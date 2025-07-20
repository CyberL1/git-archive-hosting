package importRoutes

import (
	"garg/utils"
	"net/http"
)

func ImportGitRepo(w http.ResponseWriter, r *http.Request) {
	utils.RenderPage(w, "import/git", map[string]interface{}{
		"Title": "Import a Repository from Git",
	})
}
