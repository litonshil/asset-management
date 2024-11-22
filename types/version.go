package types

type VersionReq struct {
	ID            int    `json:"id"`
	AssetID       int    `json:"asset_id"`
	VersionNumber string `json:"version_number"`
	Description   string `json:"description"`
	CreatedBy     int    `json:"created_by"`
}

type VersionResp struct {
	ID            int    `json:"id"`
	AssetID       int    `json:"asset_id"`
	VersionNumber string `json:"version_number"`
	Description   string `json:"description"`
	CreatedBy     int    `json:"created_by"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
