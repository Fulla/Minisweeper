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

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization")
// 		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 		} else {
// 			c.Next()
// 		}
// 	}
// }

func (s *Server) newRouter() *gin.Engine {
	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type"},
	// 	AllowCredentials: true,
	// 	ExposeHeaders:    []string{"Content-Length"},
	// }))
	r.OPTIONS("/:path", func(c *gin.Context) {
		log.Infof("getting option")
		c.AbortWithStatus(204)
	})

	r.Use(static.Serve("/", static.LocalFile("client/build", true)))

	api := r.Group("/api")
	api.POST("/new", s.startGame)
	api.GET("/resume", s.resumeGame)
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
