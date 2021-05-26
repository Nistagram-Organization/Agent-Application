import React from 'react'
import { Spinner as ReactSpinner } from 'react-bootstrap'

const Spinner = () => (
    <ReactSpinner animation="border" role="status">
        <span className="sr-only">Loading...</span>
    </ReactSpinner>
)

export default Spinner