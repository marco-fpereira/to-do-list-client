package dto

type AccountDetailsDTO struct {
	Username string `json:"username" copier:"username"`
	Password string `json:"password" copier:"password"`
	UserId   string `json:"userId,omitempty" copier:"UserId"`
}
