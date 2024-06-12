package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/nulhakimm/web-profile/model/domain"
	"google.golang.org/api/iterator"
)

type ProfileRepo interface {
	FindProfile(ctx context.Context, client *firestore.Client, name string) (*domain.Profile, error)
	Save(ctx context.Context, client *firestore.Client, profile *domain.Profile) error
}

type ProfileRepoImpl struct {
}

func NewProfileRepo() ProfileRepo {
	return &ProfileRepoImpl{}
}

func (repo *ProfileRepoImpl) FindProfile(ctx context.Context, client *firestore.Client, name string) (*domain.Profile, error) {

	iter := client.Collection("profiles").Where("Name", "==", name).Documents(ctx)
	defer iter.Stop()

	// Attempt to get the first document
	doc, err := iter.Next()
	if err == iterator.Done {
		// No documents found
		return nil, fmt.Errorf("profile not found")
	}
	if err != nil {
		return nil, err
	}

	var profile domain.Profile
	err = doc.DataTo(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil

}

func (repo *ProfileRepoImpl) Save(ctx context.Context, client *firestore.Client, profile *domain.Profile) error {

	_, _, err := client.Collection("profiles").Add(ctx, profile)

	if err != nil {
		return fmt.Errorf("failed create profile : %v", err)
	}

	return nil

}
