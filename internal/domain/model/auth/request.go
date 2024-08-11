package authmodel

// RegisterRequest registration request
type RegisterRequest struct {
	NIK           string `json:"nik"`
	Fullname      string `json:"fullname"`
	LegalName     string `json:"legal_name"`
	DateOfBirth   string `json:"date_of_birth"`
	PlaceOfBirth  string `json:"place_of_birth"`
	Salary        int64  `json:"salary"`
	IdentityPhoto string `json:"identity_photo"`
	Photo         string `json:"photo"`

	Password string `json:"password"`
	PIN      string `json:"pin"`
}

// ConsumerAuthRequest login request
type ConsumerAuthRequest struct {
	NIK      string `json:"nik"`
	Password string `json:"password"`
}
