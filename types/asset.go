package types

import (
	"gorm.io/datatypes"
	"time"
)

type AssetReq struct {
	Title       string         `json:"title"`
	Type        string         `json:"type"`
	UserID      int            `json:"user_id"`
	Description string         `json:"description"`
	FileURL     string         `json:"file_url"`
	Metadata    datatypes.JSON `json:"metadata"`
}

type AssetResp struct {
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

type AssetLicenseReq struct {
	AssetID    int   `json:"asset_id"`
	LicenseIDs []int `json:"license_ids"`
}

type AssignedAssetReq struct {
	ID         int `json:"id"`
	AssetID    int `json:"asset_id"`
	UserID     int `json:"user_id"`
	AssignedBy int `json:"assigned_by"`
}
