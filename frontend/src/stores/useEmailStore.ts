import { ref } from 'vue'
import { defineStore } from 'pinia'
import { useFetch } from '@vueuse/core'

export interface Email {
  message_id: string
  date: Date
  from: string
  to: string
  subject: string
  cc: string
  mime_version: string
  content_type: string
  content_transfer_encoding: string
  bcc: string
  x_from: string
  x_to: string
  x_cc: string
  x_bcc: string
  x_folder: string
  x_origin: string
  x_file_name: string
  content: string
}

export interface SumaryEmail {
  message_id: string
  date: Date
  from: string
  to: string
  subject: string
}

export const useEmailStore = defineStore('emails', () => {
  const emails = ref<SumaryEmail[]>([])

  function fetchEmails() {
    const { data, isFetching, error } = useFetch<SumaryEmail[]>('http://localhost:8080/searchEmails?page=10&page_size=5&term=charles&field=from&sort=asc')
    data.value?.forEach((email) => {
      emails.value.push({
        message_id: email.message_id,
        date: email.date,
        from: email.from,
        to: email.to,
        subject: email.subject,
      })
    })

    console.log("emails:"+ emails.value)
    console.log("isFetching:" + isFetching.value)
    console.log("error:"+ error.value)

  }

  //2dM3TqgEsjT
  function fetchEmail(message_id: string) {
    const { data } = useFetch<Email>(`http://localhost:8080/emails/${message_id}`)
    return data.value
  }
    
  return { emails, fetchEmails, fetchEmail }
})
