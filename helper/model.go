package helper

import (
	"github.com/nulhakimm/web-profile/model/domain"
	"github.com/nulhakimm/web-profile/model/web"
)

func ToProfileResponse(profile *domain.Profile) *web.ProfileResponse {

	return &web.ProfileResponse{
		Name:        profile.Name,
		Description: profile.Description,
		Email:       profile.Email,
		SocialMedia: web.SocialMedia{
			LinkedIn:  profile.SocialMedia.LinkedIn,
			Instagram: profile.SocialMedia.Instagram,
			GitHub:    profile.SocialMedia.GitHub,
		},
		Phone: profile.Phone,
		About: profile.About,
	}

}

func ToExperienceResponse(experience *domain.Experience) *web.ExperienceResponse {

	return &web.ExperienceResponse{
		CompanyName: experience.CompanyName,
		Address:     experience.Address,
		Title:       experience.Title,
		EntryYear:   experience.EntryYear,
		OutYear:     experience.OutYear,
		JobDesk:     experience.JobDesk,
	}

}

func ToExperienceResponses(experiences []*domain.Experience) []*web.ExperienceResponse {

	experienceResponses := []*web.ExperienceResponse{}
	for _, experience := range experiences {
		experienceResponses = append(experienceResponses, ToExperienceResponse(experience))
	}

	return experienceResponses

}

func ToSkillResponse(skill *domain.Skill) *web.SkillResponse {
	return &web.SkillResponse{
		Name: skill.Name,
	}
}

func ToSkillResponses(skills []*domain.Skill) []*web.SkillResponse {

	skillResponses := []*web.SkillResponse{}

	for _, skill := range skills {
		skillResponses = append(skillResponses, ToSkillResponse(skill))
	}

	return skillResponses

}

func ToEducationResponse(education *domain.Education) *web.EducationResponse {
	return &web.EducationResponse{
		Name:      education.Name,
		Address:   education.Address,
		EntryYear: education.EntryYear,
		OutYear:   education.OutYear,
		Title:     education.Title,
		Achiev:    education.Achiev,
	}
}

func ToEducationResponses(educations []*domain.Education) []*web.EducationResponse {

	educationResponses := []*web.EducationResponse{}
	for _, education := range educations {
		educationResponses = append(educationResponses, ToEducationResponse(education))
	}

	return educationResponses

}
