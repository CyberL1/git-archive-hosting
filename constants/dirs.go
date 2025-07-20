package constants

import "path/filepath"

var (
	DataDir         = filepath.Join(".", "data")
	RepositoriesDir = filepath.Join(DataDir, "repositories")
)
