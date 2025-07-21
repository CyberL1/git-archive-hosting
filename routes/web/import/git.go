package importRoutes

import (
	"garg/utils"

	"github.com/gin-gonic/gin"
)

func ImportGitRepo(c *gin.Context) {
	utils.RenderPage(c.Writer, "import/git", map[string]interface{}{
		"Title": "Import a Repository from Git",
	})
}
