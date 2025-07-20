package types

type ApiErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type ApiRepositoryResponse struct {
	Id          int64  `json:"id"`
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	CreatedAt   string `json:"createdAt"`
	OriginalUrl string `json:"originalUrl"`
}
