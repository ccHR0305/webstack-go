package index

import (
	"github.com/ccHR0305/webstack-go/configs"
	"github.com/ccHR0305/webstack-go/internal/code"
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/pkg/file"
	"github.com/ccHR0305/webstack-go/internal/services/category"
	"github.com/ccHR0305/webstack-go/internal/services/site"
	"net/http"
)

type indexResponse struct {
	CategoryTree  []*category.TreeNode
	CategorySites []*site.CategorySite
}

func (h *handler) Index() core.HandlerFunc {
	return func(c core.Context) {

		if _, ok := file.IsExists(configs.ProjectInstallMark); !ok {
			c.Redirect(http.StatusTemporaryRedirect, "/install")
		}

		categoryTree, err := h.categoryService.Tree(c)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}

		categorySites, err := h.siteService.CategorySite(c)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}

		response := indexResponse{
			CategoryTree:  categoryTree,
			CategorySites: categorySites,
		}

		c.HTML("index", response)
	}
}
