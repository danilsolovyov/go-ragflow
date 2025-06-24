package options

import "github.com/danilsolovyov/go-ragflow/parameters"

type CreateAgentSessionOptions struct {
	AgentID string         `json:"-"`
	UserID  string         `json:"user_id"`
	Begin   map[string]any `json:"-"` // Begin parameters for the session
}

func DefaultCreateAgentSessionOptions() *CreateAgentSessionOptions {
	return &CreateAgentSessionOptions{}
}

func (o *CreateAgentSessionOptions) Parameters() []parameters.Parameter {
	params := []parameters.Parameter{
		parameters.NewPathParameter("agent_id", o.AgentID),
		parameters.NewQueryParameter("user_id", o.UserID),
	}

	for key, value := range o.Begin {
		params = append(params, parameters.NewBodyParameter(key, value))
	}

	return params
}

func (o *CreateAgentSessionOptions) Merge(other *CreateAgentSessionOptions) *CreateAgentSessionOptions {
	if other == nil {
		return o
	}

	o.UserID = other.UserID
	o.Begin = other.Begin

	return o
}

func (o *CreateAgentSessionOptions) SetAgentID(agentID string) *CreateAgentSessionOptions {
	o.AgentID = agentID
	return o
}

func (o *CreateAgentSessionOptions) SetUserID(userID string) *CreateAgentSessionOptions {
	o.UserID = userID
	return o
}

func (o *CreateAgentSessionOptions) SetBegin(begin map[string]any) *CreateAgentSessionOptions {
	o.Begin = begin
	return o
}
