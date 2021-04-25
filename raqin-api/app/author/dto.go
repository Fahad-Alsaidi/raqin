package author

type NewAuthorRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type NewAuthorResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
