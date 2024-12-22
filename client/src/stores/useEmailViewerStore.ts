import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'

import type { CardEmailI } from '@/types/search'

function extractDayAndTime(isoDate: string): { day: string; time: string } {
  const [datePart, timePart] = isoDate.split('T')
  const day = datePart // YYYY-MM-DD
  const time = timePart.split('-')[0] // HH:mm:ss
  return { day, time }
}

export const useEmailViewerStore = defineStore('emailViewer', () => {
  const emailList = ref<CardEmailI[]>([])

  const pageNumber = ref<number>(1)
  const pageSize = ref<number>(50)
  const totalPages = ref<number>(0)
  const searchTerm = ref<string>('')
  const searchField = ref<string>('_all')
  const sortField= ref<string>('date')
  const sortOrder = ref<string>('desc')

  const baseUrl = import.meta.env.VITE_API_BASE_URL

  const emailSearchURL = computed(() => {
    return baseUrl + '/emails' + query.value
  })

  // http://localhost:8080/emails?page=1&page_size=10&term=charles&field=from&sort=to&order=desc
  const query = computed(() => {
    return (
      '?' +
      'page=' +
      pageNumber.value +
      '&page_size=' +
      pageSize.value +
      '&term=' +
      searchTerm.value +
      '&field=' +
      searchField.value +
      '&sort=' +
      sortField.value +
      '&order=' +
      sortOrder.value
    )
  })

  async function fetchEmails() {
    const response = await fetch(emailSearchURL.value)
    const data = await response.json()

    emailList.value = []

    if (response.ok) {
      data.data.emails?.forEach((email: CardEmailI) => {
        const { day, time } = extractDayAndTime(email.date.toString())

        emailList.value.push({
          id: email.id,
          from: email.from,
          to: email.to,
          subject: email.subject,
          date: email.date,
          day: day,
          time: time,
        })
      })

      pageNumber.value = data.data.page
      pageSize.value = data.data.page_size
      totalPages.value = data.data.total_pages

      console.log('data:', data)
    } else {
      console.log('Error fetching emails:', response.statusText)
    }
  }


  function setEmailPageNumber(page: number) {
    pageNumber.value = page
  }

  function setEmailPageSize(size: number) {
    pageSize.value = size
  }

  function setEmailSearchTerm(term: string) {
    searchTerm.value = term
  }

  function setEmailSearchField(field: string) {
    searchField.value = field
  }

  function setEmailSortField(field: string) {
    sortField.value = field
  }



  function setNextPage() {
    if (pageNumber.value < totalPages.value) {
      pageNumber.value++
    }
  }

  function setPreviousPage() {
    if (pageNumber.value > 1) {
      pageNumber.value--
    }
  }

  watch(emailSearchURL, fetchEmails)

  return {
    emailList,
    emailSearchURL,

    pageNumber,
    pageSize,
    totalPages,

    setNextPage,
    setPreviousPage,

    fetchEmails,

    setEmailPageNumber,
    setEmailPageSize,
    setEmailSearchTerm,
    setEmailSearchField,
    setEmailSortField,

  }
})
