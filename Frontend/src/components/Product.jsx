import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useRouteMatch } from 'react-router-dom'
import { getProduct } from '../reducers/productReducer'
import { Button, Col, Form, Row, Spinner } from 'react-bootstrap'
import CurrencyFormat from 'react-currency-format'

const Product = () => {
    const dispatch = useDispatch()
    const product = useSelector(state => state.products.shown)

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
                        <Form inline onSubmit={() => console.log(product.id)}>
                            <Form.Control as="input" type="number" min={1} max={product.on_stock} placeholder={1}/>
                            &nbsp;
                            <Button type="submit">Buy</Button>
                        </Form>
                    </Row>
                </Col>
            </Row>
        </div>
    )
}

export default Product