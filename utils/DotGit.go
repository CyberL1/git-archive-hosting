package utils

import "strings"

func AppendDotGitExt(path string) string {
	if !strings.HasSuffix(path, ".git") {
		return path + ".git"
	}
	return path
}

func RemoveDotGitExt(path string) string {
	return strings.TrimSuffix(path, ".git")
}
