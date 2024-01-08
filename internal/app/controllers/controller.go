package controller

import (
	"net/http"
	service "selenium-manager/internal/app/services"
	"selenium-manager/internal/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetSequences(c *gin.Context) {
	sequences, err := service.GetSequences()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, *sequences)
}

func SaveSequence(c *gin.Context) {
	var json model.Sequence
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := service.SaveSequence(json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeleteSequence(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing sequence ID"})
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := service.DeleteSequence(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeletePrecedent(c *gin.Context) {

	idSequence, idPrecedent := c.Param("id"), c.Param("idPrecedent")

	if idSequence == "" || idPrecedent == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID"})
		return
	}
	uuidSequence, err := uuid.Parse(idSequence)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuidPrecedent, err := uuid.Parse(idPrecedent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.DeletePrecedent(uuidSequence, uuidPrecedent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetTests(c *gin.Context) {
	tests, err := service.GetTests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tests)
}

func GetPatterns(c *gin.Context) {
	patterns, err := service.GetPatterns()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, *patterns)
}

func SavePattern(c *gin.Context) {
	var json model.Pattern
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := service.SavePattern(json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeletePattern(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing sequence ID"})
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := service.DeletePattern(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
