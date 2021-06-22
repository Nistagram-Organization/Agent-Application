import React from 'react'
import Navbar from './Navbar'
import { GuardedRoute, GuardProvider } from 'react-router-guards'
import { Redirect } from 'react-router-dom'
import Products from './Products'
import Product from './Product'
import Reports from './Reports'
import { useAuth0 } from '@auth0/auth0-react'


const Dashboard = () => {
    const { isAuthenticated } = useAuth0()

    const requireAuthentication = (to, from, next) => {
        if(to.meta.authenticated) {
            if(isAuthenticated) {
                next()
            }
            next.redirect(from)
        }
        next()
    }

    return (
        <div>
            <Navbar/>
            <div style={{ marginTop: '2%' }}>
                <GuardProvider guards={[requireAuthentication]}>
                    <GuardedRoute path='/dashboard/products/:id'>
                        <Product/>
                    </GuardedRoute>
                    <GuardedRoute exact path='/dashboard/products'>
                        <Products/>
                    </GuardedRoute>
                    <GuardedRoute exact path='/dashboard/reports' meta={{ authenticated: true }}>
                        <Reports/>
                    </GuardedRoute>
                    <GuardedRoute exact path='/dashboard'>
                        <Redirect to='/dashboard/products'/>
                    </GuardedRoute>
                </GuardProvider>
            </div>
        </div>
    )
}

export default Dashboard