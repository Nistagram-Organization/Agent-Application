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

const createProduct = async (product, token) => {
    const response = await axios.post(`${BASE_URL}/products`, product, { headers: { Authorization: `Bearer ${token}` } })
    return response.data
}

const editProduct = async (product, token) => {
    const response = await axios.put(`${BASE_URL}/products`, product, { headers: { Authorization: `Bearer ${token}` } })
    return response.data
}

const deleteProduct = async (id, token) => {
    const response = await axios.delete(`${BASE_URL}/products/${id}`, { headers: { Authorization: `Bearer ${token}` } })
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