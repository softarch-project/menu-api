package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/softarch-project/menu-api/models"
	"github.com/softarch-project/menu-api/repository"
)

type menuService struct {
	menuRepository repository.MenuRepository
}

type MenuService interface {
	GetShortMenu(ctx *gin.Context) ([]models.ShortMenu, error)
	GetFullMenu(ctx *gin.Context) ([]models.FullMenu, error)
	DeleteMenu(ctx *gin.Context) error
}

func NewMenuService(menuRepository repository.MenuRepository) *menuService {
	return &menuService{
		menuRepository: menuRepository,
	}
}

func (s *menuService) GetShortMenu(ctx *gin.Context) (shortMenu []models.ShortMenu, err error) {
	log.Info("Getting all short menu(s)...")
	defer log.Info("End getting all short menu(s).")

	shortMenu, err = s.menuRepository.QueryAllShortMenu(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("Get all short menu(s) successfully")
	return shortMenu, nil
}

func (s *menuService) GetFullMenu(ctx *gin.Context) (fullMenus []models.FullMenu, err error) {
	log.Info("Getting all short menu(s)...")
	defer log.Info("End getting all short menu(s).")

	fullMenus, err = s.menuRepository.QueryAllFullMenu(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("Get all short menu(s) successfully")
	return fullMenus, nil
}

func (s *menuService) DeleteMenu(ctx *gin.Context) (err error) {
	log.Infof("Deleting menu with id: %d", ctx.Param("menuId"))
	defer log.Infof("End deleting menu with id: %d", ctx.Param("menuId"))

	err = s.menuRepository.DeleteMenu(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("Delete menu successfully")
	return
}
