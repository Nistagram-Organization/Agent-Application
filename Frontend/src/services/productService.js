import axios from 'axios'

const getProducts = async () => {
    const response = await axios.get('/api/products')
    return response.data
}

const getProduct = async (id) => {
    const response = await axios.get(`/api/products/${id}`)
    return response.data
}

const createProduct = async (product, token) => {
    const response = await axios.post('/api/products', product, { headers: { Authorization: `Bearer ${token}` } })
    return response.data
}

const editProduct = async (product, token) => {
    const response = await axios.put('/api/products', product, { headers: { Authorization: `Bearer ${token}` } })
    return response.data
}

const deleteProduct = async (id, token) => {
    const response = await axios.delete(`/api/products/${id}`, { headers: { Authorization: `Bearer ${token}` } })
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