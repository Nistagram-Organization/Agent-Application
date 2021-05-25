import React from 'react'
import { Button, Card } from 'react-bootstrap'
import { Row, Col } from 'react-flexa'

const CreateProduct = ({ toggleModal }) => {
    return (
        <Card style={{ width: '200px', height: '200px' }}>
            <Card.Body>
                <Col>
                    <Row justifyContent="center">
                        <Button style={{ marginTop: '15%' }} onClick={toggleModal}>Create product</Button>
                    </Row>
                </Col>
            </Card.Body>
        </Card>
    )
}

export default CreateProduct