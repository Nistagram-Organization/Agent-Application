import React from 'react'
import { Modal } from 'react-bootstrap'
import CreateProductForm from './CreateProductForm'

const CreateProductModal = ({ visible, toggle }) =>
    (
        <Modal show={visible} onHide={toggle} size="md">
            <Modal.Header closeButton>
                <Modal.Title>Create product</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <CreateProductForm toggleModal={toggle}/>
            </Modal.Body>
        </Modal>
    )

export default CreateProductModal