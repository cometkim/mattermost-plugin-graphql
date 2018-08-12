package main

import (
	"context"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/mattermost/mattermost-server/plugin"
)

type GraphQLPlugin struct {
	plugin.MattermostPlugin
	handler *handler.Handler
}

func (p *GraphQLPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// `c.SessionId` appears to be valid only on first call
	// p.API.LogDebug("debug", "session_id", c.SessionId)
	ctx = context.WithValue(ctx, ContextSessionId, c.SessionId)

	p.handler.ContextHandler(ctx, w, r)
}

func main() {
	instance := &GraphQLPlugin{}

	schema, err := instance.initSchema()
	if err != nil {
		instance.API.LogError("Failed to init schema")
		return
	}

	instance.handler = handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	plugin.ClientMain(instance)
}
