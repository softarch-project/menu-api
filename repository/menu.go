package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/softarch-project/menu-api/models"
)

type menuRepository struct {
	db *sqlx.DB
}

type MenuRepository interface {
	QueryAllShortMenu() ([]models.ShortMenu, error)
	QueryAllFullMenu() ([]models.FullMenu, error)
	InsertShortMenu(models.ShortMenu) error
	InsertFullMenu(models.FullMenu) error
}

func NewMenuRepository(db *sqlx.DB) *menuRepository {
	return &menuRepository{
		db: db,
	}
}

var ErrFoundMoreThanOne error = errors.New("found more than one row in db")
var ErrNotFound error = errors.New("not found in db")

func (r *menuRepository) InsertShortMenu(shortMenuDAO models.ShortMenuDAO) error {
	logger := generateLogger("InsertMenu")

	_, err := r.db.Query(`
		INSERT INTO`+"`menu.shortMenu`"+`(id, name, thumbnailImage, fullPrice, discountedPercent, discountedTimePeriodId, sold, totalInStock)
		VALUES (?, ?, ?, ?, ?)
	`,
		shortMenuDAO.Id,
		shortMenuDAO.Name,
		shortMenuDAO.ThumbnailImage,
		shortMenuDAO.FullPrice,
		shortMenuDAO.DiscountedPercent,
		shortMenuDAO.DiscountedTimePeriodId,
		shortMenuDAO.Sold,
		shortMenuDAO.TotalInStock,
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Insert short menu")
	return nil
}

func (r *menuRepository) InsertFuLLMenu(fullMenuDAO models.FullMenuDAO) error {
	logger := generateLogger("InsertMenu")

	_, err := r.db.Query(`
		INSERT INTO`+"`menu.shortMenu`"+`(id, name, thumbnailImage, fullPrice, discountedPercent, discountedTimePeriodId, sold, totalInStock, LargeImage, optionsId)
		VALUES (?, ?, ?, ?, ?)
	`,
		fullMenuDAO.Id,
		fullMenuDAO.Name,
		fullMenuDAO.ThumbnailImage,
		fullMenuDAO.FullPrice,
		fullMenuDAO.DiscountedPercent,
		fullMenuDAO.DiscountedTimePeriodId,
		fullMenuDAO.Sold,
		fullMenuDAO.TotalInStock,
		fullMenuDAO.LargeImage,
		fullMenuDAO.OptionsId,
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Insert full menu")
	return nil
}

func (r *menuRepository) QueryAllShortMenu() (shortMenu []models.ShortMenu, err error) {
	logger := generateLogger("QueryAllShortMenu")

	var shortMenus []models.ShortMenu
	err = r.db.Select(&shortMenus, `
		SELECT menu.id, name, thumbnailImage, fullPrice, discountedPercent, discountedTimePeriod.begin, discountedTimePeriod.end, sold, totalInStock
		FROM menu JOIN discountedTimePeriod ON menu.id == discountedTimePeriod.id
		`)

	if err != nil {
		logger.Error(err)
		return shortMenu, err
	}

	menuLength := len(shortMenu)
	if menuLength == 0 {
		logger.Error(ErrNotFound)
		return shortMenu, ErrNotFound
	}

	return shortMenu, nil
}
