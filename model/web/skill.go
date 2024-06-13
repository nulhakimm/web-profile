package web

type SkillResponse struct {
	CompanyName string
	Address     string
	Title       string
	EntryYear   string
	OutYear     string
	JobDesk     []string
}

type SkillCreate struct {
	CompanyName string
	Address     string
	Title       string
	EntryYear   string
	OutYear     string
	JobDesk     []string
}
