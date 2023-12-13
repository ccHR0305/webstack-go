package site

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
	"github.com/ccHR0305/webstack-go/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	CategoryList(ctx core.Context) (categories []*model.Category, err error)
	PageList(ctx core.Context, searchData *SearchData) (sites []*model.Site, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id, used int64) (err error)
	Delete(ctx core.Context, id int64) (err error)
	Create(ctx core.Context, sitesData []*CreateSiteData) (successCount, failCount int64)
	CategorySite(ctx core.Context) (categorySites []*CategorySite, err error)
	UpdateSite(ctx core.Context, updateSite *UpdateSiteRequest) (err error)
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
