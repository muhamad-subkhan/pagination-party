package profilefto

type UpdateProfileRequest struct {
	FullName string `json:"FullName" form:"FullName" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
}