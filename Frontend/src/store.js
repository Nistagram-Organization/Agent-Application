import notificationReducer from './reducers/notificationReducer'
import { createStore, combineReducers, applyMiddleware } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'
import thunk from 'redux-thunk'
import productReducer from './reducers/productReducer'
import invoiceReducer from './reducers/invoiceReducer'
import modalReducer from './reducers/modalReducer'

const reducer = combineReducers({
    notification: notificationReducer,
    products: productReducer,
    invoices: invoiceReducer,
    modals: modalReducer
})

const store = createStore(
    reducer,
    composeWithDevTools(
        applyMiddleware(thunk)
    )
)

export default store