import React, { useEffect, useState } from 'react'
import Chart from 'react-apexcharts'
import reportsService from '../services/reportsService'
import Spinner from './Spinner'
import { useAuth0 } from '@auth0/auth0-react'

const Reports = () => {
    const initialSettings = {
        series: [
            {
                name: 'Sold',
                data: []
            },
            {
                name: 'Income in RSD',
                data: []
            }
        ],
        options: {
            chart: {
                type: 'bar',
                height: 430
            },
            xaxis: {
                categories: []
            }
        },
    }

    const { getAccessTokenSilently } = useAuth0()
    const [data, setData] = useState(null)
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        getAccessTokenSilently().then(token => {
            reportsService.generateReport(token).then(result => {
                if (result) {
                    result.forEach(product => {
                        initialSettings.options.xaxis.categories.push(product.product_name)
                        initialSettings.series[0].data.push(product.sold)
                        initialSettings.series[1].data.push(product.income)
                    })
                    setData(initialSettings)
                }
                setLoading(false)
            })
        })
    }, [])

    if(loading) {
        return (
            <Spinner/>
        )
    }

    if(!data) {
        return (
            <p>No products have been sold!</p>
        )
    }

    return (
        <div>
            <Chart type='bar' options={data.options} series={data.series} height={430}/>
        </div>
    )
}

export default Reports