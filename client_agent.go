package goragflow

import (
	"context"
	"net/http"

	"github.com/danilsolovyov/go-ragflow/options"
)

// GetAgents retrieves a list of agents with the specified options.
// It merges the provided options with default values and makes a GET request to the agents endpoint.
//
// @param ctx context.Context - The context for the request.
// @param opts *options.GetAgentsOptions - Options for filtering, pagination, and sorting agents.
// @return []*Agent - A slice of Agent pointers representing the retrieved agents.
// @return error - An error if the request fails.
func (c *Client) GetAgents(ctx context.Context, opts *options.GetAgentsOptions) ([]*Agent, error) {
	path := "/agents"
	opts = options.DefaultGetAgentsOptions().Merge(opts)
	params := opts.Parameters()

	result := []*Agent{}
	err := c.do(ctx, http.MethodGet, path, &result, params...)
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		item.client = c
	}

	return result, nil
}

// GetAgent retrieves an agent with the specified ID.
// It makes a GET request to the agents endpoint with the ID as a path parameter.
//
// @param ctx context.Context - The context for the request.
// @param id string - The ID of the agent to retrieve.
// @return *Agent - A pointer to the Agent.
// @return error - An error if the request fails.
func (c *Client) GetAgent(ctx context.Context, id string) (*Agent, error) {
	agents, err := c.GetAgents(ctx, &options.GetAgentsOptions{ID: id})
	if err != nil {
		return nil, err
	}

	if len(agents) == 0 {
		return nil, ErrAgentNotFound
	}

	agent := agents[0]
	agent.client = c

	return agent, nil
}
