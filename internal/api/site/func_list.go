package site

import (
	"github.com/ccHR0305/webstack-go/internal/code"
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
	"github.com/ccHR0305/webstack-go/internal/services/site"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

type listRequest struct {
	Page              int64  `form:"page,default=1"`        // 第几页
	PageSize          int64  `form:"page_size,default=20" ` // 每页显示条数
	BusinessKey       string `form:"business_key"`          // 调用方key
	BusinessSecret    string `form:"business_secret"`       // 调用方secret
	BusinessDeveloper string `form:"business_developer"`    // 调用方对接人
	Remark            string `form:"remark"`                // 备注
}
type listData struct {
	Id          int    `json:"id"`          // ID
	Thumb       string `json:"thumb"`       // 网站 logo
	Title       string `json:"title"`       // 名称简介
	Url         string `json:"url"`         // 链接
	Category    string `json:"category"`    // 分类
	CategoryId  int64  `json:"category_id"` // 分类id
	Description string `json:"description"` // 描述
	IsUsed      int64  `json:"is_used"`     // 是否启用
	CreatedAt   string `json:"created_at"`  // 创建时间
	UpdatedAt   string `json:"updated_at"`  // 更新时间
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int64 `json:"current_page"`
		PerPageCount int64 `json:"per_page_count"`
	} `json:"pagination"`
}

// List 网站列表
// @Summary 网站列表
// @Description 网站列表
// @Tags API.site
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/site [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := c.ShouldBind(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		searchData := new(site.SearchData)
		searchData.Page = req.Page
		searchData.PageSize = req.PageSize
		searchData.BusinessKey = req.BusinessKey
		searchData.BusinessSecret = req.BusinessSecret
		searchData.Remark = req.Remark
		searchData.Search = strings.TrimSpace(c.Query("search"))

		sites, err := h.siteService.PageList(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedListError,
				code.Text(code.AuthorizedListError)).WithError(err),
			)
			return
		}

		total, err := h.siteService.PageListCount(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SiteListError,
				code.Text(code.SiteListError)).WithError(err),
			)
			return
		}

		res.List = make([]listData, len(sites))
		for i, sit := range sites {
			_, err := h.hashids.HashidsEncode([]int{cast.ToInt(sit.ID)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			res.List[i] = listData{
				Id:          cast.ToInt(sit.ID),
				Thumb:       sit.Thumb,
				Title:       sit.Title,
				Url:         sit.URL,
				Category:    sit.Category.Title,
				CategoryId:  sit.Category.ID,
				Description: sit.Description,
				IsUsed:      sit.IsUsed,
				CreatedAt:   sit.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   sit.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}

		res.Pagination.Total = total
		res.Pagination.PerPageCount = req.PageSize
		res.Pagination.CurrentPage = req.Page

		c.Payload(res)
	}
}
