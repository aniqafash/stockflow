import { createBrowserRouter } from 'react-router'
import { Layout } from '../pages/Layout'
import { ProductsList } from '../pages/ProductsList'

export const router = createBrowserRouter([
  {
    element: <Layout />,
    children: [{ path: '/', element: <ProductsList /> }],
  },
])
