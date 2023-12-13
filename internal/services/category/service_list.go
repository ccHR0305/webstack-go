package category

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

type SearchData struct {
	Pid int32 // 父类ID
}

func (s *service) List(ctx core.Context) (categories []*model.Category, err error) {
	categories, err = query.Category.WithContext(ctx.RequestContext()).
		Order(query.Category.Sort).
		Find()
	if err != nil {
		return nil, err
	}

	return
}
