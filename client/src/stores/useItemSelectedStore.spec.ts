import { describe, it, expect, beforeEach } from 'vitest'
import { useItemSelectedStore } from './useItemSelectedStore'

describe('useItemSelectedStore', () => {
  let store: ReturnType<typeof useItemSelectedStore>

  beforeEach(() => {
    store = useItemSelectedStore()
  })

  it('should have initial values', () => {
    expect(store.isEmailSelected).toBe(false)
    expect(store.isPersonSelected).toBe(false)
    expect(store.isItemSelected).toBe(false)
  })

  it('should set email as selected', () => {
    store.setSelectedItemType('email')
    expect(store.isEmailSelected).toBe(true)
    expect(store.isPersonSelected).toBe(false)
    expect(store.isItemSelected).toBe(true)
  })

  it('should set person as selected', () => {
    store.setSelectedItemType('person')
    expect(store.isPersonSelected).toBe(true)
    expect(store.isEmailSelected).toBe(false)
    expect(store.isItemSelected).toBe(true)
  })

  it('should not set any type when the value is neither "email" nor "person"', () => {
    store.setSelectedItemType('other' as any)
    expect(store.isEmailSelected).toBe(false)
    expect(store.isPersonSelected).toBe(false)
    expect(store.isItemSelected).toBe(false)
  })
})
