import { beforeEach } from 'vitest'
import { createPinia, setActivePinia } from 'pinia'

beforeEach(() => {
  const pinia = createPinia()
  setActivePinia(pinia)
})
