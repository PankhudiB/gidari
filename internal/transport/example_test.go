package transport_test

import (
	"net/http"
	"net/url"

	"github.com/alpine-hodler/gidari/internal/transport"
	"github.com/alpine-hodler/gidari/pkg/repository"
)

type CustomRepositoryEncoder struct{}

func (e *CustomRepositoryEncoder) Encode(_ http.Request, _ []byte) (*repository.Raw, error) {
	// Do something with the request and data to create a repository.Raw object.
	return nil, nil
}

func ExampleRepositoryEncoderRegistry() {
	// If necessary, you can register your own RepositoryEncoder for a specific host. Of course, this would require
	// a custom build of the Gidari library.
	u, _ := url.Parse("http://test")
	transport.RepositoryEncoders.Register(u, new(CustomRepositoryEncoder))
}