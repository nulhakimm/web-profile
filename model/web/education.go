package web

type EducationResponse struct {
	Name      string
	Address   string
	EntryYear string
	OutYear   string
	Title     string
	Achiev    []string
}

type EducationCreate struct {
	Name      string
	Address   string
	EntryYear string
	OutYear   string
	Title     string
	Achiev    []string
}
