package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/nulhakimm/web-profile/model/domain"
	"google.golang.org/api/iterator"
)

type SkillRepo interface {
	Find(ctx context.Context, client *firestore.Client) ([]*domain.Skill, error)
	Save(ctx context.Context, client *firestore.Client, skill *domain.Skill) error
}

type SkillRepoImpl struct {
}

func NewSkillRepo() SkillRepo {
	return &SkillRepoImpl{}
}

func (repository *SkillRepoImpl) Find(ctx context.Context, client *firestore.Client) ([]*domain.Skill, error) {

	iter := client.Collection("skills").Documents(ctx)

	defer iter.Stop()

	skills := []*domain.Skill{}

	for {

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to iterate through documents: %v", err)
		}

		var skill domain.Skill
		err = doc.DataTo(&skill)
		if err != nil {
			return nil, fmt.Errorf("failed to decode document data: %v", err)
		}

		skills = append(skills, &skill)

	}

	return skills, nil

}

func (repository *SkillRepoImpl) Save(ctx context.Context, client *firestore.Client, skill *domain.Skill) error {

	_, _, err := client.Collection("skills").Add(ctx, skill)
	if err != nil {
		return fmt.Errorf("failed to add document : %v", err)
	}

	return nil

}
