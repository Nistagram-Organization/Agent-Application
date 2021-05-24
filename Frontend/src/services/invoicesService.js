import axios from 'axios'

const BASE_URL = process.env.REACT_APP_BASE_URL

const sendInvoice = async (invoice) => {
    const response = await axios.post(`${BASE_URL}/invoices`, invoice)
    return response.data
}

const invoicesService = {
    sendInvoice
}


export default invoicesService