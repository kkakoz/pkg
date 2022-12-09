package https

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type HttpServer struct {
	server *http.Server
	addr   string

	handler *httpHandler
}

func NewHttpServer(handler http.Handler, addr string) *HttpServer {
	newHandler := &httpHandler{handler: handler}
	return &HttpServer{
		handler: newHandler,
		server: &http.Server{
			Addr:    addr,
			Handler: newHandler},
	}
}

type httpHandler struct {
	handler  http.Handler
	isReject bool
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.isReject {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("服务已关闭"))
		return
	}
	h.handler.ServeHTTP(w, r)
}

func (h *HttpServer) Start(ctx context.Context) error {
	err := h.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("listen and server http err: %s", err)
	}
	return nil
}

func (h *HttpServer) Stop(ctx context.Context) error {
	h.handler.isReject = true
	fmt.Println("stop http server")
	shutDownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := h.server.Shutdown(shutDownCtx); err != nil {
		return fmt.Errorf("server shutdown err: %s", err)
	}
	fmt.Println("final stop http server")
	return nil
}
