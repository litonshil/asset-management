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
	versionRepo := repository.NewVersionRepository(db)
	assetService := service.NewAssetService(assetRepo)
	dependencyService := service.NewDependencyService(dependencyRepo)
	versionService := service.NewVersionService(versionRepo)
	assetController := controller.NewAssetController(assetService)
	dependencyController := controller.NewDependencyController(dependencyService)
	versionController := controller.NewVersionController(versionService)

	assetGroup := e.Group("/api/v1/assets")
	{
		assetGroup.POST("", assetController.CreateAsset)
		assetGroup.GET("", assetController.GetAssets)
		assetGroup.GET("/:id", assetController.GetAssets)
		assetGroup.PUT("/:id", assetController.UpdateAsset)
		assetGroup.DELETE("/:id", assetController.DeleteAsset)
		assetGroup.POST("/:id/assign", assetController.AssignAsset)
		assetGroup.DELETE("/:id/unassign", assetController.UnassignAsset)
		assetGroup.POST("/:id/license", assetController.AssignLicense)

		assetGroup.POST("/:id/dependencies", dependencyController.CreateDependencies)
		assetGroup.GET("/:id/dependencies", dependencyController.GetDependencies)
		assetGroup.DELETE("/:id/dependencies/:dependencyId", dependencyController.DeleteDependency)
	}

	versionGroup := e.Group("/api/v1")
	{
		versionGroup.POST("/assets/:assetId/versions", versionController.CreateVersion)
		versionGroup.GET("/assets/:assetId/versions", versionController.GetVersions)
		versionGroup.GET("/assets/:assetId/versions/:versionId", versionController.GetVersion)
		versionGroup.PUT("/assets/:assetId/versions/:versionId", versionController.UpdateVersion)
		versionGroup.DELETE("/assets/:assetId/versions/:versionId", versionController.DeleteVersion)

	}

	//GET /assets/:assetId/versions

}
