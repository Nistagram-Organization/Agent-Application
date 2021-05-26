import React, { useEffect, useState } from 'react'
import Chart from 'react-apexcharts'
import reportsService from '../services/reportsService'
import Spinner from './Spinner'

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

    const [data, setData] = useState(null)

    useEffect(() => {
        reportsService.generateReport().then(result => {
            result.forEach(product => {
                initialSettings.options.xaxis.categories.push(product.product_name)
                initialSettings.series[0].data.push(product.sold)
                initialSettings.series[1].data.push(product.income)
            })
            setData(initialSettings)
            console.log(data)
        })
    }, [])

    if(!data) {
        return (
            <Spinner/>
        )
    }

    return (
        <div>
            <Chart type='bar' options={data.options} series={data.series} height={430}/>
        </div>
    )
}

export default Reports