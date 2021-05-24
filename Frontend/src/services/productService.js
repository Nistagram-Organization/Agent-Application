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

const bookService = {
    getProducts,
    getProduct
}


export default bookService