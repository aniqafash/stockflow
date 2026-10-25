interface ProductProps {
  id: number
  sku: string
  name: string
  price: number
  stock: number
  isActive: boolean
  createdAt: string
  lastMovementAt: string | null
}

export function ProductCard({
  id,
  sku,
  name,
  price,
  stock,
  isActive,
  createdAt,
  lastMovementAt,
}: ProductProps) {
  const createdAtDate = new Date(createdAt)
  const lastMovementAtDate = lastMovementAt ? new Date(lastMovementAt) : null
  return (
    <article
      data-id={id}
      className="rounded-2xl border-2 border-amber-500 p-4 shadow-sm shadow-amber-500 hover:scale-102 hover:bg-gray-100"
    >
      <div className="flex items-center justify-between">
        <h3 className="text-bold flex flex-col gap-2 text-lg text-amber-500">
          {name}
        </h3>
        <span
          className={
            (isActive
              ? 'bg-green-200 text-green-800'
              : 'bg-red-200 text-red-800') + ' rounded-md p-2 text-sm'
          }
        >
          {isActive ? 'Active' : 'Inactive'}
        </span>
      </div>
      <div className="my-2 flex justify-between rounded-xl bg-amber-500 p-2 text-white">
        <span>{sku}</span>
        <span>Rp{price.toLocaleString('id-ID')}</span>
      </div>
      <div className="flex flex-col gap-2">
        <span>{stock} left</span>
        <span>{createdAtDate.toDateString()}</span>
        <span>
          {lastMovementAtDate
            ? lastMovementAtDate.toDateString()
            : 'No movement yet'}
        </span>
      </div>
    </article>
  )
}
