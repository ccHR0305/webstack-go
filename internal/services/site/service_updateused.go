package site

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

func (s *service) UpdateUsed(ctx core.Context, id, used int64) (err error) {
	if _, err = query.Site.WithContext(ctx.RequestContext()).
		Where(query.Site.ID.Eq(id)).
		Update(query.Site.IsUsed, used); err != nil {
		return err
	}

	return
}
