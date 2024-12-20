import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { searchTypeI } from '@/types/search'

export const useSearchTypeStore = defineStore('searchType', () => {
  const searchType = ref<searchTypeI>('emails') // Valor inicial 'emails'

  const isEmailSearchActive = ref(true)
  const isPersonSearchActive = ref(false)
  const existsSearchData = ref(true)

  // Función para cambiar el tipo de búsqueda
  function toggleSearchType(value: searchTypeI) {
    searchType.value = value // Actualiza el tipo de búsqueda

    // Actualizamos el estado reactivo de isEmailSearchActive e isPersonSearchActive
    if (value === 'emails') {
      isEmailSearchActive.value = true
      isPersonSearchActive.value = false
    } else if (value === 'persons') {
      isEmailSearchActive.value = false
      isPersonSearchActive.value = true
    }
  }

  return {
    searchType,
    isEmailSearchActive,
    isPersonSearchActive,
    existsSearchData,
    toggleSearchType
  }
})
