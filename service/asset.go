package service

import (
	"asset-management/domain"
	"asset-management/types"
)

type assetService struct {
	assetRepo domain.AssetRepository
}

func NewAssetService(assetRepo domain.AssetRepository) domain.AssetUseCase {
	return &assetService{assetRepo}
}

func (s *assetService) CreateAsset(req types.AssetReq) (types.AssetResp, error) {
	// TODO: implement this

	return types.AssetResp{}, nil
}

func (s *assetService) GetAssets() ([]types.AssetResp, error) {
	// TODO: implement this

	return []types.AssetResp{}, nil
}

func (s *assetService) GetAsset(id int) (types.AssetResp, error) {
	// TODO: implement this

	return types.AssetResp{}, nil
}

func (s *assetService) UpdateAsset(req types.AssetReq) error {
	// TODO: implement this

	return nil
}

func (s *assetService) DeleteAsset(id int) error {
	// TODO: implement this

	return nil
}

func (s *assetService) AssignAsset(req types.AssignedAssetReq) error {
	// TODO: implement this

	return nil
}

func (s *assetService) UnAssignAsset(id int) error {
	// TODO: implement this

	return nil
}
