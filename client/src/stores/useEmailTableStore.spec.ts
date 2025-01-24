import { describe, it, expect, beforeEach, vi } from 'vitest'
import { useEmailTableStore } from './useEmailTableStore'

describe('useEmailTableStore', () => {
  let store: ReturnType<typeof useEmailTableStore>
  let searchParam: string
  let pageNumber: number
  let pageSize: number
  let totalPage: number
  let setEmailSearchParams: (from: string, term: string) => void

  beforeEach(() => {
    store = useEmailTableStore()
    ;({ searchParam, pageNumber } = store)
    ;({ pageSize, totalPage, setEmailSearchParams } = store)
  })

  it('should initialize with correct default values', () => {
    expect(searchParam).toBe('')
    expect(pageNumber).toBe(1)
    expect(pageSize).toBe(10)
    expect(totalPage).toBe(0)
    expect(store.sortField).toBe('date')
    expect(store.sortOrder).toBe('desc')
    expect(store.fetchEmailsError.status).toBe(false)
    expect(store.serverError.status).toBe(false)
  })

  it('should update search parameters correctly', () => {
    setEmailSearchParams('from', 'charles')
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
