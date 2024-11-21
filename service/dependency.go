package service

import (
	"asset-management/domain"
	"asset-management/types"
)

type dependencyService struct {
	dependencyRepo domain.DependencyRepository
}

func NewDependencyService(dependencyRepo domain.DependencyRepository) domain.DependencyUseCase {
	return &dependencyService{dependencyRepo}
}

func (s *dependencyService) CreateDependency(req types.AssetDependencyReq) (domain.Dependency, error) {
	// TODO: implement this

	return domain.Dependency{}, nil
}

func (s *dependencyService) GetDependencies(assetID int) ([]domain.Dependency, error) {
	// TODO: implement this

	return []domain.Dependency{}, nil
}

func (s *dependencyService) DeleteDependency(assetID int, dependencyID int) error {
	// TODO: implement this

	return nil
}
