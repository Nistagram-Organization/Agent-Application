export const setBuyOrder = (productId, quantity) => {
    return async dispatch => {
        dispatch({
            type: 'SET_BUY_ORDER',
            productId,
            quantity
        })
    }
}

const reducer = (state = { productId: null, quantity: 0 }, action) => {
    switch (action.type) {
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