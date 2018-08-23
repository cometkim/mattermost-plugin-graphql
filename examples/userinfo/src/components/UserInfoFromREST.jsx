import React from 'react'

import config from '../config'

import Example from './Exmaple'
import UserInfo from './UserInfo'

export default class UserInfoFromREST extends React.PureComponent {
    state = {
        user: null,
        start: 0,
        end: 0,
        count: 0,
    }

    userRoute = `${config.baseUrl}/api/v4/users`

    defaultRequest = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${config.token}`,
        },
    }

    traceStart = () => {
        this.setState({
            start: Date.now(),
            end: 0,
            count: 0,
        })
    }

    traceEnd = () => {
        this.setState({
            end: Date.now(),
        })
    }

    fetchUserInfo = async () => {
        await this.traceStart()

        const [user, teams] = await Promise.all([
            this.fetchMe(),
            this.fetchMyTeams(),
        ])
        await Promise.all(
            teams.map(async team => {
                const user = await this.fetchUserByEmail(team.email)
                team['owner'] = user
            })
        )

        await this.traceEnd()

        this.setState({
            user: {
                ...user,
                joinedTeams: teams,
            },
        })
    }

    fetchMe = async () => {
        const user = await this.doFetch('me')
        return user
    }

    fetchMyTeams = async () => {
        const teams = await this.doFetch('me/teams')
        return teams.map(team => ({ ...team, displayName: team.display_name }))
    }

    fetchUserByEmail = async email => {
        const user = await this.doFetch(`email/${email}`)
        return user
    }

    doFetch = async uri => {
        await this.setState({
            count: this.state.count + 1,
        })

        const res = await fetch(`${this.userRoute}/${uri}`, this.defaultRequest)
        const data = await res.json()
        return data
    }

    render() {
        return (
            <Example name="JSON API Example" execute={this.fetchUserInfo}>
                {this.state.start && this.state.end ? (
                    <div className="trace">
                        Elapsed
                        <time className="trace__time">
                            {this.state.end - this.state.start}
                        </time>
                        millis
                        <div className="trace__count">
                            {this.state.count} queries
                        </div>
                    </div>
                ) : null}
                {this.state.user && <UserInfo {...this.state.user} />}
            </Example>
        )
    }
}
