package main

type contextKey string

func (key contextKey) String() string {
	return "kr.cometkim.mattermost-plugin-graphql.contextKey." + string(key)
}

const (
	ContextSessionId = contextKey("sessionId")
)
