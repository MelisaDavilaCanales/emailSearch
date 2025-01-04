import { describe, it, expect } from 'vitest'
import { useHighlight } from '@/composables/useHighlight'

describe('useHighlight', () => {
  const { highlightText, removeHighlightTags } = useHighlight()

  describe('highlightText', () => {
    it('should highlight the term in the text', () => {
      const text = 'This is a test sentence.'
      const highlightedText = highlightText(text, 'test')
      expect(highlightedText).toBe('This is a <span class="highlight">test</span> sentence.')
    })

    it('should not modify the text if term is not found', () => {
      const text = 'This is a test sentence.'
      const highlightedText = highlightText(text, 'example')
      expect(highlightedText).toBe('This is a test sentence.')
    })
  })

  describe('removeHighlightTags', () => {
    it('should remove highlight tags from the text', () => {
      const highlightedText = 'This is a <span class="highlight">test</span> sentence.'
      const textWithoutHighlight = removeHighlightTags(highlightedText)
      expect(textWithoutHighlight).toBe('This is a test sentence.')
    })

    it('should return the same text if no highlight tags are present', () => {
      const text = 'This is a test sentence.'
      const textWithoutHighlight = removeHighlightTags(text)
      expect(textWithoutHighlight).toBe('This is a test sentence.')
    })
  })
})
