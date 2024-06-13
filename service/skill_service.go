package service

import (
	"context"
	"fmt"

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

	skillResponses, err := service.SkillRepo.Find(ctx, service.Client)
	if err != nil {
		return nil, fmt.Errorf("failed get data profile : %v", err)
	}

	return helper.ToSkillResponses(skillResponses), nil

}

func (service *SkillServiceImpl) Save(ctx context.Context, request *web.SkillCreate) error {

	skill := &domain.Skill{
		CompanyName: request.CompanyName,
		Address:     request.Address,
		Title:       request.Title,
		EntryYear:   request.EntryYear,
		OutYear:     request.OutYear,
		JobDesk:     request.JobDesk,
	}

	err := service.SkillRepo.Save(ctx, service.Client, skill)
	if err != nil {
		return err
	}

	return nil

}
