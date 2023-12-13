package admin

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

func (s *service) PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error) {

	iAdminDo := query.Admin.WithContext(ctx.RequestContext())
	iAdminDo = iAdminDo.Where(query.Admin.IsDeleted.Eq(-1))

	if searchData.Username != "" {
		iAdminDo.Where(query.Admin.Username.Eq(searchData.Username))
	}

	if searchData.Nickname != "" {
		iAdminDo.Where(query.Admin.Nickname.Eq(searchData.Nickname))
	}

	if searchData.Mobile != "" {
		iAdminDo.Where(query.Admin.Mobile.Eq(searchData.Mobile))
	}

	total, err = iAdminDo.Count()
	if err != nil {
		return 0, err
	}

	return
}
