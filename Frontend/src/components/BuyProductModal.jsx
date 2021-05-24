import React from 'react'
import { Modal } from 'react-bootstrap'
import DeliveryInformationForm from './DeliveryInformationForm'

const BuyProductModal = ({ visible, toggle }) => {
    return (
        <Modal show={visible} onHide={toggle} size="xl">
            <Modal.Header closeButton>
                <Modal.Title>Buy a product</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <DeliveryInformationForm toggleModal={toggle}/>
            </Modal.Body>
        </Modal>
    )
}

export default BuyProductModal