import React from 'react'

const TeamInfo = ({ id, name, displayName, owner }) => (
    <div className="team-info">
        <div className="team-info__display-name">{displayName}</div>
        <span className="team-info__name">({name})</span>
        <div className="team-info__id">{id}</div>
        <div className="team-info__owner">
            <span className="team-info__owner-name">
                owner: {owner.username}
            </span>
        </div>
    </div>
)

export default TeamInfo
