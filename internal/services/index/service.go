package index

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
	"github.com/ccHR0305/webstack-go/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Index(ctx core.Context) (sites []*model.Site, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
