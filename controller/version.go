package controller

import (
	"asset-management/domain"
	"asset-management/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

type VersionController struct {
	versionSvc domain.VersionUseCase
}

func NewVersionController(versionService domain.VersionUseCase) *VersionController {
	return &VersionController{versionService}
}

func (c *VersionController) CreateVersion(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "successfully added")
}

func (c *VersionController) GetVersions(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, []types.VersionResp{})
}

func (c *VersionController) GetVersion(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, types.VersionResp{})
}

func (c *VersionController) UpdateVersion(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "successfully updated")
}

func (c *VersionController) DeleteVersion(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "successfully deleted")
}
