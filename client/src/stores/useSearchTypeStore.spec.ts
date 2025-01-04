import { describe, it, expect, beforeEach } from 'vitest'
import { useSearchTypeStore } from './useSearchTypeStore'

describe('useSearchTypeStore', () => {
  let store: ReturnType<typeof useSearchTypeStore>

  beforeEach(() => {
    store = useSearchTypeStore()
  })

  it('should initialize store correctly', () => {
    expect(store.isEmailSearchActive).toBe(true)
    expect(store.isPersonSearchActive).toBe(false)
    expect(store.searchType).toBe('emails')
    expect(store.searchFieldActive).toBe('')
    expect(store.searchTerm).toBe('')
  })

  it('should toggle search type to emails', () => {
    store.toggleSearchType('emails')
    expect(store.isEmailSearchActive).toBe(true)
    expect(store.isPersonSearchActive).toBe(false)
  })

  it('should toggle search type to persons', () => {
    store.toggleSearchType('persons')
    expect(store.isEmailSearchActive).toBe(false)
    expect(store.isPersonSearchActive).toBe(true)
  })
})
