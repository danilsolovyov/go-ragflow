package parameters_test

import (
	"net/url"
	"testing"

	"github.com/danilsolovyov/go-ragflow/parameters"
)

func TestApplyURL(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  string
		params   []parameters.Parameter
		expected string
	}{
		{
			name:    "Add query parameter",
			baseURL: "http://example.com?test=1",
			params: []parameters.Parameter{
				parameters.NewQueryParameter("query1", "value1"),
			},
			expected: "http://example.com?query1=value1&test=1",
		},
		{
			name:    "Replace path parameter",
			baseURL: "http://example.com/users/:id",
			params: []parameters.Parameter{
				parameters.NewPathParameter("id", "123"),
			},
			expected: "http://example.com/users/123",
		},
		{
			name:    "Remove path parameter if empty",
			baseURL: "http://example.com/users/:id",
			params: []parameters.Parameter{
				parameters.NewPathParameter("id", ""),
			},
			expected: "http://example.com/users",
		},
		{
			name:    "Multiple parameters in path",
			baseURL: "http://example.com/:version/:id",
			params: []parameters.Parameter{
				parameters.NewPathParameter("version", "v1"),
				parameters.NewPathParameter("id", "456"),
			},
			expected: "http://example.com/v1/456",
		},
		{
			name:    "Query and path parameters together",
			baseURL: "http://example.com/users/:id",
			params: []parameters.Parameter{
				parameters.NewPathParameter("id", "123"),
				parameters.NewQueryParameter("page", "2"),
			},
			expected: "http://example.com/users/123?page=2",
		},
		{
			name:    "Ignore body parameters",
			baseURL: "http://example.com",
			params: []parameters.Parameter{
				parameters.NewBodyParameter("data", "test"),
			},
			expected: "http://example.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u, _ := url.Parse(test.baseURL)

			err := parameters.ApplyURL(u, test.params...)
			if err != nil {
				t.Errorf("got error, err = %s", err.Error())
			}

			// Construct expected URL
			expectedURL, _ := url.Parse(test.expected)
			got := u.String()
			want := expectedURL.String()

			if got != want {
				t.Errorf("ApplyURL() = %v, want %v", got, want)
			}
		})
	}
}

func TestCreateBody(t *testing.T) {
	tests := []struct {
		name         string
		parameters   []parameters.Parameter
		expectedBody map[string]any
	}{
		{
			name: "Single Body Parameter",
			parameters: []parameters.Parameter{
				parameters.NewBodyParameter("id", 123),
			},
			expectedBody: map[string]any{
				"id": 123,
			},
		},
		{
			name: "Mixed Parameters",
			parameters: []parameters.Parameter{
				parameters.NewBodyParameter("id", 123),
				parameters.NewQueryParameter("name", "John"),
				parameters.NewPathParameter("version", "1"),
			},
			expectedBody: map[string]any{
				"id": 123,
			},
		},
		{
			name: "Multiple Body Parameters",
			parameters: []parameters.Parameter{
				parameters.NewBodyParameter("id", 123),
				parameters.NewBodyParameter("name", "Alice"),
				parameters.NewQueryParameter("role", "admin"),
			},
			expectedBody: map[string]any{
				"id":   123,
				"name": "Alice",
			},
		},
		{
			name: "No Body Parameters",
			parameters: []parameters.Parameter{
				parameters.NewQueryParameter("token", "abc123"),
				parameters.NewPathParameter("version", "1"),
			},
			expectedBody: map[string]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parameters.CreateBody(tt.parameters...)
			for key, expectedValue := range tt.expectedBody {
				if result[key] != expectedValue {
					t.Errorf("Expected %v for key %s, but got %v", expectedValue, key, result[key])
				}
			}
		})
	}
}
