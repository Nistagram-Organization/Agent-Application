import React from 'react'
import * as yup from 'yup'
import { Formik } from 'formik'
import { Button, Col, Form, Row } from 'react-bootstrap'
import { Row as FlexRow, Col as FlexCol } from 'react-flexa'
import { useDispatch, useSelector } from 'react-redux'
import { setNotification } from '../reducers/notificationReducer'
import invoicesService from '../services/invoicesService'
import { setBuyOrder } from '../reducers/invoiceReducer'
import { toggleModal } from '../reducers/modalReducer'

const deliveryInformationSchema = yup.object().shape({
    name: yup
        .string()
        .required('Name must be specified'),
    surname: yup
        .string()
        .required('Surname must be specified'),
    phone: yup
        .string()
        .required('Phone must be specified'),
    address: yup
        .string()
        .required('Address must be specified'),
    city: yup
        .string()
        .required('City must be specified'),
    zipCode: yup
        .number()
        .positive('Zip code must be a positive number')
        .required('Zip code must be specified')
})

const DeliveryInformationForm = () => {
    const dispatch = useDispatch()
    const productID = useSelector(state => state.invoices.productId)
    const quantity = useSelector(state => state.invoices.quantity)

    const buyProduct = async (values) => {
        let invoice = {
            invoiceItems: [{
                productID,
                quantity
            }],
            deliveryInformation: values
        }

        try {
            await invoicesService.sendInvoice(invoice)
            dispatch(setNotification('Product bought successfully', 'success', 3000))
        } catch (e) {
            dispatch(setNotification(e.response.data.message, 'error', 3000))
        }
        dispatch(setBuyOrder(0, 0))
        dispatch(toggleModal(''))
    }

    return (
        <FlexRow justifyContent='center'>
            <FlexCol xs={6}>
                <Formik
                    validationSchema={deliveryInformationSchema}
                    onSubmit={buyProduct}
                    initialValues={{
                        name: '',
                        surname: '',
                        phone: '',
                        address: '',
                        city: '',
                        zipCode: 0
                    }}
                >
                    {(formik) => (
                        <Form noValidate onSubmit={formik.handleSubmit}>
                            <Form.Group as={Row} controlId="name">
                                <Form.Label column sm={4}>Name</Form.Label>
                                <Col sm={7}>
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
                            <Form.Group as={Row} controlId="surname">
                                <Form.Label column sm={4}>Surname</Form.Label>
                                <Col sm={7}>
                                    <Form.Control
                                        type="text"
                                        placeholder="Surname"
                                        name="surname"
                                        value={formik.values.surname}
                                        onChange={formik.handleChange}
                                        isInvalid={!!formik.errors.surname}
                                    />
                                    <Form.Control.Feedback type="invalid">
                                        {formik.errors.surname}
                                    </Form.Control.Feedback>
                                </Col>
                            </Form.Group>
                            <Form.Group as={Row} controlId="phone">
                                <Form.Label column sm={4}>Phone</Form.Label>
                                <Col sm={7}>
                                    <Form.Control
                                        type="text"
                                        placeholder="Phone"
                                        name="phone"
                                        value={formik.values.phone}
                                        onChange={formik.handleChange}
                                        isInvalid={!!formik.errors.phone}
                                    />
                                    <Form.Control.Feedback type="invalid">
                                        {formik.errors.phone}
                                    </Form.Control.Feedback>
                                </Col>
                            </Form.Group>
                            <Form.Group as={Row} controlId="address">
                                <Form.Label column sm={4}>Address</Form.Label>
                                <Col sm={7}>
                                    <Form.Control
                                        type="text"
                                        placeholder="Address"
                                        name="address"
                                        value={formik.values.address}
                                        onChange={formik.handleChange}
                                        isInvalid={!!formik.errors.address}
                                    />
                                    <Form.Control.Feedback type="invalid">
                                        {formik.errors.address}
                                    </Form.Control.Feedback>
                                </Col>
                            </Form.Group>
                            <Form.Group as={Row} controlId="city">
                                <Form.Label column sm={4}>City</Form.Label>
                                <Col sm={7}>
                                    <Form.Control
                                        type="text"
                                        placeholder="City"
                                        name="city"
                                        value={formik.values.city}
                                        onChange={formik.handleChange}
                                        isInvalid={!!formik.errors.city}
                                    />
                                    <Form.Control.Feedback type="invalid">
                                        {formik.errors.city}
                                    </Form.Control.Feedback>
                                </Col>
                            </Form.Group>
                            <Form.Group as={Row} controlId="zipCode">
                                <Form.Label column sm={4}>Zip code</Form.Label>
                                <Col sm={7}>
                                    <Form.Control
                                        type="number"
                                        placeholder="Zip Code"
                                        name="zipCode"
                                        value={formik.values.zipCode}
                                        onChange={formik.handleChange}
                                        isInvalid={!!formik.errors.zipCode}
                                        min="1"
                                    />
                                    <Form.Control.Feedback type="invalid">
                                        {formik.errors.zipCode}
                                    </Form.Control.Feedback>
                                </Col>
                            </Form.Group>
                            <Row style={{ marginTop: '1%' }}>
                                <Col sm={4}/>
                                <Col sm={7}>
                                    <Button type="button" onClick={() => formik.submitForm()}>Buy product</Button>
                                </Col>
                            </Row>
                        </Form>
                    )}
                </Formik>
            </FlexCol>
        </FlexRow>
    )
}

export default DeliveryInformationForm