package repository

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/softarch-project/menu-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type menuRepository struct {
	db                 *mongo.Database
	resourceCollection *mongo.Collection
}

type MenuRepository interface {
	QueryAllShortMenu(ctx *gin.Context) ([]models.ShortMenu, error)
	QueryAllFullMenu(ctx *gin.Context) ([]models.FullMenu, error)
	DeleteMenu(ctx *gin.Context) error
	// InsertShortMenu(models.ShortMenu) error
	// InsertFullMenu(models.FullMenu) error
}

func NewMenuRepository(resourceCollection *mongo.Collection) *menuRepository {
	return &menuRepository{
		resourceCollection: resourceCollection,
	}
}

func (r *menuRepository) QueryAllShortMenu(ctx *gin.Context) ([]models.ShortMenu, error) {
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

func (r *menuRepository) QueryAllFullMenu(ctx *gin.Context) ([]models.FullMenu, error) {
	logger := generateLogger("QueryAllFullMenu")

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

func (r *menuRepository) DeleteMenu(c *gin.Context) (err error) {
	logger := generateLogger("DeleteMenu")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	menuId := c.Param("menuId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(menuId)
	logger.Info(objId)
	res, err := r.resourceCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		logger.Error(err)
		return err
	}
	if res.DeletedCount < 1 {
		logger.Warnf("no resources found: %v", err)
		return mongo.ErrNoDocuments
	}

	return nil
}
