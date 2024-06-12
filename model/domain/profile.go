package domain

type Profile struct {
	Name        string
	Description string
	Email       string
	SocialMedia SocialMedia
	Phone       string
	About       string
}

type SocialMedia struct {
	LinkedIn  string
	Instagram string
	GitHub    string
}
