package category

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

func (s *service) Delete(ctx core.Context, id int64) (err error) {

	if _, err = query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.ID.Eq(id)).
		Delete(); err != nil {
		return
	}

	// 删除一级目录 id 需要删除二级分类
	if _, err = query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.ParentID.Eq(id)).
		Delete(); err != nil {
		return
	}

	return
}
