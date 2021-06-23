import axios from 'axios'

const generateReport = async (token) => {
    const response = await axios.get('/api/reports', { headers: { Authorization: `Bearer ${token}` } })
    return response.data
}

const reportsService = {
    generateReport
}

export default reportsService