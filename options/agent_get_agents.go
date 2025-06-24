package options

import (
	"strconv"

	"github.com/danilsolovyov/go-ragflow/parameters"
	"github.com/danilsolovyov/go-ragflow/utils"
)

type GetAgentsOptions struct {
	Page     int    // Specifies the page on which the agents will be displayed. Defaults to 1.
	PageSize int    // The number of agents on each page. Defaults to 30.
	OrderBy  string // The field by which agents should be sorted. Available options:create_time (default); update_time
	Desc     *bool  // Indicates whether the retrieved agents should be sorted in descending order. Defaults to true.
	ID       string // The ID of the agent to retrieve.
	Name     string // The name of the agent to retrieve.
}

func DefaultGetAgentsOptions() *GetAgentsOptions {
	return &GetAgentsOptions{
		Page:     1,
		PageSize: 30,
		OrderBy:  "create_time",
		Desc:     utils.BoolPtr(true),
	}
}

func (o *GetAgentsOptions) Parameters() []parameters.Parameter {
	return []parameters.Parameter{
		parameters.NewQueryParameter("page", strconv.Itoa(o.Page)),
		parameters.NewQueryParameter("page_size", strconv.Itoa(o.PageSize)),
		parameters.NewQueryParameter("order_by", o.OrderBy),
		parameters.NewQueryParameter("desc", strconv.FormatBool(*o.Desc)),
		parameters.NewQueryParameter("id", o.ID),
		parameters.NewQueryParameter("name", o.Name),
	}
}

func (o *GetAgentsOptions) Merge(other *GetAgentsOptions) *GetAgentsOptions {
	if other == nil {
		return o
	}
	if other.Page != 0 {
		o.Page = other.Page
	}
	if other.PageSize != 0 {
		o.PageSize = other.PageSize
	}
	if other.OrderBy != "" {
		o.OrderBy = other.OrderBy
	}
	if other.Desc != nil {
		o.Desc = other.Desc
	}
	if other.ID != "" {
		o.ID = other.ID
	}
	if other.Name != "" {
		o.Name = other.Name
	}

	return o
}

func (o *GetAgentsOptions) SetPage(page int) *GetAgentsOptions {
	o.Page = page
	return o
}

func (o *GetAgentsOptions) SetPageSize(pageSize int) *GetAgentsOptions {
	o.PageSize = pageSize
	return o
}

func (o *GetAgentsOptions) SetOrderBy(orderBy string) *GetAgentsOptions {
	o.OrderBy = orderBy
	return o
}

func (o *GetAgentsOptions) SetDesc(desc bool) *GetAgentsOptions {
	o.Desc = utils.BoolPtr(desc)
	return o
}

func (o *GetAgentsOptions) SetID(id string) *GetAgentsOptions {
	o.ID = id
	return o
}

func (o *GetAgentsOptions) SetName(name string) *GetAgentsOptions {
	o.Name = name
	return o
}
