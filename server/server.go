package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Fulla/Minisweeper/gamesmanager"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	api.POST("/discover", s.discoverTile)
	return r
}

func (s *Server) Serve(c context.Context) {
	r := s.newRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
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
