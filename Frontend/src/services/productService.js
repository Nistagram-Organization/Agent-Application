import axios from 'axios'

const BASE_URL = process.env.REACT_APP_BASE_URL

const getProducts = async () => {
    const response = await axios.get(`${BASE_URL}/products`)
    return response.data
}

const getProduct = async (id) => {
    const response = await axios.get(`${BASE_URL}/products/${id}`)
    return response.data
}

const createProduct = async (product) => {
    const response = await axios.post(`${BASE_URL}/products`, product)
    return response.data
}

const editProduct = async (product) => {
    const response = await axios.put(`${BASE_URL}/products`, product)
    return response.data
}

const deleteProduct = async (id) => {
    const response = await axios.delete(`${BASE_URL}/products/${id}`)
    return response.data
}

const bookService = {
    getProducts,
    getProduct,
    createProduct,
    editProduct,
    deleteProduct
}


export default bookService