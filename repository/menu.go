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
	// QueryAllFullMenu() ([]models.FullMenu, error)
	// InsertShortMenu(models.ShortMenu) error
	// InsertFullMenu(models.FullMenu) error
}

func NewMenuRepository(db *sqlx.DB) *menuRepository {
	return &menuRepository{
		db: db,
	}
}

var ErrFoundMoreThanOne error = errors.New("found more than one row in db")
var ErrNotFound error = errors.New("not found in db")

// func (r *menuRepository) InsertShortMenu(shortMenuDAO models.ShortMenuDAO) error {
// 	logger := generateLogger("InsertMenu")

// 	_, err := r.db.Query(`
// 		INSERT INTO`+"`menu.shortMenu`"+`(id, name, thumbnailImage, fullPrice, discountedPercent, discountedTimePeriodId, sold, totalInStock)
// 		VALUES (?, ?, ?, ?, ?)
// 	`,
// 		shortMenuDAO.Id,
// 		shortMenuDAO.Name,
// 		shortMenuDAO.ThumbnailImage,
// 		shortMenuDAO.FullPrice,
// 		shortMenuDAO.DiscountedPercent,
// 		shortMenuDAO.DiscountedTimePeriodId,
// 		shortMenuDAO.Sold,
// 		shortMenuDAO.TotalInStock,
// 	)
// 	if err != nil {
// 		logger.Error(err)
// 		return err
// 	}

// 	logger.Info("Insert short menu")
// 	return nil
// }

// func (r *menuRepository) InsertFuLLMenu(fullMenuDAO models.FullMenuDAO) error {
// 	logger := generateLogger("InsertMenu")

// 	_, err := r.db.Query(`
// 		INSERT INTO`+"`menu.shortMenu`"+`(id, name, thumbnailImage, fullPrice, discountedPercent, discountedTimePeriodId, sold, totalInStock, LargeImage, optionsId)
// 		VALUES (?, ?, ?, ?, ?)
// 	`,
// 		fullMenuDAO.Id,
// 		fullMenuDAO.Name,
// 		fullMenuDAO.ThumbnailImage,
// 		fullMenuDAO.FullPrice,
// 		fullMenuDAO.DiscountedPercent,
// 		fullMenuDAO.DiscountedTimePeriodId,
// 		fullMenuDAO.Sold,
// 		fullMenuDAO.TotalInStock,
// 		fullMenuDAO.LargeImage,
// 		fullMenuDAO.OptionsId,
// 	)
// 	if err != nil {
// 		logger.Error(err)
// 		return err
// 	}

// 	logger.Info("Insert full menu")
// 	return nil
// }

func (r *menuRepository) QueryAllShortMenu() (shortMenus []models.ShortMenu, err error) {
	logger := generateLogger("QueryAllShortMenu")

	var menu []models.ShortMenuDAO
	err = r.db.Select(&menu, `
		SELECT Menu.id, name, thumbnailImage, fullPrice, discountedPercent, DiscountedTimePeriod.begin,
		DiscountedTimePeriod.end, sold, totalInStock
		FROM Menu JOIN DiscountedTimePeriod ON Menu.id = DiscountedTimePeriod.id
		`)

	if err != nil {
		logger.Error(err)
		return shortMenus, err
	}

	menuLength := len(menu)
	if menuLength == 0 {
		logger.Error(ErrNotFound)
		return shortMenus, ErrNotFound
	}

	for _, m := range menu {
		shortMenus = append(shortMenus,
			models.ShortMenu{
				Id:                m.Id,
				Name:              m.Name,
				ThumbnailImage:    m.ThumbnailImage,
				FullPrice:         m.FullPrice,
				DiscountedPercent: m.DiscountedPercent,
				DiscountedTimePeriod: struct {
					Begin string "json:\"begin\" db:\"begin\""
					End   string "json:\"end\" db:\"end\""
				}{m.Begin, m.End},
				Sold:         m.Sold,
				TotalInStock: m.TotalInStock,
			},
		)
	}
	return shortMenus, nil
}
