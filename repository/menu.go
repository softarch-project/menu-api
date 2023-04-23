package repository

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	InsertMenu(ctx *gin.Context) (models.FullMenu, error)
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
	logger.Info("Delete menu successfully")
	return nil
}

func (r *menuRepository) InsertMenu(c *gin.Context) (newMenu models.FullMenu, err error) {
	logger := generateLogger("InsertMenu")
	var validate = validator.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var menu models.FullMenu
	defer cancel()

	if err := c.BindJSON(&menu); err != nil {
		logger.Error(err)
		return newMenu, err
	}

	if validationErr := validate.Struct(&menu); validationErr != nil {
		logger.Error(validationErr)
		return newMenu, validationErr
	}

	newMenu = models.FullMenu{
		Id:                   primitive.NewObjectID(),
		Name:                 menu.Name,
		ThumbnailImage:       menu.ThumbnailImage,
		FullPrice:            menu.FullPrice,
		DiscountedPercent:    menu.DiscountedPercent,
		DiscountedTimePeriod: menu.DiscountedTimePeriod,
		Sold:                 menu.Sold,
		TotalInStock:         menu.TotalInStock,
		LargeImage:           menu.LargeImage,
		Options:              menu.Options,
	}

	result, err := r.resourceCollection.InsertOne(ctx, newMenu)
	if err != nil {
		logger.Error(err)
		return newMenu, err
	}
	logger.Info("result.InsertedID: %v\n", result.InsertedID)

	return newMenu, nil
}
