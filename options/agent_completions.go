package options

import (
	"github.com/danilsolovyov/go-ragflow/parameters"
	"github.com/danilsolovyov/go-ragflow/utils"
)

type AgentCompletionsOptions struct {
	AgentID   string         `json:"-"`
	Question  string         `json:"question"`
	Stream    *bool          `json:"stream"`
	SessionID string         `json:"session_id"` // SessionID if not set, a new session will be created
	UserID    string         `json:"user_id"`
	SyncDSL   *bool          `json:"sync_dsl"`
	Begin     map[string]any `json:"-"` // Begin parameters for the session
	Other     map[string]any `json:"-"`
}

func DefaultAgentCompletionsOptions() *AgentCompletionsOptions {
	return &AgentCompletionsOptions{
		Stream:  utils.BoolPtr(false),
		SyncDSL: utils.BoolPtr(false),
	}
}

func (o *AgentCompletionsOptions) Parameters() []parameters.Parameter {
	params := []parameters.Parameter{
		parameters.NewPathParameter("agent_id", o.AgentID),
		parameters.NewBodyParameter("question", o.Question),
		parameters.NewBodyParameter("stream", o.Stream),
		parameters.NewBodyParameter("session_id", o.SessionID),
		parameters.NewBodyParameter("user_id", o.UserID),
		parameters.NewBodyParameter("sync_dsl", o.SyncDSL),
	}

	for key, value := range o.Other {
		params = append(params, parameters.NewBodyParameter(key, value))
	}

	if o.SessionID == "" {
		for key, value := range o.Begin {
			params = append(params, parameters.NewBodyParameter(key, value))
		}
	}

	return params
}

func (o *AgentCompletionsOptions) Merge(other *AgentCompletionsOptions) *AgentCompletionsOptions {
	if other == nil {
		return o
	}
	if other.Question != "" {
		o.Question = other.Question
	}
	if other.Stream != nil {
		o.Stream = other.Stream
	}
	if other.SessionID != "" {
		o.SessionID = other.SessionID
	}
	if other.UserID != "" {
		o.UserID = other.UserID
	}
	if other.SyncDSL != nil {
		o.SyncDSL = other.SyncDSL
	}

	return o
}

func (o *AgentCompletionsOptions) SetAgentID(agentID string) *AgentCompletionsOptions {
	o.AgentID = agentID
	return o
}

func (o *AgentCompletionsOptions) SetQuestion(question string) *AgentCompletionsOptions {
	o.Question = question
	return o
}

func (o *AgentCompletionsOptions) SetStream(stream bool) *AgentCompletionsOptions {
	o.Stream = utils.BoolPtr(stream)
	return o
}

func (o *AgentCompletionsOptions) SetSessionID(sessionID string) *AgentCompletionsOptions {
	o.SessionID = sessionID
	return o
}

func (o *AgentCompletionsOptions) SetUserID(userID string) *AgentCompletionsOptions {
	o.UserID = userID
	return o
}

func (o *AgentCompletionsOptions) SetSyncDSL(syncDSL bool) *AgentCompletionsOptions {
	o.SyncDSL = utils.BoolPtr(syncDSL)
	return o
}

func (o *AgentCompletionsOptions) SetOther(other map[string]any) *AgentCompletionsOptions {
	o.Other = other
	return o
}
