package github

//Repository is a representation of a Github repository
type Repository struct {
	ID          int64
	Name        string
	Language    string
	Description string
}
