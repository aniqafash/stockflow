import useSWR from 'swr'
import { apiRequest } from '../hooks/apiRequest'
import type { Product } from '../types/product'
import { ProductCard } from '../components/ProductCard'
import { useState } from 'react'

export function ProductsList() {
  const [query, setQuery] = useState('')
  const { data, isLoading, error } = useSWR('/products', () =>
    apiRequest<Product[]>('/products')
  )
  const products = data?.data
  return (
    <section>
      <h2 className="my-4 p-4 text-xl text-amber-700">Products List</h2>
      <input
        type="text"
        className="text-md ml-4 w-[95%] rounded-2xl bg-amber-100 p-4"
        value={query}
        onChange={(event) => setQuery(event.target.value)}
      />
      {isLoading && <div>Loading...</div>}
      {error && <div>Error loading products</div>}
      {products && !error && !isLoading && (
        <div className="m-auto grid grid-cols-3 gap-2 p-4">
          {products.map((p) => {
            if (
              query != '' &&
              !p.name.toLowerCase().includes(query.toLowerCase()) &&
              !p.sku.toLowerCase().includes(query.toLowerCase())
            )
              return
            return (
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
            )
          })}
        </div>
      )}
    </section>
  )
}
