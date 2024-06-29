package model

type KYCRequest struct {
	FileURL string `json:"file_url"`
	IdType  string `json:"id_type"`
	Id      string `json:"id"`
	UserId  string `json:"user_id"`
}

type KYCAction struct {
	UserId  string `json:"user_id"`
	Approve bool   `json:"approve"`
}
