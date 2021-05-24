import productService from '../services/productService'

export const getProducts = () => {
    return async dispatch => {
        const products = await productService.getProducts()

        dispatch({
            type: 'GET_PRODUCTS',
            products
        })
    }
}

export const getProduct = (id) => {
    return async dispatch => {
        const product = await productService.getProduct(id)

        dispatch({
            type: 'GET_PRODUCT',
            product
        })
    }
}

export const setBuyOrder = (productId, quantity) => {
    return async dispatch => {
        dispatch({
            type: 'SET_BUY_ORDER',
            productId,
            quantity
        })
    }
}

const reducer = (state = { list: [], shown: null, productId: null, quantity: 0 }, action) => {
    switch (action.type) {
        case 'GET_PRODUCTS': {
            return {
                ...state,
                list: action.products
            }
        }
        case 'GET_PRODUCT': {
            return {
                ...state,
                shown: action.product
            }
        }
        case 'SET_BUY_ORDER': {
            return {
                ...state,
                productId: action.productId,
                quantity: action.quantity
            }
        }
        default:
            return state
    }
}

export default reducer