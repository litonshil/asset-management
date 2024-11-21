package controller

import (
	"asset-management/domain"
	"asset-management/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AssetController struct {
	assetSvc domain.AssetUseCase
}

func NewAssetController(assetService domain.AssetUseCase) *AssetController {
	return &AssetController{assetService}
}

func (c *AssetController) CreateAsset(ctx echo.Context) error {

	assetReq := types.AssetReq{}

	if err := ctx.Bind(&assetReq); err != nil {
		return err
	}

	_, err := c.assetSvc.CreateAsset(assetReq)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "successfully added")
}

func (c *AssetController) GetAssets(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, []domain.Asset{})
}

func (c *AssetController) GetAsset(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, domain.Asset{})
}

func (c *AssetController) UpdateAsset(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "successfully updated")
}

func (c *AssetController) DeleteAsset(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "successfully deleted")
}
