import React from 'react'

import UserInfoFromREST from './UserInfoFromREST'
import UserInfoFromGraphQL from './UserInfoFromGraphQL'

const App = () => (
    <div className="container">
        <div className="title">Mattermost User Info</div>
        <div className="wrapper">
            <div className="column">
                <UserInfoFromREST />
            </div>
            <div className="column">
                <UserInfoFromGraphQL />
            </div>
        </div>
    </div>
)

export default App
