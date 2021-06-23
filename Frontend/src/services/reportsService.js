import axios from 'axios'

const BASE_URL = process.env.REACT_APP_BASE_URL

const generateReport = async (token) => {
    const response = await axios.get(`${BASE_URL}/reports`, { headers: { Authorization: `Bearer ${token}` } })
    return response.data
}

const reportsService = {
    generateReport
}

export default reportsService