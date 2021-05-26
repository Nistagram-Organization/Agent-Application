import React from 'react'
import { Modal } from 'react-bootstrap'
import DeliveryInformationForm from './DeliveryInformationForm'
import { useDispatch, useSelector } from 'react-redux'
import { toggleModal } from '../reducers/modalReducer'
import ProductForm from './ProductForm'

const ProductModal = () => {
    const action = useSelector(state => state.modals.action)
    const visible = useSelector(state => state.modals.visible)
    const dispatch = useDispatch()

    const getModalTitle = () => {
        switch(action) {
            case 'BUY':
                return 'Buy product'
            case 'CREATE':
                return 'Create product'
            case 'EDIT':
                return 'Edit product'
        }
    }

    return (
        <Modal show={visible} onHide={() => dispatch(toggleModal(''))} size={action === 'BUY' ? 'xl' : 'md'}>
            <Modal.Header closeButton>
                <Modal.Title>{getModalTitle()}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                {
                    action === 'BUY' ? <DeliveryInformationForm/> : <ProductForm/>
                }
            </Modal.Body>
        </Modal>
    )
}

export default ProductModal