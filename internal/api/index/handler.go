package index

import (
	"github.com/ccHR0305/webstack-go/configs"
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/pkg/hash"
	"github.com/ccHR0305/webstack-go/internal/repository/mysql"
	"github.com/ccHR0305/webstack-go/internal/repository/redis"
	"github.com/ccHR0305/webstack-go/internal/services/category"
	"github.com/ccHR0305/webstack-go/internal/services/index"
	"github.com/ccHR0305/webstack-go/internal/services/site"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Index 导航网站首页
	// @Tags API.admin
	// @Router / [get]
	Index() core.HandlerFunc

	// About 导航网站关于页
	// @Tags API.admin
	// @Router /about [get]
	About() core.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	cache           redis.Repo
	hashids         hash.Hash
	indexService    index.Service
	categoryService category.Service
	siteService     site.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:          logger,
		cache:           cache,
		hashids:         hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		indexService:    index.New(db, cache),
		categoryService: category.New(db, cache),
		siteService:     site.New(db, cache),
	}
}

func (h *handler) i() {}
