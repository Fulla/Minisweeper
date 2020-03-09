package server

import "github.com/gin-gonic/gin"

func (s *Server) startGame(c *gin.Context) {
	g := s.gameMgr.StartGame(10, 10, 5)
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
