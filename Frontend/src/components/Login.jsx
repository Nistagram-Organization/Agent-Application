import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import { useHistory } from 'react-router-dom'
import { restore_login } from '../reducers/authenticationReducer'
import * as yup from 'yup'
import { Formik } from 'formik'
import { Button, Col, Form } from 'react-bootstrap'

const loginSchema = yup.object().shape({
    username: yup.string()
        .min(4, 'Username must be at least 4 characters')
        .required('Username is required'),
    password: yup.string()
        .required('Password is required')
})

const Login = () => {
    const dispatch = useDispatch()
    const history = useHistory()

    useEffect(() => {
        if (window.localStorage.getItem('token') !== null
            && window.localStorage.getItem('subject') !== null) {
            dispatch(restore_login()).then(() => history.push('/dashboard'))
        }
    }, [dispatch, history])

    return (
        <Formik
            validationSchema={loginSchema}
            onSubmit={console.log}
            initialValues={{
                username: '',
                password: ''
            }}
        >
            {({
                  handleSubmit,
                  handleChange,
                  values,
                  errors,
              }) => (
                <Form noValidate onSubmit={handleSubmit}>
                    <Form.Row>
                        <Form.Group as={Col} md="6" controlId="username">
                            <Form.Label>Username</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="Username"
                                name="username"
                                value={values.username}
                                onChange={handleChange}
                                isInvalid={!!errors.username}
                            />
                            <Form.Control.Feedback type="invalid">
                                {errors.username}
                            </Form.Control.Feedback>
                        </Form.Group>
                    </Form.Row>
                    <Form.Row>
                        <Form.Group as={Col} md="6" controlId="password">
                            <Form.Label>Password</Form.Label>
                            <Form.Control
                                type="password"
                                placeholder="Password"
                                name="password"
                                value={values.password}
                                onChange={handleChange}
                                isInvalid={!!errors.password}
                            />
                            <Form.Control.Feedback type="invalid">
                                {errors.password}
                            </Form.Control.Feedback>
                        </Form.Group>
                    </Form.Row>
                    <Button type="submit">Login</Button>
                </Form>
            )}
        </Formik>
    )
}

export default Login