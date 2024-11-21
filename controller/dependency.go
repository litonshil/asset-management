package controller

import (
	"asset-management/domain"
	"github.com/labstack/echo/v4"
)

type DependencyController struct {
	dependencySvc domain.DependencyUseCase
}

func NewDependencyController(dependencySvc domain.DependencyUseCase) *DependencyController {
	return &DependencyController{dependencySvc}
}

func (c *DependencyController) CreateDependencies(ctx echo.Context) error {
	// TODO: implement this

	return nil
}

func (c *DependencyController) GetDependencies(ctx echo.Context) error {
	// TODO: implement this

	return nil
}

func (c *DependencyController) DeleteDependency(ctx echo.Context) error {
	// TODO: implement this

	return nil
}
