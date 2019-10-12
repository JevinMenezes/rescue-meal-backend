package login

// Beer defines the properties of a beer to be added
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
