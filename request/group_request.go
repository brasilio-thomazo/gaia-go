package request

type GroupRequest struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	Locked      bool     `json:"locked"`
}
