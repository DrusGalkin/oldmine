package dto

type ResponseUser struct {
	ID        int
	Name      string
	Email     string
	CreatedAt string
	Auth      bool
	Pay       bool
	Admin     bool
}
