import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { Person } from '@/stores/usePersonStore'
import type { Email } from '@/stores/useEmailStore'



type validItem = Person | Email

export const useItemSelectedStore = defineStore('itemSelected', () => {
    const isEmailSelected = ref(true)
    const isPersonSelected = ref(false)
    const isItemSelected = ref(true)
    
    function setIteSelected(item: validItem) {
        if(item.hasOwnProperty('email')) {
            isEmailSelected.value = true
            isPersonSelected.value = false
        } else if(item.hasOwnProperty('name')) {
            isEmailSelected.value = false
            isPersonSelected.value = true
        }
    }

    return {
        isEmailSelected,
        isPersonSelected,
        isItemSelected,
        setIteSelected
    }
})