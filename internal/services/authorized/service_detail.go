package authorized

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

func (s *service) Detail(ctx core.Context, id int64) (authorizedInfo *model.Authorized, err error) {
	authorizedInfo, err = query.Authorized.WithContext(ctx.RequestContext()).
		Where(query.Authorized.IsDeleted.Eq(-1)).
		Where(query.Authorized.ID.Eq(id)).
		First()
	if err != nil {
		return nil, err
	}

	return
}
