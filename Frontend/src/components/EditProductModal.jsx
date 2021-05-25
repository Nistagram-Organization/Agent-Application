import React from 'react'
import { Modal } from 'react-bootstrap'
import EditProductForm from './EditProductForm'

const EditProductModal = ({ visible, toggle }) =>
    (
        <Modal show={visible} onHide={toggle} size="md">
            <Modal.Header closeButton>
                <Modal.Title>Edit product</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <EditProductForm toggleModal={toggle}/>
            </Modal.Body>
        </Modal>
    )

export default EditProductModal