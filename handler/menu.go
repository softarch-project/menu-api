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

func (h *menuHandler) GetAllFullMenu(c *gin.Context) {
	fullMenus, err := h.menuService.GetAllFullMenu(c)
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

func (h *menuHandler) GetShortMenu(c *gin.Context) {
	shortMenu, err := h.menuService.GetShortMenu(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, responses.MenuResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    shortMenu,
	})
}

func (h *menuHandler) GetFullMenu(c *gin.Context) {
	fullMenu, err := h.menuService.GetFullMenu(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, responses.MenuResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    fullMenu,
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

func (h *menuHandler) InsertMenu(c *gin.Context) {
	menu, err := h.menuService.InsertMenu(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, responses.MenuResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    menu,
	})
}
