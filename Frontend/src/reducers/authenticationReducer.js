import loginService from '../services/authenticationService'
import jwt_decode from 'jwt-decode'
import { setNotification } from './notificationReducer'

const local_login = (token, subject) => {
    window.localStorage.setItem('token', JSON.stringify(token))
    window.localStorage.setItem('subject', JSON.stringify(subject))
}

const local_logout = () => {
    window.localStorage.removeItem('token')
    window.localStorage.removeItem('subject')
}

export const restore_login = () => {
    return async dispatch => {
        const token = JSON.parse(window.localStorage.getItem('token'))
        const subject = JSON.parse(window.localStorage.getItem('subject'))

        dispatch({
            type: 'LOGIN',
            token,
            subject
        })
    }
}

export const login = (username, password) => {
    return async dispatch => {
        let token = ''
        try {
            token = await loginService.login(username, password)
        } catch (e) {
            dispatch(setNotification(e.response.data, 'error', 3500))
            return
        }

        const subject = jwt_decode(token).sub

        local_login(token, subject)

        dispatch({
            type: 'LOGIN',
            token,
            subject
        })
    }
}

export const logout = () => {
    return async dispatch => {
        local_logout()

        dispatch({
            type: 'LOGOUT'
        })

        dispatch({
            type: 'DESTROY_SESSION'
        })
    }
}

const reducer = (state = { token: null, subject: null }, action) => {
    switch (action.type) {
        case 'LOGIN':
            return {
                token: action.token,
                subject: action.subject
            }
        case 'LOGOUT':
            return {
                token: null,
                subject: null
            }
        default:
            return state
    }
}

export default reducer