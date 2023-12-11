package category

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

func (s *service) Detail(ctx core.Context, id int64) (category *model.Category, err error) {
	if category, err = query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.ID.Eq(id)).
		First(); err != nil {
		return nil, err
	}

	return
}
