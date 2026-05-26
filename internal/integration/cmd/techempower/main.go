package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"go.opentelemetry.io/otel/trace/noop"

	"github.com/ogen-go/ogen/internal/integration/techempower"
)

type server struct{}

func (server) JSON(ctx context.Context) (*techempower.HelloWorld, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (server) DB(ctx context.Context) (*techempower.WorldObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
func (server) Caching(ctx context.Context, params techempower.CachingParams) (techempower.WorldObjects, error) {
	_ = "STUB: not implemented"
	return *new(techempower.WorldObjects), nil
}

func (server) Queries(ctx context.Context, params techempower.QueriesParams) (techempower.WorldObjects, error) {
	_ = "STUB: not implemented"
	return *new(techempower.WorldObjects), nil
}

func (server) Updates(ctx context.Context, params techempower.UpdatesParams) (techempower.WorldObjects, error) {
	_ = "STUB: not implemented"
	return *new(techempower.WorldObjects), nil
}

func main() {
	var arg struct {
		Addr string
	}
	flag.StringVar(&arg.Addr, "addr", ":8080", "http address to listen")
	flag.Parse()

	traceProvider := noop.NewTracerProvider()
	s, err := techempower.NewServer(&server{}, techempower.WithTracerProvider(traceProvider))
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(arg.Addr, s))
}
