package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/softarch-project/menu-api/models"
	"github.com/softarch-project/menu-api/repository"
)

type menuService struct {
	menuRepository repository.MenuRepository
}

type MenuService interface {
	GetAllShortMenu() ([]models.ShortMenu, error)
}

func NewMenuService(menuRepository repository.MenuRepository) *menuService {
	return &menuService{
		menuRepository: menuRepository,
	}
}

func (s *menuService) GetAllShortMenu() (shortMenu []models.ShortMenu, err error) {
	log.Info("Getting all short menu(s)...")
	defer log.Info("End getting all short menu(s).")

	shortMenu, err = s.menuRepository.QueryAllShortMenu()
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Get all short menu(s) successfully")
	return shortMenu, err
}
