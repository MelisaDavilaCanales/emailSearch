import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { itemSelectedTypeI } from '@/types/search'

export const useItemSelectedStore = defineStore('itemSelected', () => {
    const isEmailSelected = ref(false)
    const isPersonSelected = ref(false)
    const isItemSelected = ref(false)

    function setSelectedItemType(value: itemSelectedTypeI) {
        if(value === 'email') {
          isEmailSelected.value = true
          isItemSelected.value = false
        } else if( value === 'person') {
            isPersonSelected.value = false
            isItemSelected.value = true
        }
    }

    return {
        isEmailSelected,
        isPersonSelected,
        isItemSelected,
        setSelectedItemType
    }
})
