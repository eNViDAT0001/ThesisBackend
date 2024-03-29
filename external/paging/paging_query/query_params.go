package paging_query

import (
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_params"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPagingParams(cc *gin.Context, filter paging_params.EntityFilter) (paginator paging.ParamsInput, err error) {
	if err = cc.BindQuery(&paginator); err != nil {
		return paginator, err
	}

	search := cc.QueryArray("search[]")
	fields := cc.QueryArray("fields[]")
	sort := cc.QueryArray("sorts[]")
	compare := cc.QueryArray("compare[]")

	paginator.Filter = paging_params.NewFilterBuilder().
		WithSearch(search).
		WithFields(fields).
		WithSorts(sort).
		WithCompare(compare).
		Build()

	inValidField, val := paging_params.ValidateFilter(paginator.Filter, filter)
	if len(inValidField) > 0 {
		return paginator, request.NewBadRequestError(inValidField, val, "invalid key and value")
	}
	return paginator, err
}

func SetCountListPagingQuery(input *paging.ParamsInput, tableName string, query *gorm.DB) {

	if input.Type == paging.CursorPaging && !input.Infinity {
		sort := ">"
		queryString := ""
		if input.Filter.GetSort() != nil {
			val, ok := (*input.Filter.GetSort())["id"]
			if ok {
				if val == "DESC" {
					if input.Current() == 0 {
						queryString = "Unknown"
					}
					sort = "<"
				}
			}
			if len(queryString) < 1 {
				queryString = fmt.Sprintf("%s.id %s ?", tableName, sort)
			}
		}
		if queryString != "Unknown" {
			query = query.Where(fmt.Sprintf("%s.id %s ?", tableName, sort), input.Current())
		}
	}

	if input.Filter.GetFields() != nil {
		for k, v := range *input.Filter.GetFields() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` = ?", tableName, k), v)
		}
	}

	if input.Filter.GetSearch() != nil {
		for k, v := range *input.Filter.GetSearch() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` LIKE ?", tableName, k), "%"+v+"%")
		}
	}

	if input.Filter.GetSort() != nil {
		for k, v := range *input.Filter.GetSort() {
			sort := "ASC"
			if v == "DESC" {
				sort = v
			}
			query = query.Order(fmt.Sprintf(`%s.%s %s`, tableName, k, sort))
		}
	}

	if input.Filter.GetCompare() != nil {
		for k, v := range *input.Filter.GetCompare() {
			column, condition := GetColumnAndCondition(k)
			query = query.Where(fmt.Sprintf("`%s`.`%s` %s ?", tableName, column, condition), v)
		}
	}
}
func SetPagingQuery(input *paging.ParamsInput, tableName string, query *gorm.DB) {

	query = query.Limit(input.PerPage())
	if input.Type == paging.CursorPaging && !input.Infinity {
		sort := ">"
		queryString := ""
		if input.Filter.GetSort() != nil {
			val, ok := (*input.Filter.GetSort())["id"]
			if ok {
				if val == "DESC" {
					if input.Current() == 0 {
						queryString = "Unknown"
					}
					sort = "<"
				}
			}
			if len(queryString) < 1 {
				queryString = fmt.Sprintf("%s.id %s ?", tableName, sort)
			}
		}
		if queryString != "Unknown" {
			if input.MarkerDefinition != nil {
				query = query.Where(fmt.Sprintf("%s.%s %s ?", tableName, *input.MarkerDefinition, sort), input.Current())
			} else {
				query = query.Where(fmt.Sprintf("%s.id %s ?", tableName, sort), input.Current())
			}
		}
	} else {
		offset := paging.Offset(input.Current(), input.PerPage())
		query = query.Offset(offset)
	}

	if input.Filter.GetFields() != nil {
		for k, v := range *input.Filter.GetFields() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` = ?", tableName, k), v)
		}
	}

	if input.Filter.GetSearch() != nil {
		for k, v := range *input.Filter.GetSearch() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` LIKE ?", tableName, k), "%"+v+"%")
		}
	}

	if input.Filter.GetSort() != nil {
		for k, v := range *input.Filter.GetSort() {
			sort := "ASC"
			if v == "DESC" {
				sort = v
			}
			query = query.Order(fmt.Sprintf(`%s.%s %s`, tableName, k, sort))
		}
	}

	if input.Filter.GetCompare() != nil {
		for k, v := range *input.Filter.GetCompare() {
			column, condition := GetColumnAndCondition(k)
			query = query.Where(fmt.Sprintf("`%s`.`%s` %s ?", tableName, column, condition), v)
		}
	}
}

func GetColumnAndCondition(v string) (string, string) {
	var column string
	var condition string

	index := len(v) - 1
	for index >= 0 {
		if v[index] == '_' {
			break
		}
		index--
	}

	if index == -1 {
		return "", ""
	}

	column = v[:index]
	condition = v[index+1:]

	return column, condition
}
