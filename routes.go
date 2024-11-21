package main

import (
	"asset-management/controller"
	"asset-management/repository"
	"asset-management/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	assetRepo := repository.NewAssetRepository(db)
	dependencyRepo := repository.NewDependencyRepository(db)
	assetService := service.NewAssetService(assetRepo)
	dependencyService := service.NewDependencyService(dependencyRepo)
	assetController := controller.NewAssetController(assetService)
	dependencyController := controller.NewDependencyController(dependencyService)

	assetGroup := e.Group("/api/v1/assets")

	assetGroup.POST("", assetController.CreateAsset)
	assetGroup.GET("", assetController.GetAssets)
	assetGroup.GET("/:id", assetController.GetAssets)
	assetGroup.PUT("/:id", assetController.UpdateAsset)
	assetGroup.DELETE("/:id", assetController.DeleteAsset)

	//POST /api/assets/:id/dependencies - Define dependencies for an asset
	//GET /api/assets/:id/dependencies - Get dependencies of an asset
	//DELETE /api/assets/:id/dependencies/:dependencyId - Remove a specific dependency

	assetGroup.POST("/:id/dependencies", dependencyController.CreateDependencies)
	assetGroup.GET("/:id/dependencies", dependencyController.GetDependencies)
	assetGroup.DELETE("/:id/dependencies/:dependencyId", dependencyController.DeleteDependency)

}
