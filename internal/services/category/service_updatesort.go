package category

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

func (s *service) UpdateSort(ctx core.Context, id, sort int64) (err error) {
	if _, err = query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.ID.Eq(id)).
		Update(query.Category.Sort, sort); err != nil {
		return err
	}

	return

}
