package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

func newRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func Serve(c context.Context) {
	r := newRouter()
	server := &http.Server{
		Addr:    "",
		Handler: r,
	}
	go func() {
		<-c.Done()
		log.Info("shutdown from context")
		server.Shutdown(c)
	}()

	err := server.ListenAndServe()
	if err != nil {
		log.Error(err)
	}
}
