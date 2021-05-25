import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { getProducts } from '../reducers/productReducer'
import { Row, Col } from 'react-flexa'
import ProductPreview from './ProductPreview'
import CreateProduct from './CreateProduct'
import CreateProductModal from './CreateProductModal'

const Products = () => {
    const dispatch = useDispatch()

    useEffect(() => {
        dispatch(getProducts())
    }, [])

    const products = useSelector(state => state.products.list)
    const [modalVisible, setModalVisible] = useState(false)
    const toggleModal = () => setModalVisible(!modalVisible)

    return (
        <>
            <CreateProductModal toggle={toggleModal} visible={modalVisible}/>
            <Row justifyContent='center'>
                {
                    products.map(p =>
                        <Col key={p.id} style={{ marginBottom: '1%' }}>
                            <ProductPreview {...p}/>
                        </Col>
                    )
                }
                <Col>
                    <CreateProduct toggleModal={toggleModal}/>
                </Col>
            </Row>
        </>
    )
}

export default Products