package goragflow

type DSL struct {
	Answer     []any                   `json:"answer"`
	Components map[string]DSLComponent `json:"components"`
	EmbedID    string                  `json:"embed_id"`
	Graph      DSLGraph                `json:"graph"`
	History    [][]string              `json:"history"`
	Messages   []SessionMessage        `json:"messages"`
	Path       [][]string              `json:"path"`
	Reference  []any                   `json:"reference"`
}

type DSLComponent struct {
	Downstream []string        `json:"downstream"`
	Name       string          `json:"name"`
	Upstream   []string        `json:"upstream"`
	Params     map[string]any  `json:"params"`
	Obj        DSLComponentObj `json:"obj"`
}

type DSLComponentObj struct {
	ComponentName string                 `json:"component_name"`
	Name          string                 `json:"name"`
	Type          string                 `json:"type"`
	Inputs        []DSLComponentObjInput `json:"inputs"`
	Output        DSLComponentObjOutput  `json:"output"`
	Params        map[string]any         `json:"params"`
}

type DSLComponentObjInput struct {
	ComponentID string `json:"component_id"`
	Content     string `json:"content"`
}

type DSLComponentObjOutput struct {
	ComponentID any `json:"component_id"`
	Content     any `json:"content"`
}

type Query struct {
	ComponentID string `json:"component_id"`
	Type        string `json:"type"`
}

type DSLGraph struct {
	Edges []Edge `json:"edges"`
	Nodes []Node `json:"nodes"`
}

type Edge struct {
	ID           string  `json:"id"`
	Label        *string `json:"label,omitempty"`
	Source       string  `json:"source"`
	Target       string  `json:"target"`
	MarkerEnd    string  `json:"markerEnd,omitempty"`
	SourceHandle *string `json:"sourceHandle,omitempty"`
	Style        *Style  `json:"style,omitempty"`
	TargetHandle string  `json:"targetHandle,omitempty"`
	Type         string  `json:"type,omitempty"`
	ZIndex       *int64  `json:"zIndex,omitempty"`
	Selected     *bool   `json:"selected,omitempty"`
}

type Style struct {
	Stroke      string `json:"stroke"`
	StrokeWidth int64  `json:"strokeWidth"`
}

type Node struct {
	Data             Data      `json:"data"`
	Dragging         bool      `json:"dragging"`
	Height           *int64    `json:"height,omitempty"`
	ID               string    `json:"id"`
	Measured         Measured  `json:"measured"`
	Position         Position  `json:"position"`
	PositionAbsolute *Position `json:"positionAbsolute,omitempty"`
	Selected         bool      `json:"selected"`
	SourcePosition   string    `json:"sourcePosition"`
	TargetPosition   string    `json:"targetPosition"`
	Type             string    `json:"type"`
	Width            *int64    `json:"width,omitempty"`
	DragHandle       *string   `json:"dragHandle,omitempty"`
	Resizing         *bool     `json:"resizing,omitempty"`
	Style            *Measured `json:"style,omitempty"`
}

type Data struct {
	Form  Form   `json:"form"`
	Label string `json:"label"`
	Name  string `json:"name"`
}

type Form struct {
	Prologue                 *string  `json:"prologue,omitempty"`
	QueryType                *string  `json:"query_type,omitempty"`
	TopN                     *int64   `json:"top_n,omitempty"`
	Query                    []Query  `json:"query,omitempty"`
	CategoryDescription      any      `json:"category_description,omitempty"`
	FrequencyPenaltyEnabled  *bool    `json:"frequencyPenaltyEnabled,omitempty"`
	FrequencyPenalty         *float64 `json:"frequency_penalty,omitempty"`
	LlmID                    string   `json:"llm_id,omitempty"`
	MaxTokensEnabled         *bool    `json:"maxTokensEnabled,omitempty"`
	MaxTokens                *int64   `json:"max_tokens,omitempty"`
	MessageHistoryWindowSize *int64   `json:"message_history_window_size,omitempty"`
	Parameter                string   `json:"parameter,omitempty"`
	PresencePenaltyEnabled   *bool    `json:"presencePenaltyEnabled,omitempty"`
	PresencePenalty          *float64 `json:"presence_penalty,omitempty"`
	Temperature              *float64 `json:"temperature,omitempty"`
	TemperatureEnabled       *bool    `json:"temperatureEnabled,omitempty"`
	TopPEnabled              *bool    `json:"topPEnabled,omitempty"`
	TopP                     *float64 `json:"top_p,omitempty"`
	Email                    *string  `json:"email,omitempty"`
	Channel                  *string  `json:"channel,omitempty"`
	Language                 *string  `json:"language,omitempty"`
	Lang                     *string  `json:"lang,omitempty"`
	TimePeriod               *string  `json:"time_period,omitempty"`
	Type                     *string  `json:"type,omitempty"`
	UserType                 *string  `json:"user_type,omitempty"`
	WebApikey                *string  `json:"web_apikey,omitempty"`
	Cite                     *bool    `json:"cite,omitempty"`
	Parameters               []any    `json:"parameters,omitempty"`
	Prompt                   *string  `json:"prompt,omitempty"`
	KBIDS                    []any    `json:"kb_ids,omitempty"`
	KeywordsSimilarityWeight *float64 `json:"keywords_similarity_weight,omitempty"`
	SimilarityThreshold      *float64 `json:"similarity_threshold,omitempty"`
	Text                     *string  `json:"text,omitempty"`
}

type Measured struct {
	Height int64 `json:"height"`
	Width  int64 `json:"width"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
