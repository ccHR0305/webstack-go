package authorized

import (
	"github.com/ccHR0305/webstack-go/configs"
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
	"github.com/ccHR0305/webstack-go/internal/repository/redis"

	"gorm.io/gorm"
)

func (s *service) DeleteAPI(ctx core.Context, id int64) (err error) {
	// 先查询 id 是否存在
	authorizedAPI, err := query.AuthorizedAPI.WithContext(ctx.RequestContext()).
		Where(query.AuthorizedAPI.IsDeleted.Eq(-1)).
		Where(query.AuthorizedAPI.ID.Eq(id)).
		First()
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	_, err = query.AuthorizedAPI.WithContext(ctx.RequestContext()).
		Where(query.AuthorizedAPI.ID.Eq(id)).
		UpdateColumnSimple(
			query.AuthorizedAPI.IsDeleted.Value(1),
			query.AuthorizedAPI.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		)
	if err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixSignature+authorizedAPI.BusinessKey, redis.WithTrace(ctx.Trace()))
	return
}
