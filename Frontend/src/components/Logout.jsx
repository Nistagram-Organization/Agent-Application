/*eslint-disable no-unused-vars*/
import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import { logout } from '../reducers/authenticationReducer'
import { useHistory } from 'react-router-dom'

const Logout = () => {
    const dispatch = useDispatch()
    const history = useHistory()

    useEffect(() => {
        dispatch(logout()).then(() => history.push('/dashboard/products'))
    }, [dispatch])

    return null
}

export default Logout