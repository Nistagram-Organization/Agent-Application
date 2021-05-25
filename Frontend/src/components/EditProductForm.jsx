import React, { useState } from 'react'
import * as yup from 'yup'
import { Formik } from 'formik'
import { Button, Col, Form, Row } from 'react-bootstrap'
import ImageUploader from 'react-images-upload'
import { toBase64 } from '../image_utils'
import { useDispatch, useSelector } from 'react-redux'
import { setNotification } from '../reducers/notificationReducer'
import { useHistory } from 'react-router-dom'
import productService from '../services/productService'

const productSchema = yup.object().shape({
    name: yup
        .string()
        .required('Product name must be specified'),
    description: yup
        .string()
        .required('Product description must be specified'),
    price: yup
        .number()
        .required('Product price must be specified')
        .positive('Product price must be a positive number'),
    on_stock: yup
        .number()
        .required('Product stock must be specified')
        .min(0, 'Stock must be >= 0')
})

const EditProductForm = ({ toggleModal }) => {
    const dispatch = useDispatch()
    const history = useHistory()
    const product = useSelector(state => state.products.shown)

    const [image, setImage] = useState(product.image)

    const submitForm = async (values) => {
        if(!image) {
            dispatch(setNotification('Product image must be specified', 'error', 3000))
            return
        }

        const productToEdit = {
            ...values,
            image: image
        }
        try {
            productToEdit.id = product.id
            await productService.editProduct(productToEdit)
            dispatch(setNotification('Product edited successfully', 'success', 3000))
            toggleModal()
            history.push('/dashboard/products')
        } catch (e) {
            dispatch(setNotification(e.response.data.message, 'error', 3000))
        }
    }

    return (
        <Formik
            validationSchema={productSchema}
            onSubmit={submitForm}
            initialValues={{
                name: product.name,
                description: product.description,
                price: product.price,
                on_stock: product.on_stock
            }}
        >
            {(formik) => (
                <Form noValidate onSubmit={formik.handleSubmit}>
                    <Form.Group as={Row} controlId="name">
                        <Form.Label column sm={4}>Name</Form.Label>
                        <Col sm={8}>
                            <Form.Control
                                type="text"
                                placeholder="Name"
                                name="name"
                                value={formik.values.name}
                                onChange={formik.handleChange}
                                isInvalid={!!formik.errors.name}
                            />
                            <Form.Control.Feedback type="invalid">
                                {formik.errors.name}
                            </Form.Control.Feedback>
                        </Col>
                    </Form.Group>
                    <Form.Group as={Row} controlId="description">
                        <Form.Label column sm={4}>Description</Form.Label>
                        <Col sm={8}>
                            <Form.Control
                                type="text"
                                placeholder="Description"
                                name="description"
                                value={formik.values.description}
                                onChange={formik.handleChange}
                                isInvalid={!!formik.errors.description}
                            />
                            <Form.Control.Feedback type="invalid">
                                {formik.errors.description}
                            </Form.Control.Feedback>
                        </Col>
                    </Form.Group>
                    <Form.Group as={Row} controlId="price">
                        <Form.Label column sm={4}>Price</Form.Label>
                        <Col sm={8}>
                            <Form.Control
                                type="number"
                                placeholder="Price"
                                name="price"
                                value={formik.values.price}
                                onChange={formik.handleChange}
                                isInvalid={!!formik.errors.price}
                            />
                            <Form.Control.Feedback type="invalid">
                                {formik.errors.price}
                            </Form.Control.Feedback>
                        </Col>
                    </Form.Group>
                    <Form.Group as={Row} controlId="on_stock">
                        <Form.Label column sm={4}>On stock</Form.Label>
                        <Col sm={8}>
                            <Form.Control
                                type="number"
                                placeholder="On stock"
                                name="on_stock"
                                min={0}
                                value={formik.values.on_stock}
                                onChange={formik.handleChange}
                                isInvalid={!!formik.errors.on_stock}
                            />
                            <Form.Control.Feedback type="invalid">
                                {formik.errors.on_stock}
                            </Form.Control.Feedback>
                        </Col>
                    </Form.Group>
                    <Form.Group as={Row}>
                        <Form.Label column sm={4}>Product image</Form.Label>
                        <Col sm={8}>
                            <ImageUploader
                                onChange={async (i) => {
                                    if (i.length === 0) {
                                        setImage(undefined)
                                        return
                                    }
                                    setImage(await toBase64(i[0]))
                                }}
                                imgExtension={['.jpg', '.png', '.jpeg']}
                                buttonText='Choose image'
                                label='Max file size: 5mb, accepted: jpg, jpeg, png'
                                singleImage={true}
                                buttonType='button'
                                withPreview={true}
                                defaultImages={[image]}
                            />
                        </Col>
                    </Form.Group>
                    <Button type='submit'>Edit product</Button>
                </Form>
            )}
        </Formik>
    )
}

export default EditProductForm