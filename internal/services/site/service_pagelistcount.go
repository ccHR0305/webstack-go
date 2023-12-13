package site

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

func (s *service) PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error) {

	iSiteDo := query.Site.WithContext(ctx.RequestContext())
	if searchData.Search != "" {
		iSiteDo = iSiteDo.Where(query.Site.Title.Like("%" + searchData.Search + "%"))
	}
	if total, err = iSiteDo.Count(); err != nil {
		return 0, err
	} else {
		return
	}
}
