package goragflow

type SessionMessage struct {
	Content   string  `json:"content"`
	CreatedAt float64 `json:"created_at"`
	ID        string  `json:"id"`
	Reference []any   `json:"reference"`
	Role      string  `json:"role"`
}
