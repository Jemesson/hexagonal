package http

import (
	"github.com/gin-gonic/gin"
	"hexagonal/src/core/dto"
	"hexagonal/src/core/ports"
)

type http struct {
	gamePort ports.GamePort
}

func NewHTTPHandler(gameUseCase ports.GamePort) *http {
	return &http{
		gamePort: gameUseCase,
	}
}

func (handler *http) Get(c *gin.Context) {
	game, err := handler.gamePort.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, game)
}

func (handler *http) Create(c *gin.Context) {
	body := dto.BodyCreate{}
	c.BindJSON(&body)

	game, err := handler.gamePort.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, dto.BuildResponseCreate(game))
}

func (handler *http) RevealCell(c *gin.Context) {
	body := dto.BodyRevealCell{}
	c.BindJSON(&body)

	game, err := handler.gamePort.Reveal(c.Param("id"), body.Row, body.Col)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, dto.BuildResponseRevealCell(game))
}
