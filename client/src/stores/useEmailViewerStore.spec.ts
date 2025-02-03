import { describe, it, expect, beforeEach, vi } from 'vitest'
import { useEmailViewerStore } from './useEmailViewerStore'

// Mock global fetch for API calls
global.fetch = vi.fn()

describe('useEmailViewerStore', () => {
  let store: ReturnType<typeof useEmailViewerStore>

  beforeEach(() => {
    store = useEmailViewerStore()
    vi.clearAllMocks()
  })

  it('should initialize with correct default values', () => {
    expect(store.pageNumber).toBe(1)
    expect(store.pageSize).toBe(0)
    expect(store.totalPages).toBe(0)
    expect(store.isEmailListLoading).toBe(false)
    expect(store.emailListType).toBe('')
    expect(store.fetchEmailsError.status).toBe(false)
  })

  it('should update search parameters correctly with setEmailSearchParams', () => {
    store.setEmailSearchParams('from', 'charles')
    expect(store.searchTerm).toBe('charles')
    expect(store.searchField).toBe('from')
    expect(store.pageNumber).toBe(1)
  })

  it('should update email list type correctly with setEmailListType', () => {
    store.setEmailListType('to')
    expect(store.emailListType).toBe('to')
  })

  it('should toggle visibility of all sent emails correctly', () => {
    store.toggleShowAllSentEmails(true)
    expect(store.isAllSentEmailsVisible).toBe(true)
    store.toggleShowAllSentEmails(false)
    expect(store.isAllSentEmailsVisible).toBe(false)
  })

  it('should toggle visibility of all copied emails correctly', () => {
    store.toggleShowAllCopiedEmails(true)
    expect(store.isAllCopiedEmailsVisible).toBe(true)
    store.toggleShowAllCopiedEmails(false)
    expect(store.isAllCopiedEmailsVisible).toBe(false)
  })

  it('should handle error correctly if fetchEmails fails', async () => {
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: false,
      json: async () => ({ message: 'Error fetching emails' }),
    } as Response)

    await store.fetchEmails()

    expect(store.fetchEmailsError.status).toBe(true)
    expect(store.fetchEmailsError.message).toBe('Error fetching emails')
    expect(store.isEmailListLoading).toBe(false)
  })

  it('should handle server error during email detail fetching', async () => {
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: false,
      json: async () => ({ message: 'Error fetching email detail' }),
    } as Response)

    await store.fetchEmail('msg123')

    expect(store.fetchEmailsError.status).toBe(true)
    expect(store.fetchEmailsError.message).toBe('Error fetching email detail')
    expect(store.isEmailDetailLoading).toBe(false)
  })
})
