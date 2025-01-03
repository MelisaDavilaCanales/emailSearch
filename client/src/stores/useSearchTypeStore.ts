import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { searchTypeI } from '@/types'

export const useSearchTypeStore = defineStore('searchType', () => {
  const isEmailSearchActive = ref(true)
  const isPersonSearchActive = ref(false)

  const searchType = ref<searchTypeI>('emails')
  const searchFieldActive = ref('')
  const searchTerm = ref<string>('')

  function setEmailSearchTerm(term: string) {
    searchTerm.value = term
  }

  function setSearchFieldActive(value: string) {
    searchFieldActive.value = value
  }

  function toggleSearchType(value: searchTypeI) {
    searchType.value = value

    if (value === 'emails') {
      isEmailSearchActive.value = true
      isPersonSearchActive.value = false
    } else if (value === 'persons') {
      isEmailSearchActive.value = false
      isPersonSearchActive.value = true
    }
  }

  return {
    isEmailSearchActive,
    isPersonSearchActive,
    searchType,
    searchFieldActive,
    searchTerm,

    setEmailSearchTerm,
    setSearchFieldActive,
    toggleSearchType,
  }
})
