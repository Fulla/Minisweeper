package server

import (
	"context"
	"net/http"

	"github.com/Fulla/Minisweeper/gamesmanager"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

type Server struct {
	gameMgr *gamesmanager.GamesManager
}

func NewServer(mgr *gamesmanager.GamesManager) *Server {
	return &Server{
		gameMgr: mgr,
	}
}

func (s *Server) newRouter() *gin.Engine {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("client/build", true)))

	api := r.Group("/api")
	api.POST("/new", s.startGame)
	api.GET("/resume", s.resumeGame)
	return r
}

func (s *Server) Serve(c context.Context) {
	r := s.newRouter()
	httpServer := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	go func() {
		<-c.Done()
		log.Info("shutdown from context")
		httpServer.Shutdown(c)
	}()

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Error(err)
	}
}
