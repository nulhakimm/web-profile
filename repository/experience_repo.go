package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/nulhakimm/web-profile/model/domain"
	"google.golang.org/api/iterator"
)

type ExperienceRepo interface {
	Find(ctx context.Context, client *firestore.Client) ([]*domain.Experience, error)
	Save(ctx context.Context, client *firestore.Client, Experience *domain.Experience) error
}

type ExperienceImpl struct {
}

func NewExperienceRepo() ExperienceRepo {
	return &ExperienceImpl{}
}

func (repo *ExperienceImpl) Find(ctx context.Context, client *firestore.Client) ([]*domain.Experience, error) {

	iter := client.Collection("experiences").Documents(ctx)

	defer iter.Stop()

	Experiences := []*domain.Experience{}

	for {

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to iterate through documents: %v", err)
		}

		var Experience domain.Experience
		err = doc.DataTo(&Experience)
		if err != nil {
			return nil, fmt.Errorf("failed to decode document data: %v", err)
		}

		Experiences = append(Experiences, &Experience)

	}

	return Experiences, nil

}

func (repo *ExperienceImpl) Save(ctx context.Context, client *firestore.Client, Experience *domain.Experience) error {

	_, _, err := client.Collection("experiences").Add(ctx, Experience)
	if err != nil {
		return fmt.Errorf("failed to create new Experience : %v", err)
	}

	return nil

}
