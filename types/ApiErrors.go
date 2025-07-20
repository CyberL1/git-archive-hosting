package types

const (
	ApiErrorCodeInvalidRequestBody     = "INVALID_REQUEST_BODY"
	ApiErrorCodeRepositoryUrlRequired  = "REQUIRED_REPOSITORY_URL"
	ApiErrorCodeRepositoryUrlBadSchema = "REPOSITORY_URL_BAD_SCHEMA"
	ApiErrorCodeRepositoryCloneFailed  = "REPOSITORY_CLONE_FAILED"
)

const (
	ApiErrorMessageInvalidRequestBody = "Invalid request body"
	ApiErrorMessageRepositoryUrlRequired  = "Repository URL is required"
	ApiErrorMessageRepositoryUrlBadSchema = "Repository URL must start with http:// or https://"
	ApiErrorMessageRepositoryCloneFailed = "Failed to clone the repository"
)
