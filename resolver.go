package main

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/mattermost/mattermost-server/model"
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

	// Nil check required for the [Issue](https://github.com/graphql-go/graphql/issues/134)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Contextual resolver example
func (p *GraphQLPlugin) resolveCurrentUser(param graphql.ResolveParams) (interface{}, error) {
	user, err := p.API.GetUser(param.Context.Value(ContextCurrentUserId).(string))
	if err != nil {
		return nil, err
	}
	return user, nil
}

// List resolver example
func (p *GraphQLPlugin) resolveAllTeams(param graphql.ResolveParams) (interface{}, error) {
	teams, err := p.API.GetTeams()

	if err != nil {
		return nil, err
	}
	return teams, nil
}

// Relational resolver example
func (p *GraphQLPlugin) resolveTeamOwner(param graphql.ResolveParams) (interface{}, error) {
	if team, ok := param.Source.(*model.Team); ok {
		owner, err := p.API.GetUserByEmail(team.Email)

		if err != nil {
			return nil, err
		}
		return owner, nil
	}
	return nil, nil
}

func (p *GraphQLPlugin) resolveTeamsForUser(param graphql.ResolveParams) (interface{}, error) {
	if user, ok := param.Source.(*model.User); ok {
		// Not supported by plugin API yet
		// teams, err := p.API.GetTeamsForUser(user.Id)
		teams, err := p.API.GetTeams()
		joinedTeams := make([]*model.Team, 0)
		for _, team := range teams {
			membership, _ := p.API.GetTeamMember(team.Id, user.Id)
			if membership != nil {
				joinedTeams = append(joinedTeams, team)
			}
		}

		if err != nil {
			return nil, err
		}
		return joinedTeams, nil
	}
	return nil, nil
}
