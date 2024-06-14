package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/nulhakimm/web-profile/model/domain"
	"google.golang.org/api/iterator"
)

type EducationRepo interface {
	Find(ctx context.Context, client *firestore.Client) ([]*domain.Education, error)
	Create(ctx context.Context, client *firestore.Client, education *domain.Education) error
}

type EducationRepoImpl struct {
}

func NewEducationRepo() EducationRepo {
	return &EducationRepoImpl{}
}

func (repository *EducationRepoImpl) Find(ctx context.Context, client *firestore.Client) ([]*domain.Education, error) {

	iter := client.Collection("educations").Documents(ctx)
	defer iter.Stop()

	educations := []*domain.Education{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to iterate through documents: %v", err)
		}

		var education domain.Education
		err = doc.DataTo(&education)
		if err != nil {
			return nil, fmt.Errorf("failed to decode document : %v", err)
		}

		educations = append(educations, &education)
	}

	return educations, nil

}

func (repository *EducationRepoImpl) Create(ctx context.Context, client *firestore.Client, education *domain.Education) error {

	_, _, err := client.Collection("educations").Add(ctx, education)
	if err != nil {
		return fmt.Errorf("failed to create new document : %v", err)
	}

	return nil

}
