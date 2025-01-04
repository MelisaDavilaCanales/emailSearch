import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'

import type { IPerson, IRequestError, IServerErrorResponse } from '@/types/index'

export const usePersonStore = defineStore('persons', () => {
  const personList = ref<IPerson[]>([])
  const selectedPersonEmail = ref<string>('')

  const totalPage = ref<number>(0)
  const pageNumber = ref<number>(1)
  const pageSize = ref<number>(0)
  const searchParam = ref<string>('')
  const sortField = ref<string>('name')
  const sortOrder = ref<string>('asc')

  const isPersonsLoading = ref<boolean>(false)

  const fetchPersonsError = ref<IRequestError>({
    status: false,
    message: '',
  })

  const serverError = ref<IServerErrorResponse>({
    status: false,
    code: 0,
    message: '',
  })

  const baseUrl = import.meta.env.VITE_API_URL

  const query = computed(() => {
    return (
      '?' +
      'page=' +
      pageNumber.value +
      '&page_size=' +
      pageSize.value +
      searchParam.value +
      '&sort=' +
      sortField.value +
      '&order=' +
      sortOrder.value
    )
  })

  const personSearchURL = computed(() => {
    return baseUrl + '/persons' + query.value
  })

  async function fetchPersons() {
    personList.value = []
    isPersonsLoading.value = true

    try {
      const response = await fetch(personSearchURL.value)

      if (!response.ok) {
        const responseData: IServerErrorResponse = await response.json()
        fetchPersonsError.value = {
          status: true,
          message: responseData.message,
        }
        isPersonsLoading.value = false
        return
      }

      const data = await response.json()
      if (data && data.data && data.data.persons !== null) {
        data.data.persons.forEach((person: IPerson) => {
          personList.value.push({
            id: person.id,
            name: person.name,
            email: person.email,
          })
        })
      }

      totalPage.value = data.data.total_pages || 0
      pageNumber.value = data.data.page || 0
      pageSize.value = data.data.page_size || 0

      restoreFetchPersonError()
    } catch {
      serverError.value = {
        status: true,
        code: 500,
        message: 'Internal Server Error',
      }
    } finally {
      isPersonsLoading.value = false
    }
  }

  function restoreFetchPersonError() {
    fetchPersonsError.value = {
      status: false,
      message: '',
    }
  }

  function setPersonSearchParams(field: string, term: string) {
    pageNumber.value = 1
    searchParam.value = '&field=' + field + '&term=' + term
  }

  function setSelectedPersonEmail(email: string) {
    selectedPersonEmail.value = email
  }

  function sortPersonsByField(field: string) {
    sortField.value = field

    if (sortOrder.value == 'asc') {
      sortOrder.value = 'desc'
    } else {
      sortOrder.value = 'asc'
    }
  }

  function setNextPage() {
    if (pageNumber.value < totalPage.value) {
      pageNumber.value++
    }
  }

  function setPreviousPage() {
    if (pageNumber.value > 1) {
      pageNumber.value--
    }
  }

  function setPageSize(size: number) {
    pageSize.value = size
  }

  watch(query, fetchPersons)

  return {
    fetchPersonsError,
    isPersonsLoading,
    personList,
    searchParam,
    selectedPersonEmail,
    serverError,
    sortOrder,
    sortField,
    pageNumber,
    pageSize,
    totalPage,

    fetchPersons,
    setNextPage,
    setPageSize,
    setPreviousPage,
    setSelectedPersonEmail,
    setPersonSearchParams,
    sortPersonsByField,
  }
})
