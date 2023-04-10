package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) generateMenu(c *gin.Context) {
	mask := strings.ToLower(c.Param("mask"))

	menu, err := h.services.Menu.GenerateMenu(mask)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, ErrServiceFailure)
		return
	}

	c.JSON(http.StatusOK, menu)
}
