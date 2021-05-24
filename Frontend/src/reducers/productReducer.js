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

const reducer = (state = { list: [], shown: null }, action) => {
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
        default:
            return state
    }
}

export default reducer