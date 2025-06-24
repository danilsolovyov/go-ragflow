package goragflow

// Completions is the response from the /completions endpoint.
type Completions struct {
	Answer    string             `json:"answer"`
	ID        string             `json:"id"`
	Param     []CompletionsParam `json:"param"`
	Reference Reference          `json:"reference"`
	SessionID string             `json:"session_id"`
}

type CompletionsParam struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

type Reference any // TODO: implement me - it's maybe files
