package options

import (
	"github.com/danilsolovyov/go-ragflow/parameters"
	"github.com/danilsolovyov/go-ragflow/utils"
)

type CompletionsOptions struct {
	AgentID   string         `json:"-"`          // Required one of <AgentID, ChatID>. Default: empty.
	ChatID    string         `json:"-"`          // Required one of <AgentID, ChatID>. Default: empty.
	Question  string         `json:"question"`   // Required. Default: empty
	Stream    *bool          `json:"stream"`     // Indicates whether to output responses in a streaming way. Default: true
	SessionID string         `json:"session_id"` // The ID of session. If it is not provided, a new session will be generated. Default: empty
	UserID    string         `json:"user_id"`    // The optional user-defined ID. Valid only when no session_id is provided. Default: empty
	SyncDSL   *bool          `json:"sync_dsl"`   // Whether to synchronize the changes to existing sessions when an agent is modified. Default: false
	Begin     map[string]any `json:"-"`          // Begin parameters for the session. Default: empty
	Other     map[string]any `json:"-"`          // Other parameters for the session. Default: empty
}

func DefaultCompletionsOptions() *CompletionsOptions {
	return &CompletionsOptions{
		Stream:  utils.BoolPtr(true),
		SyncDSL: utils.BoolPtr(false),
	}
}

func (o *CompletionsOptions) Parameters() []parameters.Parameter {
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

func (o *CompletionsOptions) Merge(other *CompletionsOptions) *CompletionsOptions {
	if other == nil {
		return o
	}
	if other.AgentID != "" {
		o.AgentID = other.AgentID
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
	if other.Begin != nil {
		o.Begin = other.Begin
	}
	if other.Other != nil {
		o.Other = other.Other
	}

	return o
}

func (o *CompletionsOptions) SetAgentID(agentID string) *CompletionsOptions {
	o.AgentID = agentID
	return o
}

func (o *CompletionsOptions) SetQuestion(question string) *CompletionsOptions {
	o.Question = question
	return o
}

func (o *CompletionsOptions) SetStream(stream bool) *CompletionsOptions {
	o.Stream = utils.BoolPtr(stream)
	return o
}

func (o *CompletionsOptions) SetSessionID(sessionID string) *CompletionsOptions {
	o.SessionID = sessionID
	return o
}

func (o *CompletionsOptions) SetUserID(userID string) *CompletionsOptions {
	o.UserID = userID
	return o
}

func (o *CompletionsOptions) SetSyncDSL(syncDSL bool) *CompletionsOptions {
	o.SyncDSL = utils.BoolPtr(syncDSL)
	return o
}

func (o *CompletionsOptions) SetOther(other map[string]any) *CompletionsOptions {
	o.Other = other
	return o
}
