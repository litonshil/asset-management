package domain

import (
	"asset-management/types"
	"gorm.io/datatypes"
	"time"
)

type Asset struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Type        string         `json:"type"`
	UserID      int            `json:"user_id"`
	Description string         `json:"description"`
	FileURL     string         `json:"file_url"`
	Metadata    datatypes.JSON `json:"metadata"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type AssetUseCase interface {
	CreateAsset(req types.AssetReq) (types.AssetResp, error)
	GetAssets() ([]types.AssetResp, error)
	GetAsset(id int) (types.AssetResp, error)
	UpdateAsset(req types.AssetReq) error
	DeleteAsset(id int) error
}

type AssetRepository interface {
	CreateAsset(req Asset) (Asset, error)
	GetAssets() ([]Asset, error)
	GetAsset(id int) (Asset, error)
	UpdateAsset(req Asset) error
	DeleteAsset(id int) error
}
