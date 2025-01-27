import { describe, it, expect, beforeEach, vi } from 'vitest'
import { usePersonStore } from './usePersonStore'

describe('usePersonStore', () => {
  let store: ReturnType<typeof usePersonStore>

  beforeEach(() => {
    store = usePersonStore()
  })

  it('should initialize store correctly', () => {
    expect(store.personList).toEqual([])
    expect(store.selectedPersonEmail).toBe('')
    expect(store.totalPage).toBe(0)
    expect(store.pageNumber).toBe(1)
    expect(store.pageSize).toBe(0)
    expect(store.searchParam).toBe('')
    expect(store.sortField).toBe('name')
    expect(store.sortOrder).toBe('asc')
    expect(store.isPersonsLoading).toBe(false)
    expect(store.fetchPersonsError).toEqual({ status: false, message: '' })
    expect(store.serverError).toEqual({ status: false, code: 0, message: '' })
  })

  it('should update search parameters correctly', () => {
    store.setPersonSearchParams('name', 'John')
    expect(store.searchParam).toBe('&field=name&term=John')
    expect(store.pageNumber).toBe(1)
  })

  it('should update selected person email correctly', () => {
    store.setSelectedPersonEmail('test@example.com')
    expect(store.selectedPersonEmail).toBe('test@example.com')
  })

  it('should toggle sort order correctly', () => {
    store.sortPersonsByField('name')
    expect(store.sortField).toBe('name')
    expect(store.sortOrder).toBe('desc')
    store.sortPersonsByField('name')
    expect(store.sortOrder).toBe('asc')
  })

  it('should update pagination correctly', () => {
    store.totalPage = 2
    expect(store.pageNumber).toBe(1)
    store.setNextPage()
    expect(store.pageNumber).toBe(2)
    store.setPreviousPage()
    expect(store.pageNumber).toBe(1)
    store.setPageSize(20)
    expect(store.pageSize).toBe(20)
  })

  it('should handle fetch errors correctly', async () => {
    global.fetch = vi.fn().mockResolvedValueOnce({
      ok: false,
      json: async () => ({ message: 'Fetch error' }),
    })
    await store.fetchPersons()
    expect(store.fetchPersonsError.status).toBe(true)
    expect(store.fetchPersonsError.message).toBe('Fetch error')
  })

  it('should handle server errors correctly', async () => {
    global.fetch = vi.fn().mockRejectedValueOnce(new Error('Internal Server Error'))
    await store.fetchPersons()
    expect(store.serverError.status).toBe(true)
    expect(store.serverError.message).toBe('Something went wrong')
  })
})
