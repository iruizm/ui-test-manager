package api

import (
	"net/http"
	"polonium/internal/pkg/model"
	"polonium/internal/pkg/persistence"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserInput struct {
	Username  string `json:"username" binding:"required"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}

// SaveSequence godoc
// @Summary Saves a Sequence
// @Description Save Sequence given Sequence data
// @ID save-sequence
// @Produce json
// @Param fileName string true "File Name"
// @Param id uuid.UUID false "Sequence ID"
// @Param precedents []*Sequence false "Precedents"
// @Success 200 {object} model.Sequence
// @Router /api/sequence/save/
func SaveSequence(c *gin.Context) {
	var json model.Sequence
	err := c.ShouldBindJSON(&json)
	if err != nil {
		id := c.Param("file_name")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		json = *model.NewSequence(id)

	}
	persistence.SaveSequence(&json)
	c.JSON(http.StatusOK, json.Id.String())
}

// SaveSequence godoc
// @Summary Removes a Sequence
// @Description Remove Sequence given Sequence ID
// @ID remove-sequence
// @Produce json
// @Param id uuid.UUID true "Sequence ID"
// @Success 200 {object} model.Sequence
// @Router /api/sequence/remove/
func RemoveSequence(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	persistence.RemoveSequence(&id)
	c.JSON(http.StatusOK, id.String())
}
