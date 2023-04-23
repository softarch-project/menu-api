package service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/softarch-project/menu-api/models"
	"github.com/softarch-project/menu-api/repository"
)

type menuService struct {
	menuRepository repository.MenuRepository
}

type MenuService interface {
	GetShortMenu(ctx context.Context) ([]models.ShortMenu, error)
	GetFullMenu(ctx context.Context) ([]models.FullMenu, error)
}

func NewMenuService(menuRepository repository.MenuRepository) *menuService {
	return &menuService{
		menuRepository: menuRepository,
	}
}

func (s *menuService) GetShortMenu(ctx context.Context) (shortMenu []models.ShortMenu, err error) {
	log.Info("Getting all short menu(s)...")
	defer log.Info("End getting all short menu(s).")

	shortMenu, err = s.menuRepository.QueryAllShortMenu(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Get all short menu(s) successfully")
	return shortMenu, err
}

func (s *menuService) GetFullMenu(ctx context.Context) (fullMenus []models.FullMenu, err error) {
	log.Info("Getting all short menu(s)...")
	defer log.Info("End getting all short menu(s).")

	fullMenus, err = s.menuRepository.QueryAllFullMenu(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Get all short menu(s) successfully")
	return fullMenus, err
}
