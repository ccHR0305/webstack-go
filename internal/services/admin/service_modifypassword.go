package admin

import (
	"github.com/ccHR0305/webstack-go/configs"
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/pkg/password"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/query"
	"github.com/ccHR0305/webstack-go/internal/repository/redis"
)

func (s *service) ModifyPassword(ctx core.Context, id int64, newPassword string) (err error) {

	if _, err = query.Admin.WithContext(ctx.RequestContext()).
		Where(query.Admin.ID.Eq(id)).
		UpdateColumnSimple(
			query.Admin.Password.Value(password.GeneratePassword(newPassword)),
			query.Admin.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		); err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), redis.WithTrace(ctx.Trace()))
	return
}
