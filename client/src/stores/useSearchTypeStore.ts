import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { searchTypeI } from '@/types/search'

export const useSearchTypeStore = defineStore('searchType', () => {
  const searchType = ref<searchTypeI>('emails')
  const searchFieldActive = ref('')
  const searchTerm = ref<string>('')

  const isEmailSearchActive = ref(true)
  const isPersonSearchActive = ref(false)
  const existsSearchData = ref(false)

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

    function setEmailSearchTerm(term: string) {
    searchTerm.value = term
    console.log('searchTerm:', searchTerm.value)
  }

  return {
    searchType,
    isEmailSearchActive,
    isPersonSearchActive,
    existsSearchData,

    searchFieldActive,
    setSearchFieldActive,

    toggleSearchType,

    searchTerm,
    setEmailSearchTerm
  }
})
