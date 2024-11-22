package domain

import (
	"time"
)

type AssetLicense struct {
	ID         int       `json:"id"`
	AssetID    int       `json:"asset_id"`
	LicenseIDs []int     `json:"license_ids"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
