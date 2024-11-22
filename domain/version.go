package domain

import (
	"asset-management/types"
	"time"
)

type Version struct {
	ID            int       `json:"id"`
	AssetID       int       `json:"asset_id"`
	VersionNumber string    `json:"version_number"`
	Description   string    `json:"description"`
	CreatedBy     int       `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type VersionUseCase interface {
	CreateVersion(req types.VersionReq) (types.VersionResp, error)
	GetVersions() ([]types.VersionResp, error)
	GetVersion(id int) (types.VersionResp, error)
	UpdateVersion(req types.VersionReq) error
	DeleteVersion(id int) error
}

type VersionRepository interface {
	CreateVersion(req Version) (Version, error)
	GetVersions() ([]Version, error)
	GetVersion(id int) (Version, error)
	UpdateVersion(req Version) error
	DeleteVersion(id int) error
}
