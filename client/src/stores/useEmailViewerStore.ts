import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'

import type { ICardEmail, IEmail, IRequestError, IServerErrorResponse } from '@/types'
import { useFormatData } from '@/composables/useFormatData'

export const useEmailViewerStore = defineStore('emailViewer', () => {
  const emailDetail = ref<IEmail | null>(null)
  const isEmailDetailLoading = ref<boolean>(false)

  const emailList = ref<ICardEmail[]>([])
  const emailListType = ref<string>('')
  const isEmailListLoading = ref<boolean>(false)
  const isFetchEmailsByDefault = ref<boolean>(false)

  const pageNumber = ref<number>(1)
  const pageSize = ref<number>(0)
  const totalPages = ref<number>(0)
  const searchTerm = ref<string>('')
  const searchField = ref<string>('_all')
  const sortField = ref<string>('date')
  const sortOrder = ref<string>('desc')

  const isAllSentEmailsVisible = ref(false)
  const isAllCopiedEmailsVisible = ref(false)

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

  const params = new URLSearchParams()

  const { formatDate, convertStringListToArray } = useFormatData()

  const emailSearchURL = computed(() => {
    return emailBaseUrl + '?' + query.value
  })

  const query = computed(() => {
    params.set('term', searchTerm.value)
    params.set('field', searchField.value)
    params.set('page', String(pageNumber.value))
    params.set('page_size', String(pageSize.value))
    params.set('sort', sortField.value)
    params.set('order', sortOrder.value)

    return params.toString()
  })

  async function fetchEmails() {
    emailList.value = []
    isEmailListLoading.value = true

    try {
      console.log('Fetch Emails URL: ', emailSearchURL.value)
      const response = await fetch(emailSearchURL.value)

      if (!response.ok) {
        const responseData: IServerErrorResponse = await response.json()
        fetchEmailsError.value = {
          status: true,
          message: responseData.message,
        }
        isEmailListLoading.value = false
        return
      }

      const data = await response.json()
      const emails = data.data.emails

      // If no "from" emails are found, it automatically searches by "to" and then by "cc" to display if any type of emails exist
      if (isFetchEmailsByDefault.value && emails === null) {
        if (searchField.value === 'from') {
          searchField.value = 'to'
          return
        }
        if (searchField.value === 'to') {
          searchField.value = 'cc'
          return
        }
      }

      // It is set to false so that if the user tries to search for a specific type of email, it doesn't search by default again
      if (isFetchEmailsByDefault.value && emails !== null) {
        isFetchEmailsByDefault.value = false
      }

      emailListType.value = searchField.value

      data.data.emails?.forEach((email: ICardEmail) => {
        const formattedDate = formatDate(email.date.toString())

        emailList.value.push({
          id: email.id,
          from: email.from,
          to: email.to,
          subject: email.subject,
          date: formattedDate,
        })
      })

      pageNumber.value = data.data?.page || 0
      pageSize.value = data.data?.page_size || 0
      totalPages.value = data.data?.total_pages || 0

      restoreFetchError()
    } catch (e) {
      console.log('Fetch Emails ERROR: ', e)

      serverError.value = {
        status: true,
        code: 500,
        message: 'Something went wrong',
      }
    } finally {
      isEmailListLoading.value = false
    }
  }

  async function fetchEmail(emailId: string) {
    if (emailDetail.value?.id === emailId) {
      return
    }

    emailDetail.value = null
    isEmailDetailLoading.value = true

    console.log('Fetch Email URL', `${emailBaseUrl}/${emailId}`)
    const response = await fetch(`${emailBaseUrl}/${emailId}`)

    try {
      if (!response.ok) {
        const responseData: IRequestError = await response.json()
        fetchEmailsError.value = {
          status: true,
          message: responseData.message,
        }
        isEmailDetailLoading.value = false

        return
      }

      const data = await response.json()
      const email = data.data

      if (email !== null) {
        const formattedDate = formatDate(email.date.toString())
        const toArray = convertStringListToArray(email.to)
        const ccArray = convertStringListToArray(email.cc)

        emailDetail.value = {
          id: email.id,
          message_id: email.message_id,
          date: formattedDate,
          from: email.from,
          to: email.to,
          toArray: toArray,
          subject: email.subject,
          cc: email.cc,
          ccArray: ccArray,
          x_folder: email.x_folder,
          x_origin: email.x_origin,
          x_file_name: email.x_file_name,
          content: email.content,
        }
      }

      restoreFetchError()
    } catch (e) {
      console.log('Fetch Email ERROR', e)
      serverError.value = {
        status: true,
        code: 500,
        message: 'Something went wrong',
      }
    } finally {
      isEmailDetailLoading.value = false
    }
  }

  function restoreFetchError() {
    fetchEmailsError.value = {
      status: false,
      message: '',
    }
  }

  function setEmailListType(field: string) {
    emailListType.value = field
    searchField.value = field
  }

  function setEmailSearchParams(field: string, term: string) {
    pageNumber.value = 1

    searchTerm.value = term
    searchField.value = field
  }

  function setFetchEmailsListByDefault(value: boolean) {
    isFetchEmailsByDefault.value = value
  }

  const toggleShowAllSentEmails = (value: boolean) => {
    isAllSentEmailsVisible.value = value
  }

  const toggleShowAllCopiedEmails = (value: boolean) => {
    isAllCopiedEmailsVisible.value = value
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

  function setPageSize(size: number) {
    pageSize.value = size
  }

  // Reset isAllSentEmailsVisible and isAllCopiedEmailsVisible when emailDetail changes,
  // so that when a new email is selected the copied and sent emails should be hidden
  watch(emailDetail, () => {
    isAllSentEmailsVisible.value = false
    isAllCopiedEmailsVisible.value = false
  })

  watch(emailSearchURL, fetchEmails)

  return {
    emailDetail,
    emailList,
    emailListType,
    emailSearchURL,
    fetchEmailsError,
    isAllCopiedEmailsVisible,
    isAllSentEmailsVisible,
    isFetchEmailsByDefault,
    isEmailDetailLoading,
    isEmailListLoading,
    pageNumber,
    pageSize,
    searchTerm,
    searchField,
    serverError,
    totalPages,

    setEmailListType,
    setEmailSearchParams,
    setFetchEmailsListByDefault,
    setNextPage,
    setPageSize,
    setPreviousPage,
    fetchEmail,
    fetchEmails,
    toggleShowAllCopiedEmails,
    toggleShowAllSentEmails,
  }
})
