import React from 'react'

const Example = ({ children, name, execute }) => (
    <div className="example">
        <h1 className="example__title">{name}</h1>
        <button className="example__button" onClick={execute}>
            Click to Fetch
        </button>
        {children}
    </div>
)

export default Example
