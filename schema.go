package main

import (
	"time"

	"github.com/graphql-go/graphql"
)

var (
	userType *graphql.Object
	teamType *graphql.Object

	dateScalar = graphql.NewScalar(
		graphql.ScalarConfig{
			Name: "Date",
			Serialize: func(value interface{}) interface{} {
				switch v := value.(type) {
				case time.Time:
					return v.String()
				default:
					return nil
				}
			},
			ParseValue: func(value interface{}) interface{} {
				switch v := value.(type) {
				case int64:
					return time.Unix(0, v*int64(time.Millisecond))
				default:
					return nil
				}
			},
		},
	)
)

func (p *GraphQLPlugin) initSchema() (graphql.Schema, error) {
	userType = graphql.NewObject(
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
					Type: graphql.NewNonNull(dateScalar),
				},
				"updateAt": &graphql.Field{
					Type: graphql.NewNonNull(dateScalar),
				},
			},
		},
	)

	teamType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Team",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"displayName": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"createAt": &graphql.Field{
					Type: graphql.NewNonNull(dateScalar),
				},
				"updateAt": &graphql.Field{
					Type: graphql.NewNonNull(dateScalar),
				},
			},
		},
	)

	userType.AddFieldConfig("joinedTeams", &graphql.Field{
		Type:    graphql.NewList(teamType),
		Resolve: p.resolveTeamsForUser,
	})

	teamType.AddFieldConfig("owner", &graphql.Field{
		Type:    graphql.NewNonNull(userType),
		Resolve: p.resolveTeamOwner,
	})

	queryType := graphql.NewObject(
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
				"allTeams": &graphql.Field{
					Type:    graphql.NewList(teamType),
					Resolve: p.resolveAllTeams,
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
