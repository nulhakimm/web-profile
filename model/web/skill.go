package web

type SkillResponse struct {
	Name string `json:"name,omitempty"`
}

type SkillCreate struct {
	Name string `json:"name,omitempty"`
}
