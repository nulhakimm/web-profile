package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/nulhakimm/web-profile/helper"
	"github.com/nulhakimm/web-profile/model/domain"
	"github.com/nulhakimm/web-profile/model/web"
	"github.com/nulhakimm/web-profile/repository"
)

type SkillService interface {
	Find(ctx context.Context) ([]*web.SkillResponse, error)
	Save(ctx context.Context, request *web.SkillCreate) error
}

type SkillServiceImpl struct {
	SkillRepo repository.SkillRepo
	Client    *firestore.Client
}

func NewSkillService(skillRepo repository.SkillRepo, client *firestore.Client) SkillService {
	return &SkillServiceImpl{
		SkillRepo: skillRepo,
		Client:    client,
	}
}

func (service *SkillServiceImpl) Find(ctx context.Context) ([]*web.SkillResponse, error) {

	skillResponse, err := service.SkillRepo.Find(ctx, service.Client)
	if err != nil {
		return nil, err
	}

	return helper.ToSkillResponses(skillResponse), nil

}

func (service *SkillServiceImpl) Save(ctx context.Context, request *web.SkillCreate) error {

	skill := &domain.Skill{
		Name: request.Name,
	}

	err := service.SkillRepo.Save(ctx, service.Client, skill)
	if err != nil {
		return err
	}

	return nil

}
