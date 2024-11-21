package types

type AssetDependencyReq struct {
	ID           int `json:"id"`
	AssetID      int `json:"asset_id"`
	DependencyID int `json:"dependency_id"`
}
