import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'

import type { ISumaryEmail, IRequestError, IServerErrorResponse } from '@/types/index'
import { useFormatData } from '@/composables/useFormatData'

export const useEmailTableStore = defineStore('emailTable', () => {
  const emailList = ref<ISumaryEmail[]>([])

  const pageNumber = ref<number>(1)
  const pageSize = ref<number>(10)
  const totalPage = ref<number>(0)
  const searchParam = ref<string>('')
  const searchTerm = ref<string>('')
  const searchField = ref<string>('')
  const sortField = ref<string>('date')
  const sortOrder = ref<string>('desc')

  const isEmailsLoading = ref<boolean>(false)

  const fetchEmailsError = ref<IRequestError>({
    status: false,
    message: '',
  })

  const serverError = ref<IServerErrorResponse>({
    status: false,
    code: 0,
    message: '',
  })

  const emailBaseUrl = import.meta.env.VITE_API_URL + '/emails'

  const { formatDate, cleanQuery } = useFormatData()

  const emailSearchURL = computed(() => {
    return emailBaseUrl + cleanedQuery.value
  })

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

  const cleanedQuery = computed(() => {
    return cleanQuery(query.value)
  })

  async function fetchEmails() {
    emailList.value = []
    isEmailsLoading.value = true

    try {
      const response = await fetch(emailSearchURL.value)

      if (!response.ok) {
        const responseData: IServerErrorResponse = await response.json()
        fetchEmailsError.value = {
          status: true,
          message: responseData.message,
        }
        isEmailsLoading.value = false
        return
      }

      const data = await response.json()

      data.data.emails?.forEach((email: ISumaryEmail) => {
        const formattedDate = formatDate(email.date.toString())
        const toArray = email.to.split(',').map((email: string) => email.trim())

        emailList.value.push({
          id: email.id,
          date: formattedDate,
          from: email.from,
          to: email.to,
          toArray: toArray,
          subject: email.subject,
        })
      })

      pageNumber.value = data.data?.page || 0
      pageSize.value = data.data?.page_size || 0
      totalPage.value = data.data?.total_pages || 0

      restoreFetchError()
    } catch {
      serverError.value = {
        status: true,
        code: 500,
        message: 'Something went wrong',
      }
    } finally {
      isEmailsLoading.value = false
    }
  }

  function restoreFetchError() {
    fetchEmailsError.value = {
      status: false,
      message: '',
    }
  }

  function setEmailSearchField(field: string) {
    searchField.value = field
  }

  function setEmailSearchParams(field: string, term: string) {
    pageNumber.value = 1

    searchTerm.value = term
    searchField.value = field
    searchParam.value = '&field=' + field + '&term=' + term
  }

  function sortEmailsByField(field: string) {
    sortField.value = field
    pageNumber.value = 1

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

  watch(query, fetchEmails)

  return {
    emailList,
    fetchEmailsError,
    emailSearchURL,
    isEmailsLoading,
    pageNumber,
    pageSize,
    searchField,
    searchParam,
    serverError,
    sortField,
    sortOrder,
    totalPage,

    fetchEmails,
    setEmailSearchField,
    setEmailSearchParams,
    setNextPage,
    setPageSize,
    setPreviousPage,
    sortEmailsByField,
  }
})
