package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/softarch-project/menu-api/responses"
	"github.com/softarch-project/menu-api/service"
)

type menuHandler struct {
	menuService service.MenuService
}

func NewMenuHandler(menuService service.MenuService) menuHandler {
	return menuHandler{
		menuService: menuService,
	}
}

func (h *menuHandler) GetAllShortMenu(c *gin.Context) {
	shortMenus, err := h.menuService.GetAllShortMenu(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, responses.MenuResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    shortMenus,
	})
}
