package authmodel

// Account struct containing account data
type Account struct {
	ID            int64  `json:"id" db:"id"`
	NIK           string `json:"-" db:"nik"`
	NIKIndex      string `json:"-" db:"nik_index"`
	Password      string `json:"-" db:"password"`
	PIN           string `json:"-" db:"pin"`
	Fullname      string `json:"fullname" db:"fullname"`
	LegalName     string `json:"legal_name" db:"legal_name"`
	DateOfBirth   string `json:"-" db:"date_of_birth"`
	PlaceOfBirth  string `json:"-" db:"place_of_birth"`
	Salary        int64  `json:"-" db:"salary"`
	IdentityPhoto string `json:"-" db:"identity_photo"`
	Photo         string `json:"-" db:"photo"`
}

// TokenData struct containing token data
type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
