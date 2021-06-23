import axios from 'axios'

const sendInvoice = async (invoice) => {
    const response = await axios.post('/api/invoices', invoice)
    return response.data
}

const invoicesService = {
    sendInvoice
}


export default invoicesService