package profiledto

type UpdateProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"FullName" form:"FullName" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
}