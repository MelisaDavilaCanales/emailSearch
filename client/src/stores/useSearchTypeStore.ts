import { ref } from 'vue'
import { defineStore } from 'pinia'


export const useSearchTypeStore = defineStore('searchType', () => {
  const isEmailSearchActive = ref(true)
  const isPersonSearchActive = ref(false)
  const existsSearchData = ref(true)

    function toggleSearchType(value: string) {
        if(value === 'email') {
            isEmailSearchActive.value = true
            isPersonSearchActive.value = false
        } else if(value === 'person') {
            isEmailSearchActive.value = false
            isPersonSearchActive.value = true
        }
    }

    return {
        isEmailSearchActive,
        isPersonSearchActive,
        existsSearchData,
        toggleSearchType
    }
  
})
