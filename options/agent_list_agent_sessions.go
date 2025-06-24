package options

import (
	"strconv"

	"github.com/danilsovyov/go-ragflow/parameters"
	"github.com/danilsovyov/go-ragflow/utils"
)

// ListAgentSessionsOptions represents the options for listing sessions.
type ListAgentSessionsOptions struct {
	AgentID  string // The ID of the agent.
	Page     int    // Specifies the page on which the sessions will be displayed. Defaults to 1.
	PageSize int    // The number of sessions on each page. Defaults to 30.
	OrderBy  string // The field by which sessions should be sorted. Available options:create_time (default); update_time
	Desc     *bool  // Indicates whether the retrieved sessions should be sorted in descending order. Defaults to true.
	ID       string // The ID of the agent session to retrieve.
	UserID   string // The optional user-defined ID passed in when creating session.
	DSL      *bool  // Indicates whether to include the dsl field of the sessions in the response. Defaults to true.
}

// DefaultListAgentSessionsOptions returns a new ListAgentSessionsOptions with default values.
func DefaultListAgentSessionsOptions() *ListAgentSessionsOptions {
	return &ListAgentSessionsOptions{
		Page:     1,
		PageSize: 30,
		OrderBy:  "create_time",
		Desc:     utils.BoolPtr(true),
		ID:       "",
		UserID:   "",
		DSL:      utils.BoolPtr(true),
	}
}

func (o *ListAgentSessionsOptions) Parameters() []parameters.Parameter {
	return []parameters.Parameter{
		parameters.NewPathParameter("agent_id", o.AgentID),
		parameters.NewQueryParameter("page", strconv.Itoa(o.Page)),
		parameters.NewQueryParameter("page_size", strconv.Itoa(o.PageSize)),
		parameters.NewQueryParameter("order_by", o.OrderBy),
		parameters.NewQueryParameter("desc", strconv.FormatBool(*o.Desc)),
		parameters.NewQueryParameter("id", o.ID),
		parameters.NewQueryParameter("user_id", o.UserID),
		parameters.NewQueryParameter("dsl", strconv.FormatBool(*o.DSL)),
	}
}

// Merge merges the given ListAgentSessionsOptions into the current instance.
func (o *ListAgentSessionsOptions) Merge(other *ListAgentSessionsOptions) *ListAgentSessionsOptions {
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
	if other.UserID != "" {
		o.UserID = other.UserID
	}
	if other.DSL != nil {
		o.DSL = other.DSL
	}

	return o
}

// SetAgentID sets the agentID field of ListAgentSessionsOptions.
//
// @param agentID The agentID value to set.
func (o *ListAgentSessionsOptions) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// SetPage sets the page field of ListAgentSessionsOptions.
//
// @param page The page value to set. Must be greater than 0.
func (o *ListAgentSessionsOptions) SetPage(page int) {
	if page < 1 {
		page = 1
	}

	o.Page = page
}

// SetPageSize sets the pageSize field of ListAgentSessionsOptions.
//
// @param pageSize The pageSize value to set. Must be greater than 0.
func (o *ListAgentSessionsOptions) SetPageSize(pageSize int) {
	if pageSize < 1 {
		pageSize = 1
	}

	o.PageSize = pageSize
}

// SetOrderBy sets the orderBy field of ListAgentSessionsOptions.
//
// @param orderBy The orderBy value to set. Must be one of "create_time", "update_time".
func (o *ListAgentSessionsOptions) SetOrderBy(orderBy string) {
	allowedOrderBy := []string{"create_time", "update_time"}
	for _, v := range allowedOrderBy {
		if v == orderBy {
			o.OrderBy = orderBy
			return
		}
	}

	o.OrderBy = orderBy
}

// SetDesc sets the desc field of ListAgentSessionsOptions.
//
// @param desc The desc value to set.
func (o *ListAgentSessionsOptions) SetDesc(desc bool) {
	o.Desc = &desc
}

// SetID sets the ID field of ListAgentSessionsOptions.
//
// @param id The ID value to set.
func (o *ListAgentSessionsOptions) SetID(id string) {
	o.ID = id
}

// SetUserID sets the user ID for the ListAgentSessionsOptions.
//
// @param userID The user ID to be set.
func (o *ListAgentSessionsOptions) SetUserID(userID string) {
	o.UserID = userID
}

// SetDSL sets the DSL field of ListAgentSessionsOptions.
//
// @param dsl The DSL value to set.
func (o *ListAgentSessionsOptions) SetDSL(dsl bool) {
	o.DSL = &dsl
}
