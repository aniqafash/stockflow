import useSWR from 'swr'
import { apiRequest } from '../hooks/apiRequest'
import type { Product } from '../types/product'
import { ProductCard } from '../components/ProductCard'

export function ProductsList() {
  const { data, isLoading, error } = useSWR('/products', () =>
    apiRequest<Product[]>('/products')
  )
  const products = data?.data
  return (
    <section>
      <h2 className="my-4 p-4 text-xl text-amber-700">Products List</h2>
      {isLoading && <div>Loading...</div>}
      {error && <div>Error loading products</div>}
      {products && !error && !isLoading && (
        <div className="m-auto grid grid-cols-3 gap-2 p-4">
          {products.map((p) => (
            <ProductCard
              id={p.id}
              sku={p.sku}
              name={p.name}
              price={p.price}
              stock={p.stock}
              isActive={p.isActive}
              createdAt={p.createdAt}
              lastMovementAt={p.lastMovementAt}
            />
          ))}
        </div>
      )}
    </section>
  )
}
