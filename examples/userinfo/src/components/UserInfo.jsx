import React from 'react'

import TeamInfo from './TeamInfo'

const UserInfo = ({ id, username, email, joinedTeams }) => (
    <div className="user-info">
        <div className="user-info__username">{username}</div>
        <span className="user-info__email">{`<${email}>`}</span>
        <div className="user-info__id">ID: {id}</div>
        <ul className="user-info__team-list">
            <div className="user-info__team-list-title">Teams</div>
            {joinedTeams.map(team => (
                <li className="user-info__team-list-item" key={team.id}>
                    <TeamInfo {...team} />
                </li>
            ))}
        </ul>
    </div>
)

export default UserInfo
