<script setup lang="ts">

import { defineProps, defineEmits } from 'vue'

interface Props {
  message_id: string
  date: string
  from: string
  subject: string
}

const { message_id, date, from, subject } = defineProps<Props>()

const emit = defineEmits(['showPersonDetail'])

const handlerShowPersonDetail = (personEmail: string) => {
  emit('showPersonDetail', personEmail)
}

</script>

<template>
  <div class="w-full overflow-x-hidden flex-none">
    <div>
      <div class="flex justify-between">
        <div>
          <p>
            <span class="font-bold mr-2">Message ID:</span>
            <span v-html="message_id"></span>
          </p>
          <p>
            <span class="font-bold mr-2">Date:</span>
            <span v-html="date"></span>
          </p>
          <p>
            <span class="font-bold mr-2">From:</span>
            <span
              class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 bg-gray-300 hover:bg-gray-400  inline-block whitespace-nowrap"
              @click="handlerShowPersonDetail(from)" v-html="from">
            </span>
          </p>
        </div>
        <div class="-mt-1 mr-2 w-16 lg:hidden xl:block">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
            <path fill="#050505"
              d="M48 64C21.5 64 0 85.5 0 112c0 15.1 7.1 29.3 19.2 38.4L236.8 313.6c11.4 8.5 27 8.5 38.4 0L492.8 150.4c12.1-9.1 19.2-23.3 19.2-38.4c0-26.5-21.5-48-48-48L48 64zM0 176L0 384c0 35.3 28.7 64 64 64l384 0c35.3 0 64-28.7 64-64l0-208L294.4 339.2c-22.8 17.1-54 17.1-76.8 0L0 176z" />
          </svg>
        </div>
      </div>

      <p class="mr-2 mb-1" v-if="subject && subject.trim() !== ''">
        <span class="font-bold">Subject:</span> <span v-html="subject"></span>
      </p>
    </div>
  </div>
</template>
