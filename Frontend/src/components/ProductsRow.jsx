import React from 'react'
import ProductListItem from './ProductPreview'
import { Col, Row } from 'react-bootstrap'

const ProductsRow = ({ products }) => (
    <Row xs={3} md={4} lg={4} style={{ marginBottom: '1%' }}>
        {
            products.map((product, i) =>
                <Col key={i} xs={3}>
                <ProductListItem
                    key={product.id}
                    id={product.id}
                    name={product.name}
                    image={product.image}
                    price={product.price}
                />
                </Col>
            )
        }
    </Row>
)

export default ProductsRow