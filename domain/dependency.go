package domain

import (
	"asset-management/types"
	"gorm.io/gorm"
	"time"
)

type Dependency struct {
	ID           int            `json:"id"`
	AssetID      int            `json:"asset_id"`
	DependencyID int            `json:"dependency_id"`
	Description  string         `json:"description"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type DependencyUseCase interface {
	CreateDependency(req types.AssetDependencyReq) (Dependency, error)
	GetDependencies(assetID int) ([]Dependency, error)
	DeleteDependency(assetID int, dependencyID int) error
}

type DependencyRepository interface {
	CreateDependency(req Dependency) (Dependency, error)
	GetDependencies(assetID int) ([]Dependency, error)
	DeleteDependency(assetID int, dependencyID int) error
}
