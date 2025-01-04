import { describe, it, expect } from 'vitest'
import { useFormatData } from './useFormatData'

describe('useFormatData', () => {
  const { formatDate, convertToArray } = useFormatData()

  describe('formatDate', () => {
    it('should format the date correctly', () => {
      const isoDate = '2023-01-01T12:30:45'
      const formattedDate = formatDate(isoDate)
      expect(formattedDate).toBe('2023-01-01 12:30')
    })

    it('should return the same date when no time is present', () => {
      const isoDate = '2023-01-01T00:00:00'
      const formattedDate = formatDate(isoDate)
      expect(formattedDate).toBe('2023-01-01 00:00')
    })
  })

  describe('convertToArray', () => {
    it('should convert a comma-separated string to an array of emails', () => {
      const emails = 'email1@example.com, email2@example.com, email3@example.com'
      const emailArray = convertToArray(emails)
      expect(emailArray).toEqual(['email1@example.com', 'email2@example.com', 'email3@example.com'])
    })

    it('should trim spaces around emails', () => {
      const emails = ' email1@example.com , email2@example.com , email3@example.com '
      const emailArray = convertToArray(emails)
      expect(emailArray).toEqual(['email1@example.com', 'email2@example.com', 'email3@example.com'])
    })
  })
})
