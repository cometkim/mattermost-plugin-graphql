package main

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/mattermost/mattermost-server/model"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"username": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"firstname": &graphql.Field{
				Type: graphql.String,
			},
			"lastname": &graphql.Field{
				Type: graphql.String,
			},
			"nickname": &graphql.Field{
				Type: graphql.String,
			},
			"createAt": &graphql.Field{
				Type: graphql.Int,
			},
			"updateAt": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func (p *GraphQLPlugin) resolveUser(param graphql.ResolveParams) (interface{}, error) {
	id, _ := param.Args["id"].(string)
	username, _ := param.Args["username"].(string)
	email, _ := param.Args["email"].(string)

	var user *model.User
	var err *model.AppError

	if len(id) > 0 {
		user, err = p.API.GetUser(id)
	} else if len(username) > 0 {
		user, err = p.API.GetUserByUsername(username)
	} else if len(email) > 0 {
		user, err = p.API.GetUserByEmail(email)
	} else {
		return nil, errors.New("Arguments must contain one of {id, username, email}")
	}

	// I don't understand why `return user, err` will dereferences null
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *GraphQLPlugin) resolveCurrentUser(param graphql.ResolveParams) (interface{}, error) {
	session, _ := p.API.GetSession(param.Context.Value(ContextSessionId).(string))
	user, err := p.API.GetUser(session.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *GraphQLPlugin) initSchema() (graphql.Schema, error) {
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: userType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"username": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: p.resolveUser,
				},
				"me": &graphql.Field{
					Type:    userType,
					Resolve: p.resolveCurrentUser,
				},
			},
		},
	)

	schemaConfig := graphql.SchemaConfig{
		Query: queryType,

		// TODO: Bind GraphQL mutations to APIs
		// Mutation: mutationType,

		// TODO: Bind GraphQL subscriptions to Hooks
		// Subscription: subscriptionType,
	}

	return graphql.NewSchema(schemaConfig)
}
