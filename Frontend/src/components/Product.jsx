import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useHistory, useRouteMatch } from 'react-router-dom'
import { getProduct } from '../reducers/productReducer'
import { Button, Col, Form, Row } from 'react-bootstrap'
import CurrencyFormat from 'react-currency-format'
import { Formik } from 'formik'
import * as yup from 'yup'
import { setBuyOrder } from '../reducers/invoiceReducer'
import { setNotification } from '../reducers/notificationReducer'
import productService from '../services/productService'
import Spinner from './Spinner'
import ProductModal from './ProductModal'
import { toggleModal } from '../reducers/modalReducer'
import { useAuth0 } from '@auth0/auth0-react'

const buyModalSchema = yup.object().shape({
    quantity: yup
        .number()
        .positive('Quantity must be a positive number')
        .required('Quantity must be specified')
})

const Product = () => {
    const dispatch = useDispatch()
    const history = useHistory()
    const product = useSelector(state => state.products.shown)
    const { getAccessTokenSilently } = useAuth0()
    const [ token, setToken ] = useState(null)

    useEffect(() => {
        getAccessTokenSilently().then(t => setToken(t))
    }, [])

    const openBuyModal = async (values) => {
        dispatch(setBuyOrder(product.id, values.quantity))
        dispatch(toggleModal('BUY'))
    }

    const openEditModal = () => {
        dispatch(toggleModal('EDIT'))
    }

    const deleteProduct = async () => {
        try {
            await productService.deleteProduct(product.id, token)
            dispatch(setNotification('Product deleted successfully', 'success', 3000))
            history.push('/dashboard/products')
        } catch (e) {
            dispatch(setNotification(e.response.data.message, 'error', 3000))
        }
    }

    const idMatch = useRouteMatch('/dashboard/products/:id')

    useEffect(() => {
        if (!product || product.id !== idMatch) {
            idMatch && dispatch(getProduct(idMatch.params.id))
        }
    }, [])

    if (!product) {
        return (
            <Spinner/>
        )
    }

    return (
        <div>
            <ProductModal/>
            <Row>
                <Col>
                    <h1>{product.name}</h1>
                </Col>
            </Row>
            <Row>
                <Col>
                    <img src={product.image} alt={product.name} style={{ width: '300px', height: '300px' }}/>
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
                                            <Button type="submit">Buy</Button>
                                        </Col>
                                    </Row>
                                </Form>)}
                        </Formik>
                    </Row>
                    <Row style={{ marginTop: '1%' }}>
                        <Col sm={7}>
                            <Button onClick={openEditModal}>Edit product</Button>
                        </Col>
                    </Row>
                    <Row style={{ marginTop: '1%' }}>
                        <Col sm={7}>
                            <Button variant="danger" onClick={deleteProduct}>Delete product</Button>
                        </Col>
                    </Row>
                </Col>
            </Row>
        </div>
    )
}

export default Product