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

func (h *menuHandler) GetShortMenu(c *gin.Context) {
	shortMenus, err := h.menuService.GetShortMenu(c)
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

func (h *menuHandler) GetFullMenu(c *gin.Context) {
	fullMenus, err := h.menuService.GetFullMenu(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, responses.MenuResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    fullMenus,
	})
}

func (h *menuHandler) DeleteMenu(c *gin.Context) {
	err := h.menuService.DeleteMenu(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, responses.MenuResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}
