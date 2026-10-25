export type Product = {
  id: number
  sku: string
  name: string
  price: number
  stock: number
  isActive: boolean
  createdAt: string
  lastMovementAt: string | null
}
