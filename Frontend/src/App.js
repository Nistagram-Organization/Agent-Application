import React from 'react'
import Toaster from './components/Toaster'
import { Switch, Route, Redirect } from 'react-router-dom'
import Dashboard from './components/Dashboard'

const App = () => {
    return (
        <div className='container'>
        <Toaster/>
        <Switch>
            <Route path='/dashboard'>
                <Dashboard/>
            </Route>
            <Route exact path='*'>
                <Redirect to='/dashboard'/>
            </Route>
        </Switch>
    </div>
)
}

export default App
