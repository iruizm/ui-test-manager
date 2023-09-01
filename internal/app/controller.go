package api

import (
	"net/http"
	"polonium/internal/pkg/model"
	"polonium/internal/pkg/persistence"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetSequences(c *gin.Context) {
	sequences, err := persistence.GetSequences()
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
	zero := uuid.UUID{}
	if json.Id.String() == zero.String() {
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

func DeletePrecedent(c *gin.Context) {

	idSequence, idPrecedent := c.Param("id"), c.Param("idPrecedent")

	if idSequence == "" || idPrecedent == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID"})
		return
	}
	uuidSequence, err := uuid.Parse(idSequence)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sequence ID"})
		return
	}
	uuidPrecedent, err := uuid.Parse(idPrecedent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid precedent ID"})
		return
	}

	err = persistence.DeletePrecedent(&uuidSequence, &uuidPrecedent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete precedent"})
		return
	}
	c.JSON(http.StatusOK, uuidSequence.String()+"|"+uuidPrecedent.String())
}
