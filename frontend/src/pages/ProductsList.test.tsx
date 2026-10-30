import { render, screen } from '@testing-library/react'
import { describe, expect, it, vi } from 'vitest'
import { ProductsList } from './ProductsList'
import useSWR from 'swr'

vi.mock('swr')
const mockUseSWR = vi.mocked(useSWR)

describe('Products List Page', () => {
  it('should show all products', () => {
    mockUseSWR.mockReturnValue({
      data: {
        data: [
          {
            id: 1,
            sku: 'JKT-BLCK-1234',
            name: 'Jaket Varsity Hitam',
            price: 199000,
            stock: 137,
            isActive: true,
            createdAt: '2026-08-01T00:00:00Z',
          },
          {
            id: 2,
            sku: 'SPT-RED-4567',
            name: 'Sepatu Heels Merah',
            price: 155000,
            stock: 48,
            isActive: true,
            createdAt: '2026-08-05T00:00:00Z',
          },
          {
            id: 3,
            sku: 'JKT-BLUE-1234',
            name: 'Jaket Bomber Biru',
            price: 185000,
            stock: 76,
            isActive: false,
            createdAt: '2026-08-07T00:00:00Z',
          },
        ],
        isLoading: false,
        error: undefined,
        mutate: vi.fn(),
        isValidating: false,
      },
    } as any)

    render(<ProductsList />)

    const productCards = screen.getAllByRole('article')
    expect(productCards.length).toBe(3)
    expect(screen.getAllByText(/Inactive/i).length).toBe(1)
    expect(screen.getAllByText(/^Active/i).length).toBe(2)
  })
})
