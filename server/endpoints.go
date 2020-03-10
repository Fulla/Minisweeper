package server

import (
	"github.com/Fulla/Minisweeper/game"
	"github.com/gin-gonic/gin"
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
	g := s.gameMgr.StartGame(size.Files, size.Columns, size.Mines)
	exported, err := s.gameMgr.ExportClientBoard(g)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(201, gin.H{"data": exported})
}

func (s *Server) resumeGame(c *gin.Context) {
	g := s.gameMgr.GetGame()
	exported, err := s.gameMgr.ExportClientBoard(g)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{"data": exported})
}

func (s *Server) discoverTile(c *gin.Context) {
	g := s.gameMgr.GetGame()
	logrus.Info("here")
	var point game.Point
	err := c.BindJSON(&point)
	if err != nil {
		c.Error(err)
		return
	}
	logrus.Infof("after parsing %+v", point)
	g.Discover(point)
	logrus.Info("after discover")
	exported, err := s.gameMgr.ExportClientBoard(g)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	logrus.Info("exporting")
	c.JSON(200, gin.H{"data": exported})
}
