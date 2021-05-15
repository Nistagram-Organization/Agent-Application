import React from 'react'
import store from './store'
import axios from 'axios'
import Toaster from './components/Toaster'
import { Switch, Route, Redirect } from 'react-router-dom'
import Dashboard from './components/Dashboard'
import Login from './components/Login'
import Logout from './components/Logout'

axios.interceptors.request.use(
    request => {
        const state = store.getState()
        const token = state.authentication.token

        if (token) {
            request.headers['Authorization'] = `Bearer ${token}`
        }
        return request
    },
    error => Promise.reject(error)
)

const App = () => (
    <div className='container'>
        <Toaster/>
        <Switch>
            <Route path='/login'>
                <Login/>
            </Route>
            <Route path='/logout'>
                <Logout/>
            </Route>
            <Route path='/dashboard'>
                <Dashboard/>
            </Route>
            <Route exact path='*'>
                <Redirect to='/dashboard'/>
            </Route>
        </Switch>
    </div>
)

export default App
