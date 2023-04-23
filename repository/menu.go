package repository

import (
	"context"
	"errors"

	"github.com/softarch-project/menu-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "gopkg.in/mgo.v2/bson"
)

type menuRepository struct {
	db                 *mongo.Database
	resourceCollection *mongo.Collection
}

type MenuRepository interface {
	QueryAllShortMenu(ctx context.Context) ([]models.ShortMenu, error)
	QueryAllFullMenu(ctx context.Context) ([]models.FullMenu, error)
	// InsertShortMenu(models.ShortMenu) error
	// InsertFullMenu(models.FullMenu) error
}

func NewMenuRepository(resourceCollection *mongo.Collection) *menuRepository {
	return &menuRepository{
		resourceCollection: resourceCollection,
	}
}

var ErrFoundMoreThanOne error = errors.New("found more than one row in db")
var ErrNotFound error = errors.New("not found in db")

func (r *menuRepository) QueryAllShortMenu(ctx context.Context) ([]models.ShortMenu, error) {
	logger := generateLogger("QueryAllShortMenu")

	var shortMenus []models.ShortMenu

	results, err := r.resourceCollection.Find(ctx, bson.M{})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Warnf("no resources found: %v", err)
			return nil, err
		}
		logger.Warnf("find resources failed: %v", err)
		return nil, err
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var menu models.ShortMenu
		if err = results.Decode(&menu); err != nil {
			logger.Warnf("decode resource failed: %v", err)
			return nil, err
		}
		shortMenus = append(shortMenus, menu)
	}

	logger.Info("find short menus successfully")
	return shortMenus, nil
}

func (r *menuRepository) QueryAllFullMenu(ctx context.Context) ([]models.FullMenu, error) {
	logger := generateLogger("QueryAllShortMenu")

	var fullMenus []models.FullMenu

	results, err := r.resourceCollection.Find(ctx, bson.M{})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Warnf("no resources found: %v", err)
			return nil, err
		}
		logger.Warnf("find resources failed: %v", err)
		return nil, err
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var menu models.FullMenu
		if err = results.Decode(&menu); err != nil {
			logger.Warnf("decode resource failed: %v", err)
			return nil, err
		}
		fullMenus = append(fullMenus, menu)
	}

	logger.Info("find full menus successfully")
	return fullMenus, nil
}
