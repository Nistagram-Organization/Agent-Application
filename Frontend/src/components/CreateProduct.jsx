import React from 'react'
import { Button, Card } from 'react-bootstrap'
import { Row, Col } from 'react-flexa'
import { toggleModal } from '../reducers/modalReducer'
import { useDispatch } from 'react-redux'

const CreateProduct = () => {
    const dispatch = useDispatch()

    return (
        <Card style={{ width: '200px', height: '200px' }}>
            <Card.Body>
                <Col>
                    <Row justifyContent="center">
                        <Button style={{ marginTop: '15%' }} onClick={() => dispatch(toggleModal('CREATE'))}>Create product</Button>
                    </Row>
                </Col>
            </Card.Body>
        </Card>
    )
}

export default CreateProduct