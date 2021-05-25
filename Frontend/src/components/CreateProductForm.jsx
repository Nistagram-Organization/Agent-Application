import React, { useState } from 'react'
import * as yup from 'yup'
import { Formik } from 'formik'
import { Button, Col, Form, Row } from 'react-bootstrap'
import ImageUploader from 'react-images-upload'
import { toBase64 } from '../image_utils'
import { useDispatch } from 'react-redux'
import { setNotification } from '../reducers/notificationReducer'
import productService from '../services/productService'
import { getProducts } from '../reducers/productReducer'

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

const CreateProductForm = ({ toggleModal }) => {
    const dispatch = useDispatch()

    const [image, setImage] = useState(null)

    const submitForm = async (values) => {
        if(!image) {
            dispatch(setNotification('Product image must be specified', 'error', 3000))
            return
        }
        const product = {
            ...values,
            image: await toBase64(image[0])
        }
        try {
            await productService.createProduct(product)
            dispatch(setNotification('Product created successfully', 'success', 3000))
            dispatch(getProducts())
            toggleModal()
        } catch (e) {
            dispatch(setNotification(e.message, 'error', 3000))
        }
    }

    return (
        <Formik
            validationSchema={productSchema}
            onSubmit={submitForm}
            initialValues={{
                name: '',
                description: '',
                price: 0,
                on_stock: 0
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
                                onChange={(i) => setImage(i)}
                                imgExtension={['.jpg', '.png', '.jpeg']}
                                buttonText='Choose image'
                                label='Max file size: 5mb, accepted: jpg, jpeg, png'
                                singleImage={true}
                                buttonType='button'
                            />
                        </Col>
                    </Form.Group>
                    <Button type='submit'>Create product</Button>
                </Form>
            )}
        </Formik>
    )
}

export default CreateProductForm