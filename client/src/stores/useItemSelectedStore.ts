import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { itemSelectedTypeI } from '@/types'

export const useItemSelectedStore = defineStore('itemSelected', () => {
  const isEmailSelected = ref(false)
  const isPersonSelected = ref(false)
  const isItemSelected = ref(false)

  function setSelectedItemType(value: itemSelectedTypeI) {
    if (value === 'email') {
      isItemSelected.value = true
      isEmailSelected.value = true
      isPersonSelected.value = false
    } else if (value === 'person') {
      isItemSelected.value = true
      isPersonSelected.value = true
      isEmailSelected.value = false
    }
  }

  return {
    isEmailSelected,
    isPersonSelected,
    isItemSelected,

    setSelectedItemType,
  }
})
