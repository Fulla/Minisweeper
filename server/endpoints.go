package server

import (
	"github.com/Fulla/Minisweeper/game"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Dimensions struct {
	Files   int `json:"files"`
	Columns int `json:"columns"`
	Mines   int `json:"mines"`
}

func (s *Server) startGame(c *gin.Context) {
	var size Dimensions
	err := c.BindJSON(&size)
	if err != nil {
		c.Error(err)
		return
	}
	logrus.Infof("starting game with configuration: %+v", size)
	g := s.gameMgr.StartGame(c.Request.Context(), size.Files, size.Columns, size.Mines)
	defer s.gameMgr.FreeGame()
	exported, err := s.gameMgr.ExportClientBoard(g)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(201, gin.H{"data": exported})
}

func (s *Server) resumeGame(c *gin.Context) {
	g := s.gameMgr.GetGame(c.Request.Context())
	defer s.gameMgr.FreeGame()
	exported, err := s.gameMgr.ExportClientBoard(g)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{"data": exported})
}

func (s *Server) discoverTile(c *gin.Context) {
	g := s.gameMgr.GetGame(c.Request.Context())
	defer s.gameMgr.FreeGame()
	state := g.State()
	if state != "initial" && state != "playing" {
		c.AbortWithError(400, errors.Errorf("Bad action: Trying to discover tile while not in active game"))
		return
	}
	var point game.Point
	err := c.BindJSON(&point)
	if err != nil {
		c.Error(err)
		return
	}
	g.Discover(point)
	exported, err := s.gameMgr.ExportClientBoard(g)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{"data": exported})
}
