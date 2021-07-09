import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import { BrowserRouter as Router } from 'react-router-dom'
import { Provider } from 'react-redux'
import store from './store'
import { Auth0Provider } from '@auth0/auth0-react'

ReactDOM.render(
    <Router>
        <Provider store={store}>
            <Auth0Provider
                domain="dev-6w-2hyw1.eu.auth0.com"
                clientId="scV8c0nay2dGleOBxq5CtaCP9idlz7U0"
                audience="http://nistagram-agent"
                redirectUri="http://localhost:3000/dashboard/products"
                useRefreshTokens={true}
                cacheLocation={'localstorage'}
            >
            <App/>
            </Auth0Provider>
        </Provider>
    </Router>,
    document.getElementById('root')
)
