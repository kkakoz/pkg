package app_test

import (
	"context"
	app2 "github.com/kkakoz/pkg/app"
	"github.com/kkakoz/pkg/app/grpcs"
	"github.com/kkakoz/pkg/app/grpcs/pb"
	"github.com/kkakoz/pkg/app/https"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
	"testing"
)

func TestApp(t *testing.T) {

	http.HandleFunc("/aa", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/aa", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})

	app := app2.NewApp("test", https.NewHttpServer(http.DefaultServeMux, ":8081"), https.NewHttpServer(mux2, ":8082"))

	err := app.Start(context.Background())
	if err != nil {
		t.Log(err)
	}
}

func TestGrpcApp(t *testing.T) {

	s := grpc.NewServer()

	reflection.Register(s)
	hellpb.RegisterHelloServer(s, &HelloHandler{})

	servers := []app2.Server{
		grpcs.NewGrpcServer(":9999", s),
	}

	app := app2.NewApp("test", servers...)

	err := app.Start(context.Background())
	if err != nil {
		t.Log(err)
	}
}

type HelloHandler struct {
	hellpb.UnsafeHelloServer
}

func (h HelloHandler) Hello(ctx context.Context, req *hellpb.Req) (*hellpb.Res, error) {
	return &hellpb.Res{
		Res: req.Name,
	}, nil
}
