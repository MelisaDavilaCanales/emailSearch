<script setup lang="ts">

import { defineProps, defineEmits } from 'vue';

interface Props {
  emailList: string[] | undefined
  highlightedEmails: string[]
  isAllSentEmailsVisible: boolean
}

const emit = defineEmits(['showPersonDetail', 'toggleShowAllEmails'])

const { emailList, highlightedEmails, isAllSentEmailsVisible } = defineProps<Props>();

const handlerShowPersonDetail = (emailAddress: string) => {
  emit('showPersonDetail', emailAddress)
}

const handlerToggleShowAllEmails = (value: boolean) => {
  emit('toggleShowAllEmails', value)
}

</script>

<template>
  <div>
    <div v-if="emailList">
      <!-- First items -->
      <span
        class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 bg-gray-300 hover:bg-gray-400  inline-block whitespace-nowrap mx-1 mb-1"
        v-for="(emailAddress, index) in highlightedEmails.slice(0, 7)" :key="index"
        @click="handlerShowPersonDetail(emailAddress)">
        <span v-html="emailAddress"></span>
        <span v-if="index < emailList.length - 1">,</span>
      </span>
      <!-- Option to show remaining items -->
      <span @click="handlerToggleShowAllEmails(true)" v-if="emailList.length > 7"
        class="opacity-60 cursor-pointer hover:opacity-100" :class="{ 'hidden': isAllSentEmailsVisible }">...
        See more</span>
      <!-- Remaining items -->
      <span v-if="isAllSentEmailsVisible && emailList.length > 7">
        <div
          class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 bg-gray-300 hover:bg-gray-400  inline-block whitespace-nowrap mx-1 mb-1"
          v-for="(emailAddress, index) in highlightedEmails.slice(7,)" :key="index"
          @click="handlerShowPersonDetail(emailAddress)">
          <span v-html="emailAddress"></span>
          <span v-if="index < emailList.length - 8">,</span>
        </div>
      </span>
      <!-- Option to hide remaining items -->
      <span @click="handlerToggleShowAllEmails(false)" v-if="isAllSentEmailsVisible && emailList.length > 7"
        class="opacity-60 cursor-pointer hover:opacity-100">
        See less
      </span>
    </div>
  </div>
</template>
