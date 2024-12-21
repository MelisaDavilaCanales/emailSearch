import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'

import type { SumaryEmail, Email } from '@/types/search'

function extractDayAndTime(isoDate: string): { day: string; time: string } {
  const [datePart, timePart] = isoDate.split('T')
  const day = datePart // YYYY-MM-DD
  const time = timePart.split('-')[0] // HH:mm:ss
  return { day, time }
}

export const useEmailStore = defineStore('emails', () => {
  const emailList = ref<SumaryEmail[]>([])
  const emailDetails = ref<Email | null>(null)

  const pageNumber = ref<number>(1)
  const pageSize = ref<number>(50)
  const tatalPages = ref<number>(0)
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
      data.data.emails?.forEach((email: SumaryEmail) => {
        const { day, time } = extractDayAndTime(email.date.toString())
        const toArray = email.to.split(',').map((email: string) => email.trim())

        emailList.value.push({
          id: email.id,
          // date: new Date(email.date),
          date: email.date,
          from: email.from,
          to: email.to,
          toArray: toArray,
          subject: email.subject,
          day: day,
          time: time,
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

  async function fetchEmail(message_id: string) {
    const response = await fetch(`http://localhost:8080/emails/${message_id}`)
    const data = await response.json()

    if (response.ok) {
      const email = data.data
      const toArray = email.to.split(',').map((email: string) => email.trim())

      emailDetails.value = {
        id: email.id,
        message_id: email.message_id,
        date: email.date,
        from: email.from,
        to: email.to,
        toArray: toArray,
        subject: email.subject,
        cc: email.cc,
        mime_version: email.mime_version,
        content_type: email.content_type,
        content_transfer_encoding: email.content_transfer_encoding,
        bcc: email.bcc,
        x_from: email.x_from,
        x_to: email.x_to,
        x_cc: email.x_cc,
        x_bcc: email.x_bcc,
        x_folder: email.x_folder,
        x_origin: email.x_origin,
        x_file_name: email.x_file_name,
        content: email.content,
      }

      console.log('Email Details:', emailDetails.value)
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

  function setEmailSearchTerm(term: string) {
    searchTerm.value = term
  }

  function setEmailSearchField(field: string) {
    searchField.value = field
  }

  function setEmailSortField(field: string) {
    sortField.value = field
  }

  function sortEmailsByField(field: string) {
    setEmailSortField(field)

    if (sortOrder.value == 'asc') {
      sortOrder.value = 'desc'
      return
    }

    sortOrder.value = 'asc'
  }

  watch(emailSearchURL, fetchEmails)

  return {
    emailList,
    emailDetails,
    emailSearchURL,

    pageNumber,
    pageSize,
    tatalPages,

    fetchEmails,
    fetchEmail,

    setEmailPageNumber,
    setEmailPageSize,
    setEmailSearchTerm,
    setEmailSearchField,
    setEmailSortField,

    sortEmailsByField,
  }
})
