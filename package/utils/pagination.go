package utils

import (
	"math"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit    int64       `json:"limit"`
	Page     int64       `json:"page"`
	SortBy   string      `json:"sort_by"`
	SortDesc int64       `json:"sort_desc"`
	Count    int64       `json:"count"`
	Pages    int64       `json:"pages"`
	Records  interface{} `json:"records"`
}

func NewPagination(limit interface{}, page interface{}, sort_by interface{}, sort_desc interface{}, with_total bool) *Pagination {
	p := &Pagination{
		Limit: 10,
		Page:  1,
	}
	if !with_total {
		p.Pages = -1
	}
	switch v := limit.(type) {
	case int64:
		p.Limit = v
	case int:
	case float64:
		p.Limit = int64(v)
	case string:
		v_int, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			p.Limit = int64(v_int)
		}
	}

	switch v := page.(type) {
	case int64:
		p.Page = v
	case int:
	case float64:
		p.Page = int64(v)
	case string:
		v_int, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			p.Page = int64(v_int)
		}
	}

	switch v := sort_by.(type) {
	case string:
		p.SortBy = v
	}
	switch v := sort_desc.(type) {
	case int64:
		p.SortDesc = v
	}
	return p
}

func (p *Pagination) GetOffset() int {
	return int((p.GetPage() - 1) * p.GetLimit())
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return int(p.Limit)
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return int(p.Page)
}

func (p *Pagination) GetSort() string {
	if p.SortBy == "" {
		p.SortBy = ""
	} else {
		if p.SortDesc == 1 {
			p.SortBy = p.SortBy + " DESC"
		} else {
			p.SortBy = p.SortBy + " ASC"
		}
	}
	return p.SortBy
}
func GetScopePagination(value interface{}, pagination *Pagination, query *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if pagination.Pages != -1 && query != nil {
		var count int64
		query.Session(&gorm.Session{}).Count(&count)
		pagination.Count = count
		pages := int64(math.Ceil(float64(count) / float64(pagination.GetLimit())))
		pagination.Pages = pages
	} else {
		pagination.Pages = 0
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
