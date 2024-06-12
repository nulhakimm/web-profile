package web

type ProfileResponse struct {
	Name        string      `firestore:"name"`
	Description string      `firestore:"description"`
	Email       string      `firestore:"email"`
	SocialMedia SocialMedia `firestore:"social_media"`
	Phone       string      `firestore:"phone"`
	About       string      `firestore:"about"`
}

type ProfileCreate struct {
	Name        string      `firestore:"name"`
	Description string      `firestore:"description"`
	Email       string      `firestore:"email"`
	SocialMedia SocialMedia `firestore:"social_media"`
	Phone       string      `firestore:"phone"`
	About       string      `firestore:"about"`
}

type SocialMedia struct {
	LinkedIn  string `firestore:"linked_in"`
	Instagram string `firestore:"instagram"`
	GitHub    string `firestore:"github"`
}
