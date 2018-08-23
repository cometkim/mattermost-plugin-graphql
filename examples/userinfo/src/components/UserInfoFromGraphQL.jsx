import React from 'react'
import ApolloClient from 'apollo-boost'
import gql from 'graphql-tag'

import config from '../config'

import Example from './Exmaple'
import UserInfo from './UserInfo'

export default class UserInfoFromGraphQL extends React.PureComponent {
    state = {
        user: null,
        start: 0,
        end: 0,
    }

    client = new ApolloClient({
        uri: `${config.baseUrl}/plugins/kr.cometkim.mattermost-plugin-graphql`,
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${config.token}`,
        },
    })

    traceStart = () => {
        this.setState({
            start: Date.now(),
            end: 0,
        })
    }

    traceEnd = () => {
        this.setState({
            end: Date.now(),
        })
    }

    fetchUserInfo = async () => {
        this.traceStart()

        const { data } = await this.client.query({
            query: gql`
                {
                    me {
                        id
                        username
                        email
                        joinedTeams {
                            id
                            name
                            displayName
                            owner {
                                username
                            }
                        }
                    }
                }
            `,
        })

        this.traceEnd()

        this.setState({
            user: data.me,
        })
    }

    render() {
        const elapsedMillis = this.state.end - this.state.start

        return (
            <div className="graphql-example">
                <Example name="GraphQL Example" execute={this.fetchUserInfo}>
                    {elapsedMillis > 0 && (
                        <div className="trace">
                            Elapsed
                            <time className={`trace__time ${elapsedMillis > 100 ? 'danger' : ''}`}>
                                {elapsedMillis}
                            </time>
                            millis
                            <div className="trace__count">Only 1 query!!</div>
                        </div>
                    )}
                    {this.state.user && <UserInfo {...this.state.user} />}
                </Example>
            </div>
        )
    }
}
