package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/softarch-project/menu-api/config"
	"github.com/softarch-project/menu-api/handler"
	"github.com/softarch-project/menu-api/repository"
	"github.com/softarch-project/menu-api/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	App      *gin.Engine
	Database *mongo.Client
	Config   *config.Config
}

func NewHTTPServer(config *config.Config, db *mongo.Client) *Server {
	gin.SetMode(config.App.GinMode)
	app := gin.Default()
	return &Server{
		App:      app,
		Database: db,
		Config:   config,
	}
}

func (server *Server) SetUpRouter() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	server.App.Use(cors.New(config))
	server.App.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	db := server.Database.Database("Menu")
	coll := db.Collection("menus")

	menuRepository := repository.NewMenuRepository(coll)

	menuService := service.NewMenuService(menuRepository)

	menuHandler := handler.NewMenuHandler(menuService)

	server.App.GET("/", handler.HealthCheckHandler)
	server.App.GET("/shortMenu", menuHandler.GetAllShortMenu)
	server.App.GET("/fullMenu", menuHandler.GetAllFullMenu)
	server.App.GET("/:menuName/short", menuHandler.GetShortMenu)
	server.App.GET("/:menuName/full", menuHandler.GetFullMenu)
	server.App.DELETE("/deleteMenu/:menuId", menuHandler.DeleteMenu)
	server.App.POST("/menu", menuHandler.InsertMenu)
}

func (server *Server) Start() {
	server.SetUpRouter()

	port := server.Config.App.Port

	log.Infof("Server is starting on port: %s", port)
	server.App.Run(":" + port)
}
