import React from 'react'
import { Card } from 'react-bootstrap'
import NavbarLink from './NavbarLink'
import { Nav } from 'react-bootstrap'
import CurrencyFormat from 'react-currency-format'
import { LinkContainer } from 'react-router-bootstrap'

const ProductPreview = ({ id, name, price, image }) => {
    return (
        <Card>
            <LinkContainer to={`/dashboard/products/${id}`}>
                <Nav.Link>
                    <Card.Img variant="top" src={image} style={{ width: '200px', height: '200px' }}/>
                </Nav.Link>
            </LinkContainer>
            <Card.Body>
                <Card.Title>
                    <NavbarLink text={name} url={`/dashboard/products/${id}`}/>
                </Card.Title>
                <Card.Text>
                    <CurrencyFormat value={price} displayType={'text'} thousandSeparator={true} suffix={' RSD'}
                                    renderText={value => <b>{value}</b>}/>
                </Card.Text>
            </Card.Body>
        </Card>
    )
}

export default ProductPreview