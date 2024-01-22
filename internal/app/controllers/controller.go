package controller

import (
	"net/http"
	service "ui-test-manager/internal/app/services"
	"ui-test-manager/internal/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	s service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{
		s: *s,
	}
}

func (c *Controller) GetSequences(ctx *gin.Context) {
	sequences, err := c.s.GetSequences()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, *sequences)
}

func (c *Controller) SaveSequence(ctx *gin.Context) {
	var json model.Sequence
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.s.SaveSequence(json)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) DeleteSequence(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing sequence ID"})
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.s.DeleteSequence(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) GetPatterns(ctx *gin.Context) {
	patterns, err := c.s.GetPatterns()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, *patterns)
}

func (c *Controller) SavePattern(ctx *gin.Context) {
	var json model.Pattern
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.s.SavePattern(json)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) DeletePattern(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing sequence ID"})
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.s.DeletePattern(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) DeletePrecedent(ctx *gin.Context) {

	idSequence, idPrecedent := ctx.Param("id"), ctx.Param("idPrecedent")

	if idSequence == "" || idPrecedent == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID"})
		return
	}
	uuidSequence, err := uuid.Parse(idSequence)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuidPrecedent, err := uuid.Parse(idPrecedent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.s.DeletePrecedent(uuidSequence, uuidPrecedent)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) GetTests(ctx *gin.Context) {
	tests, err := c.s.GetTests()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tests)
}
