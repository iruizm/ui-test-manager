package api

import (
	"net/http"
	"polonium/internal/pkg/model"
	"polonium/internal/pkg/persistence"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetSequences(c *gin.Context) {
	sequences := persistence.GetSequences()
	c.JSON(http.StatusOK, *sequences)
}

func SaveSequence(c *gin.Context) {
	var json model.Sequence
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sequence"})
		return
	}
	if id := c.Param("id"); id == "" {
		json = *model.NewSequence(json.Name, json.Content)
	}

	persistence.SaveSequence(&json)
	c.JSON(http.StatusOK, json.Id.String())
}

func DeleteSequence(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing sequence ID"})
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	persistence.DeleteSequence(&uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sequence"})
		return
	}
	c.JSON(http.StatusOK, uuid.String())
}
