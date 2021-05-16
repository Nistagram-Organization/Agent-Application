import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { getProducts } from '../reducers/productReducer'
import { chunk } from 'lodash'
import ProductsRow from './ProductsRow'

const Products = () => {
    const dispatch = useDispatch()

    useEffect(() => {
        dispatch(getProducts())
    }, [])

    const products = useSelector(state => state.products.list)

    return (
        <div style={{ marginTop: '2%' }}>
            {
                chunk(products, 4)
                    .map((productsChunk, i) =>
                        <ProductsRow key={i} products={productsChunk}/>
                    )
            }
        </div>
    )
}

export default Products