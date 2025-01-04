import { describe, it, expect, beforeEach, vi } from 'vitest'
import { useEmailTableStore } from './useEmailTableStore'

describe('useEmailTableStore', () => {
  let store: ReturnType<typeof useEmailTableStore>

  beforeEach(() => {
    store = useEmailTableStore()
  })

  it('should initialize with correct default values', () => {
    expect(store.searchParam).toBe('')
    expect(store.pageNumber).toBe(1)
    expect(store.pageSize).toBe(10)
    expect(store.totalPage).toBe(0)
    expect(store.sortField).toBe('date')
    expect(store.sortOrder).toBe('desc')
    expect(store.fetchEmailsError.status).toBe(false)
    expect(store.serverError.status).toBe(false)
  })

  it('should update search parameters correctly', () => {
    store.setEmailSearchParams('from', 'charles')
    expect(store.searchParam).toBe('&field=from&term=charles')
    expect(store.pageNumber).toBe(1)
  })

  it('should toggle sort order correctly', () => {
    store.sortEmailsByField('from')
    expect(store.sortField).toBe('from')
    expect(store.sortOrder).toBe('asc')
    store.sortEmailsByField('from')
    expect(store.sortOrder).toBe('desc')
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
    await store.fetchEmails()
    expect(store.fetchEmailsError.status).toBe(true)
    expect(store.fetchEmailsError.message).toBe('Fetch error')
  })

  it('should handle server errors correctly', async () => {
    global.fetch = vi.fn().mockRejectedValueOnce(new Error('Internal Server Error'))
    await store.fetchEmails()
    expect(store.serverError.status).toBe(true)
    expect(store.serverError.message).toBe('Internal Server Error')
  })
})
