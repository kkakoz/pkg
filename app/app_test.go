package app_test

import (
	"context"
	app2 "github.com/kkakoz/pkg/app"
	"github.com/kkakoz/pkg/app/https"
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

	app := app2.NewApp("test", https.NewHttpServer(http.DefaultServeMux, ":8081"))

	err := app.Start(context.Background())
	if err != nil {
		t.Log(err)
	}
}
