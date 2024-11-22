package service

import (
	"asset-management/domain"
	"asset-management/types"
)

type versionService struct {
	versionRepo domain.VersionRepository
}

func NewVersionService(versionRepo domain.VersionRepository) domain.VersionUseCase {
	return &versionService{versionRepo}
}

func (s *versionService) CreateVersion(req types.VersionReq) (types.VersionResp, error) {
	// TODO: implement this

	return types.VersionResp{}, nil
}

func (s *versionService) GetVersions() ([]types.VersionResp, error) {
	// TODO: implement this

	return []types.VersionResp{}, nil
}

func (s *versionService) GetVersion(id int) (types.VersionResp, error) {
	// TODO: implement this

	return types.VersionResp{}, nil
}

func (s *versionService) UpdateVersion(req types.VersionReq) error {
	// TODO: implement this

	return nil
}

func (s *versionService) DeleteVersion(id int) error {
	// TODO: implement this

	return nil
}
