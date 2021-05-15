import React from 'react'
import Navbar from './Navbar'
import { GuardedRoute, GuardProvider } from 'react-router-guards'
import { Redirect } from 'react-router-dom'
import Products from './Products'
import Product from './Product'


const Dashboard = () => {

    const authenticated = false
    const requireAuthentication = (to, from, next) => {
        if(to.meta.authenticated) {
            if(authenticated) {
                next()
            }
            next.redirect(from)
        }
        next()
    }

    return (
        <div>
            <Navbar authenticated={authenticated}/>
            <GuardProvider guards={[requireAuthentication]}>
                <GuardedRoute path='/dashboard/products/:id'>
                    <Product/>
                </GuardedRoute>
                <GuardedRoute exact path='/dashboard/products'>
                    <Products/>
                </GuardedRoute>
                <GuardedRoute exact path='/dashboard/reports' meta={{ authenticated: true }}>
                    <p>reports</p>
                </GuardedRoute>
                <GuardedRoute exact path='/dashboard'>
                    <Redirect to='/dashboard/products'/>
                </GuardedRoute>
            </GuardProvider>
        </div>
    )
}

export default Dashboard