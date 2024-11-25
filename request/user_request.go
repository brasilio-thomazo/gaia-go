package request

type UserRequest struct {
	GroupID         int    `json:"group_id"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	JobTitle        string `json:"job_title"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Locked          bool   `json:"locked"`
}
