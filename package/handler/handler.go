package handler

import (
	_ "autokatolog/docs"
	"autokatolog/package/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cars := router.Group("/cars")
	{
		cars.POST("/", h.createCar)
		cars.GET("/", h.getAllCars)
		cars.GET("/:reg_num", h.getCarByRegNum)
		cars.PUT("/:reg_num", h.updateCar)
		cars.DELETE("/:reg_num", h.deleteCar)
	}

	owners := router.Group("/owners")
	{
		owners.POST("/", h.CreateOwner)
	}

	return router
}
