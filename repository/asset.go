package repository

import (
	"asset-management/domain"
	"gorm.io/gorm"
)

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) domain.AssetRepository {
	return &assetRepository{
		db: db,
	}
}

func (r *assetRepository) CreateAsset(req domain.Asset) (domain.Asset, error) {
	// TODO: implement this

	return domain.Asset{}, nil
}

func (r *assetRepository) GetAssets() ([]domain.Asset, error) {
	// TODO: implement this

	return []domain.Asset{}, nil
}

func (r *assetRepository) GetAsset(id int) (domain.Asset, error) {
	// TODO: implement this

	return domain.Asset{}, nil
}

func (r *assetRepository) UpdateAsset(req domain.Asset) error {
	// TODO: implement this

	return nil
}

func (r *assetRepository) DeleteAsset(id int) error {
	// TODO: implement this

	return nil
}

func (r *assetRepository) AssignAsset(req domain.AssignedAsset) error {
	// TODO: implement this

	return nil
}

func (r *assetRepository) UnAssignAsset(id int) error {
	// TODO: implement this

	return nil
}

func (r *assetRepository) AssignLicense(req domain.AssetLicense) error {
	// TODO: implement this

	return nil
}
