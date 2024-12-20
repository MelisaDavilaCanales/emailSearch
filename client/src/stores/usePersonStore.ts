import { ref } from 'vue'
import { defineStore } from 'pinia'

export interface Person {
  name: string
  email: Date
}

export const useEmailStore = defineStore('persons', () => {
  const persons = ref<Person[]>([])

  return { persons }
})
