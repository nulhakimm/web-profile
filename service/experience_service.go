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

type ExperienceService interface {
	Find(ctx context.Context) ([]*web.ExperienceResponse, error)
	Save(ctx context.Context, request *web.ExperienceCreate) error
}

type ExperienceServiceImpl struct {
	ExperienceRepo repository.ExperienceRepo
	Client         *firestore.Client
}

func NewExperienceService(ExperienceRepo repository.ExperienceRepo, client *firestore.Client) ExperienceService {
	return &ExperienceServiceImpl{
		ExperienceRepo: ExperienceRepo,
		Client:         client,
	}
}

func (service *ExperienceServiceImpl) Find(ctx context.Context) ([]*web.ExperienceResponse, error) {

	experienceResponses, err := service.ExperienceRepo.Find(ctx, service.Client)
	if err != nil {
		return nil, fmt.Errorf("failed get data profile : %v", err)
	}

	return helper.ToExperienceResponses(experienceResponses), nil

}

func (service *ExperienceServiceImpl) Save(ctx context.Context, request *web.ExperienceCreate) error {

	experience := &domain.Experience{
		CompanyName: request.CompanyName,
		Address:     request.Address,
		Title:       request.Title,
		EntryYear:   request.EntryYear,
		OutYear:     request.OutYear,
		JobDesk:     request.JobDesk,
	}

	err := service.ExperienceRepo.Save(ctx, service.Client, experience)
	if err != nil {
		return err
	}

	return nil

}
