package menu

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

type SearchOneData struct {
	Id     int64 // 用户ID
	IsUsed int64 // 是否启用 1:是  -1:否
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (menu *model.Menu, err error) {
	iMenuDo := query.Menu.WithContext(ctx.RequestContext())
	iMenuDo = iMenuDo.Where(query.Menu.IsDeleted.Eq(-1))
	if searchOneData.Id != 0 {
		iMenuDo = iMenuDo.Where(query.Menu.ID.Eq(searchOneData.Id))
	}
	if searchOneData.IsUsed != 0 {
		iMenuDo = iMenuDo.Where(query.Menu.IsUsed.Eq(searchOneData.IsUsed))
	}
	menu, err = iMenuDo.First()
	if err != nil {
		return nil, err
	}

	return
}
