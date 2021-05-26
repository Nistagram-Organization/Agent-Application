export const toggleModal = (action) => {
    return async dispatch => {
        dispatch({
            type: 'TOGGLE_MODAL',
            action: action
        })
    }
}

const reducer = (state = { visible: false, action: '' }, action) => {
    switch (action.type) {
        case 'TOGGLE_MODAL': {
            return {
                visible: !state.visible,
                action: action.action
            }
        }
        default:
            return state
    }
}

export default reducer