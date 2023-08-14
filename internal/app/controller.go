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
		name := c.Param("name")
		content := c.Param("content")
		json = *model.NewSequence(name, content)
	}
	persistence.SaveSequence(&json)
	c.JSON(http.StatusOK, json.Id.String())
}

func DeleteSequence(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	persistence.DeleteSequence(&id)
	c.JSON(http.StatusOK, id.String())
}
