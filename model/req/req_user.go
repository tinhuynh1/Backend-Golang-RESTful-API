package req

type ReqUpdateUser struct {
	FullName string `json:"fullName,onmitempty" validate:"required"`
	Email    string `json:"email,onmitempty" validate:"required"`
}
