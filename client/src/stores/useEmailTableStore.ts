import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'

import type { SumaryEmail } from '@/types/search'

function extractDayAndTime(isoDate: string): { day: string; time: string } {
  const [datePart, timePart] = isoDate.split('T')
  const day = datePart // YYYY-MM-DD
  const time = timePart.split('-')[0] // HH:mm:ss
  return { day, time }
}

export const useEmailTableStore = defineStore('emailTable', () => {
  const emailList = ref<SumaryEmail[]>([])

  const pageNumber = ref<number>(1)
  const pageSize = ref<number>(50)
  const tatalPages = ref<number>(0)
  const searchTerm = ref<string>('')
  const searchField = ref<string>('_all')
  const searchParam = ref<string>('')
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
      searchParam.value +
      '&sort=' +
      sortField.value +
      '&order=' +
      sortOrder.value
    )
  })

  async function fetchEmails() {
    console.log('FETCHING emailSearchURL:', emailSearchURL.value)
    const response = await fetch(emailSearchURL.value)
    const data = await response.json()

    emailList.value = []

    if (response.ok) {
      data.data.emails?.forEach((email: SumaryEmail) => {
        const { day, time } = extractDayAndTime(email.date.toString())
        const dateFormatted = day + ' ' + time

        const toArray = email.to.split(',').map((email: string) => email.trim())

        emailList.value.push({
          id: email.id,
          date: dateFormatted,
          from: email.from,
          to: email.to,
          toArray: toArray,
          subject: email.subject,
        })
      })

      pageNumber.value = data.data.page
      pageSize.value = data.data.page_size
      tatalPages.value = data.data.total_pages

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

  // function setEmailSearchTerm(term: string) {
  //   searchTerm.value = term
  //   console.log('searchTerm:', searchTerm.value)
  // }

  function setEmailSearchField(field: string) {
    searchField.value = field
  }

  function setEmailSearchParams(field: string, term: string) {
    if (field === '' && term !== '') {
      searchTerm.value = term
      searchParam.value = '&field=' + searchField.value + '&term=' + term
      return
    } else if (term === '' && field !== '') {
      searchField.value = field
      searchParam.value = '&field=' + field + '&term=' + searchTerm.value
      return
    }

    searchField.value = field
    searchTerm.value = term
    searchParam.value =  '&field=' + field + '&term=' + term
  }

  function setEmailSortField(field: string) {
    sortField.value = field
  }

  function sortEmailsByField(field: string) {
    sortField.value = field
    pageNumber.value = 1

    if (sortOrder.value == 'asc') {
      sortOrder.value = 'desc'

      return
    }

    sortOrder.value = 'asc'
  }

  function setNextPage() {
    if (pageNumber.value < tatalPages.value) {
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

    sortOrder,
    sortField,

    pageNumber,
    pageSize,
    tatalPages,

    searchTerm,
    searchField,

    setNextPage,
    setPreviousPage,

    fetchEmails,

    setEmailPageNumber,
    setEmailPageSize,
    setEmailSearchParams,
    // setEmailSearchTerm,
    setEmailSearchField,
    setEmailSortField,

    sortEmailsByField,
  }
})

