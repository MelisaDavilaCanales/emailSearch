import { ref } from 'vue'
import { defineStore } from 'pinia'
// import { useFetch } from '@vueuse/core'

export interface Email {
  id: string
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
  id: string;
  date: Date;
  day: string;
  time: string;
  from: string;
  to: string;
  toArray: string[];
  subject: string;
}

export const useEmailStore = defineStore('emails', () => {
  const emails = ref<SumaryEmail[]>([]);
  const emailDetail = ref<Email | null>(null);

  function extractDayAndTime(isoDate: string): { day: string, time: string } {
    const [datePart, timePart] = isoDate.split('T');
    const day = datePart; // YYYY-MM-DD
    const time = timePart.split('-')[0]; // HH:mm:ss
    return { day, time };
  }

  async function fetchEmails() {
    const response = await fetch('http://localhost:8080/emails?term=charles&field=&page=1&page_size=50&sort=date&order=asc');
    const data = await response.json();

    if (response.ok) {
      data.data.emails.forEach((email: SumaryEmail) => {
        const { day, time } = extractDayAndTime(email.date.toString());

        const toArray = typeof email.to === 'string' ? email.to.split(',').map((email: string) => email.trim()) : email.to;

        emails.value.push({
          id: email.id,
          date: new Date(email.date),
          from: email.from,
          to: email.to,
          toArray: toArray,
          subject: email.subject,
          day: day,
          time: time,
        });
      });
    } else {
      console.log('Error fetching emails:', response.statusText);
    }
  }

  async function fetchEmail(message_id: string) {
    const response = await fetch(`http://localhost:8080/emails/${message_id}`);
    const data = await response.json();
    console.log("data:" + data);

    if (response.ok) {
      emailDetail.value = data.data;
    } else {
      console.log('Error fetching email:', response.statusText);
    }
  }

  return { emails,  emailDetail, fetchEmails, fetchEmail };
});


