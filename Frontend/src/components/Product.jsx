import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useRouteMatch } from 'react-router-dom'
import { getProduct, setBuyOrder } from '../reducers/productReducer'
import { Button, Col, Form, Row, Spinner } from 'react-bootstrap'
import CurrencyFormat from 'react-currency-format'
import BuyProductModal from './BuyProductModal.jsx'
import { Formik } from 'formik'
import * as yup from 'yup'

const buyModalSchema = yup.object().shape({
    quantity: yup
        .number()
        .positive('Quantity must be a positive number')
        .required('Quantity must be specified')
})

const Product = () => {
    const dispatch = useDispatch()
    const product = useSelector(state => state.products.shown)

    const [modalVisible, setModalVisible] = useState(false)

    const toggleModal = () => setModalVisible(!modalVisible)

    const openBuyModal = async (values) => {
        dispatch(setBuyOrder(product.id, values.quantity))
        toggleModal()
    }

    const idMatch = useRouteMatch('/dashboard/products/:id')

    useEffect(() => {
        if (!product || product.id !== idMatch) {
            idMatch && dispatch(getProduct(idMatch.params.id))
        }
    }, [])

    if (!product) {
        return (
            <Spinner animation="border" role="status">
                <span className="sr-only">Loading...</span>
            </Spinner>
        )
    }

    return (
        <div>
            <BuyProductModal visible={modalVisible} toggle={toggleModal}/>
            <Row>
                <Col>
                    <h1>{product.name}</h1>
                </Col>
            </Row>
            <Row>
                <Col>
                    <img src={product.image} alt={product.name}/>
                </Col>
                <Col>
                    <Row>
                        <p>{product.description}</p>
                    </Row>
                    <Row>
                        <CurrencyFormat value={product.price} displayType={'text'} thousandSeparator={true} suffix={' RSD'}
                                        renderText={value => <b>{value}</b>}/>
                    </Row>
                    <Row>
                        <Formik
                            validationSchema={buyModalSchema}
                            onSubmit={openBuyModal}
                            initialValues={{
                                quantity: 0
                            }}
                        >
                            {(formik) => (
                                <Form noValidate onSubmit={formik.handleSubmit}>
                                    <Form.Group as={Row} controlId="quantity">
                                        <Form.Label column sm={4}>Quantity</Form.Label>
                                        <Col sm={7}>
                                            <Form.Control
                                                type="number"
                                                placeholder="Quantity"
                                                name="quantity"
                                                value={formik.values.quantity}
                                                onChange={formik.handleChange}
                                                isInvalid={!!formik.errors.quantity}
                                            />
                                            <Form.Control.Feedback type="invalid">
                                                {formik.errors.quantity}
                                            </Form.Control.Feedback>
                                        </Col>
                                    </Form.Group>
                                    <Row style={{ marginTop: '1%' }}>
                                        <Col sm={4}/>
                                        <Col sm={7}>
                                            <Button type="button" onClick={() => formik.submitForm()}>Buy</Button>
                                        </Col>
                                    </Row>
                                </Form>)}
                        </Formik>
                    </Row>
                </Col>
            </Row>
        </div>
    )
}

export default Product