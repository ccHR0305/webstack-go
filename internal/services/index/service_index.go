package index

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

func (s *service) Index(ctx core.Context) (sites []*model.Site, err error) {
	sites, err = query.Site.WithContext(ctx.RequestContext()).Find()
	if err != nil {
		return nil, err
	}
	return
}
