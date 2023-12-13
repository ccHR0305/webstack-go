package category

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
)

type UpdateCategory struct {
	Name string `json:"name"` // 菜单名称
	Icon string `json:"icon"` // 图标
}

func (s *service) Modify(ctx core.Context, id int64, updateCategory *UpdateCategory) (err error) {
	if _, err = query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.ID.Eq(id)).
		Updates(model.Category{Title: updateCategory.Name, Icon: updateCategory.Icon}); err != nil {
		return
	}

	return
}
