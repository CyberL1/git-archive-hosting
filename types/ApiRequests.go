package types

type ApiRepositoryImportRequest struct {
	RepositoryUrl string `json:"repositoryUrl"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}
