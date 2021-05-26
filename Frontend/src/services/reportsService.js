import axios from 'axios'

const BASE_URL = process.env.REACT_APP_BASE_URL

const generateReport = async () => {
    const response = await axios.get(`${BASE_URL}/reports`)
    return response.data
}

const reportsService = {
    generateReport
}

export default reportsService