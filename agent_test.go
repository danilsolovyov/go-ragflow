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

func rawResponseBody(data []byte) []byte {
	return []byte(`{"data":` + string(data) + `}`)
}

func TestAgent_GetMe_Integration(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(rawResponseBody([]byte(`[{"id":"test-agent","title":"Test Agent"}]`)))
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
	agent := goragflow.NewAgent("test-agent", client)
	ctx := t.Context()

	me, err := agent.GetMe(ctx)
	require.NoError(t, err)
	assert.Equal(t, "test-agent", me.ID)
	assert.Equal(t, "Test Agent", me.Title)
}

func TestAgent_ListSessions_Integration(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(rawResponseBody([]byte(`[{"id":"sess1","agent_id":"test-agent"}]`)))
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
	agent := goragflow.NewAgent("test-agent", client)
	ctx := t.Context()

	sessions, err := agent.ListSessions(ctx, nil)
	require.NoError(t, err)
	assert.Len(t, sessions, 1)
	assert.Equal(t, "sess1", sessions[0].ID)
}

func TestAgent_CreateSession_Integration(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(rawResponseBody([]byte(`{"id":"sess2","agent_id":"test-agent","user_id":"user1"}`)))
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
	agent := goragflow.NewAgent("test-agent", client)
	ctx := t.Context()

	sess, err := agent.CreateSession(ctx, &options.CreateAgentSessionOptions{UserID: "user1"})
	require.NoError(t, err)
	assert.Equal(t, "sess2", sess.ID)
	assert.Equal(t, "test-agent", sess.AgentID)
	assert.Equal(t, "user1", sess.UserID)
}

func TestAgent_DeleteSessions_Integration(t *testing.T) {
	var called bool
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		called = true
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(rawResponseBody([]byte(`{}`)))
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
	agent := goragflow.NewAgent("test-agent", client)
	ctx := t.Context()

	err := agent.DeleteSessions(ctx, []string{"sess1", "sess2"})
	require.NoError(t, err)
	assert.True(t, called, "expected handler to be called")
}
