package serve

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/enuesaa/walkin/pkg/graph"
)

func New(repos repository.Repos) Servectl {
	return Servectl{
		repos:   repos,
		wsconns: NewWsConns(),
		Port:    3000,
	}
}

type Servectl struct {
	repos   repository.Repos
	wsconns WsConns
	Port    int
}

func (s *Servectl) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *Servectl) Listen() error {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// see https://github.com/99designs/gqlgen/issues/1664
	gqhandle := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))
	gqhandle.AddTransport(&transport.Websocket{})
	app.Any("/graph", echo.WrapHandler(gqhandle))
	app.GET("/graph/playground", echo.WrapHandler(playground.Handler("graph", "/graph")))

	return app.Start(s.Addr())
}
