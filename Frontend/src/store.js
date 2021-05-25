import notificationReducer from './reducers/notificationReducer'
import authenticationReducer from './reducers/authenticationReducer'
import { createStore, combineReducers, applyMiddleware } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'
import thunk from 'redux-thunk'
import productReducer from './reducers/productReducer'
import invoiceReducer from './reducers/invoiceReducer'

const reducer = combineReducers({
    authentication: authenticationReducer,
    notification: notificationReducer,
    products: productReducer,
    invoices: invoiceReducer
})

const rootReducer = (state, action) => {
    if (action.type === 'DESTROY_SESSION') {
        state = undefined
    }
    return reducer(state, action)
}

const store = createStore(
    rootReducer,
    composeWithDevTools(
        applyMiddleware(thunk)
    )
)

export default store