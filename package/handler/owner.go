package handler

import (
	"autokatolog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateOwner(c *gin.Context) {
	var owner autokatolog.People
	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.CreateOwner(owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}
