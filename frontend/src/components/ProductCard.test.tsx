import { describe, expect, it } from 'vitest'
import { render, screen } from '@testing-library/react'
import { ProductCard } from './ProductCard'

describe('ProductCard component', async () => {
  it('should show all required info', () => {
    const mockProduct = {
      id: 1,
      sku: 'JKT-BLCK-1234',
      name: 'Jaket Varsity Hitam',
      price: 199000,
      stock: 137,
      isActive: true,
      createdAt: '2026-08-01T00:00:00Z',
    }

    render(
      <ProductCard
        id={mockProduct.id}
        sku={mockProduct.sku}
        name={mockProduct.name}
        price={mockProduct.price}
        stock={mockProduct.stock}
        isActive={mockProduct.isActive}
        createdAt={mockProduct.createdAt}
        lastMovementAt={null}
      />
    )

    expect(screen.getByText(/JKT-BLCK-1234/i)).toBeInTheDocument()
    expect(screen.getByText(/Jaket Varsity Hitam/i)).toBeInTheDocument()
    expect(screen.getByText(/Rp199\.000/i)).toBeInTheDocument()
    expect(screen.getByText(/137 left/i)).toBeInTheDocument()
    expect(screen.getByText(/Sat Aug 01 2026/i)).toBeInTheDocument()
    expect(screen.getByText(/No movement yet/i)).toBeInTheDocument()    
    expect(screen.getByText(/Active/i)).toBeInTheDocument()    
  })
})
