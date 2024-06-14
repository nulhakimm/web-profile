package web

type ExperienceResponse struct {
	CompanyName string
	Address     string
	Title       string
	EntryYear   string
	OutYear     string
	JobDesk     []string
}

type ExperienceCreate struct {
	CompanyName string
	Address     string
	Title       string
	EntryYear   string
	OutYear     string
	JobDesk     []string
}
