import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'

import type { CardEmailI, Email } from '@/types/search'

function extractDayAndTime(isoDate: string): { day: string; time: string } {
  const [datePart, timePart] = isoDate.split('T')
  const day = datePart
  const time = timePart.split(':').slice(0, 2).join(':') // Extrae solo hh:mm
  return { day, time }
}

export const useEmailViewerStore = defineStore('emailViewer', () => {
  const emailList = ref<CardEmailI[]>([])
  const emailDetail = ref<Email | null>(null)
  const emailListType = ref<string>("from")

  const pageNumber = ref<number>(1)
  const pageSize = ref<number>(0)
  const totalPages = ref<number>(0)
  const searchTerm = ref<string>('')
  const searchField = ref<string>('_all')
  const sortField= ref<string>('date')
  const sortOrder = ref<string>('desc')

  const searchParam = ref<string>('')

  const baseUrl = import.meta.env.VITE_API_BASE_URL

  const emailSearchURL = computed(() => {
    return baseUrl + '/emails' + query.value
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

  async function fetchEmails() {
    const response = await fetch(emailSearchURL.value)
    const data = await response.json()

    emailList.value = []

    if (response.ok) {

      if (data.data.emails === null) {
        if (searchField.value === 'from') {
          searchField.value = 'to'
          searchParam.value = '&field=to&term=' + searchTerm.value
        }

        emailListType.value = searchField.value
      }

      data.data.emails?.forEach((email: CardEmailI) => {
      const { day, time } = extractDayAndTime(email.date.toString())

        const dateFormatted = day + ' ' + time

        emailList.value.push({
          id: email.id,
          from: email.from,
          to: email.to,
          subject: email.subject,
          date: dateFormatted,
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


  async function fetchEmail(message_id: string) {
    const response = await fetch(`http://localhost:8080/emails/${message_id}`)
    const data = await response.json()

    if (response.ok) {
      const email = data.data

      const toArray = email.to
        .split(',')
        .map((email: string) => email.trim())
        .filter((email: string) => email !== '');

      const ccArray = email.cc
        .split(',')
        .map((email: string) => email.trim())
        .filter((email: string) => email !== '');

      emailDetail.value = {
        id: email.id,
        message_id: email.message_id,
        date: email.date,
        from: email.from,
        to: email.to,
        toArray: toArray,
        subject: email.subject,
        cc: email.cc,
        ccArray: ccArray,
        bcc: email.bcc,
        x_folder: email.x_folder,
        x_origin: email.x_origin,
        x_file_name: email.x_file_name,
        content: email.content,
      }

      console.log('Email Details:', emailDetail.value)
    } else {
      console.log('Error fetching email:', response.statusText)
    }
  }

  function setEmailPageNumber(page: number) {
    pageNumber.value = page
  }

  function setEmailPageSize(size: number) {
    pageSize.value = size
  }

  function setEmailSortField(field: string) {
    sortField.value = field
  }

  function setEmailListType(field: string) {
    emailListType.value = field

    searchParam.value = '&field=' + field + '&term=' + searchTerm.value
    searchField.value = field
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

  function setEmailSearchParams(field: string, term: string) { // ### Refactor
    pageNumber.value = 1
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

  watch(emailSearchURL, fetchEmails)
  watch(emailListType, fetchEmails)

  return {
    emailList,
    emailDetail,
    emailSearchURL,
    emailListType,

    pageNumber,
    pageSize,
    totalPages,

    searchTerm,

    setNextPage,
    setPreviousPage,

    fetchEmails,
    fetchEmail,

    setEmailSearchParams,
    setEmailPageNumber,
    setEmailPageSize,
    // setEmailSearchTerm,
    // setEmailSearchField,
    setEmailSortField,

    setEmailListType,

  }
})
