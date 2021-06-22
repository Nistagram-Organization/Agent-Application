import React from 'react'
import { Button, Card } from 'react-bootstrap'
import { Row, Col } from 'react-flexa'
import { toggleModal } from '../reducers/modalReducer'
import { useDispatch } from 'react-redux'
import { useAuth0 } from '@auth0/auth0-react'

const CreateProduct = () => {
    const dispatch = useDispatch()
    const { isAuthenticated } = useAuth0()

    if(!isAuthenticated) {
        return null
    }

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