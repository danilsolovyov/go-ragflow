package goragflow_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	goragflow "github.com/danilsolovyov/go-ragflow"
	"github.com/danilsolovyov/go-ragflow/options"
)

func rawSessionResponse(data []byte) []byte {
	return []byte(`{"data":` + string(data) + `}`)
}

func TestSession_Completions_Integration(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(rawSessionResponse([]byte(`{"answer":"Hello, world!"}`)))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	clientOpts := &options.ClientOptions{
		Scheme:    "http",
		Host:      ts.Listener.Addr().String(),
		APIKey:    "test",
		Transport: http.DefaultTransport,
	}
	client := goragflow.NewClient(clientOpts)

	sess := &goragflow.Session{ID: "sess1", AgentID: "agent1"}
	sess.SetClient(client)
	ctx := t.Context()

	comp, err := sess.Completions(ctx, &options.CompletionsOptions{Question: "Hi"})
	require.NoError(t, err)
	assert.Equal(t, "Hello, world!", comp.Answer)
}

func TestSession_DeleteSession_Integration(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	clientOpts := &options.ClientOptions{
		Scheme:    "http",
		Host:      ts.Listener.Addr().String(),
		APIKey:    "test",
		Transport: http.DefaultTransport,
	}
	client := goragflow.NewClient(clientOpts)

	sess := &goragflow.Session{ID: "sess1", AgentID: "agent1"}
	sess.SetClient(client)
	ctx := t.Context()

	err := sess.DeleteSession(ctx)
	require.NoError(t, err)
}

func TestSession_DeleteSession_Errors(t *testing.T) {
	ctx := t.Context()
	client := &goragflow.Client{}

	sessNoID := &goragflow.Session{AgentID: "agent1"}
	sessNoID.SetClient(client)
	err := sessNoID.DeleteSession(ctx)
	require.Error(t, err)

	sessNoAgent := &goragflow.Session{ID: "sess1"}
	sessNoAgent.SetClient(client)
	err = sessNoAgent.DeleteSession(ctx)
	require.NoError(t, err)
}
