package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/nulhakimm/web-profile/helper"
	"github.com/nulhakimm/web-profile/model/domain"
	"github.com/nulhakimm/web-profile/model/web"
	"github.com/nulhakimm/web-profile/repository"
)

type EducationService interface {
	Find(ctx context.Context) ([]*web.EducationResponse, error)
	Save(ctx context.Context, request *web.EducationCreate) error
}

type EducationServiceImpl struct {
	EducationRepo repository.EducationRepo
	Client        *firestore.Client
}

func NewEducationService(educationRepo repository.EducationRepo, client *firestore.Client) EducationService {
	return &EducationServiceImpl{
		EducationRepo: educationRepo,
		Client:        client,
	}
}

func (service *EducationServiceImpl) Find(ctx context.Context) ([]*web.EducationResponse, error) {

	eduResponse, err := service.EducationRepo.Find(ctx, service.Client)
	if err != nil {
		return nil, err
	}

	return helper.ToEducationResponses(eduResponse), nil

}

func (service *EducationServiceImpl) Save(ctx context.Context, request *web.EducationCreate) error {

	education := &domain.Education{
		Name:      request.Name,
		Address:   request.Address,
		EntryYear: request.EntryYear,
		OutYear:   request.OutYear,
		Title:     request.Title,
		Achiev:    request.Achiev,
	}

	err := service.EducationRepo.Create(ctx, service.Client, education)
	if err != nil {
		return err
	}

	return nil

}
