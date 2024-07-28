package invoke

import (
	"github.com/enuesaa/walkhttp/pkg/serve/schema"
)

func NewConfig() Workspace {
	return Workspace{
		BaseUrl: "https://",
	}
}

type Workspace struct {
	BaseUrl string `json:"baseUrl"` // like `https://example.com`
	// Histories []History `json:"histories"`
}

type History struct {
	Path    string          `json:"path"`    // like `/aaa`
	Method  string          `json:"method"`  // like `GET`
	Headers []schema.Header `json:"headers"` // default headers
	Body    string          `json:"body"`    // default body
}
