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
