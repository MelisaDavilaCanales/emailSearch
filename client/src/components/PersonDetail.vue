<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import Pagination from '@/components/ExplorerDataTablePagination.vue'
import LoadingSpinner from '@/components/LoadingSpinner.vue'
import BannerServerError from '@/components/BannerServerError.vue'

import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { usePersonStore } from '@/stores/usePersonStore'
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'
import { useHighlight } from '@/composables/useHighlight'

const { fetchEmail, setEmailListType, setNextPage, setPreviousPage, setPageSize } = useEmailViewerStore()
const { emailList, isEmailListLoading, pageSize, pageNumber, totalPages, emailListType, fetchEmailsError } = storeToRefs(useEmailViewerStore())

const { selectedPersonEmail } = storeToRefs(usePersonStore())

const { setSelectedItemType } = useItemSelectedStore()

const { searchTerm } = storeToRefs(useSearchTypeStore())

const { highlightText } = useHighlight()


const showEmailDetail = (emailId: string) => {
  fetchEmail(emailId)
  setSelectedItemType('email')
}

const highlightedEmails = computed(() => {
  const term = searchTerm.value || '';
  return emailList.value.map(email => {
    return {
      ...email,
      date: highlightText(email.date, term),
      from: highlightText(email.from, term),
      to: highlightText(email.to, term),
      subject: highlightText(email.subject, term),
    };
  });
});

</script>

<template>
  <div class="h-full w-full bg-grayExtraSoft rounded-md px-4 pt-4 pb-2">
    <!-- person's data -->
    <div class="flex space-x-2">
      <div class="w-14  flex items-center">
        <img src="../assets/img/person.png" alt="" class="">
      </div>
      <div class="flex flex-col justify-center w-full">
        <p class="font-bold">{{ selectedPersonEmail }}</p>
        <hr class="border-t border-grayDark mt-1">
      </div>
    </div>

    <!-- person's emails -->
    <div class="flex justify-between space-x-2 my-1 pr-4">
      <!-- type emails title -->
      <div class="flex items-end pt-6">
        <p class="text-primaryMiddle text-sm font-semibold">
          <button v-if="emailListType == 'from' && emailList.length > 0">- Sent Emails -</button>
          <span v-if="emailListType == 'to' && emailList.length > 0">- Received Emails -</span>
          <span v-if="emailListType == 'cc' && emailList.length > 0">- Copied Emails -</span>
        </p>
      </div>

      <!-- button's section to filter emails by type -->
      <div class="space-x-2">
        <button class="text-white text-xs py-1 px-3 rounded"
          :class="emailListType === 'from' ? 'bg-primary cursor-not-allowed' : 'bg-primarySoft hover:bg-primary/80'"
          :disabled="emailListType === 'from'" @click="setEmailListType('from')">
          Sent
        </button>
        <button class="text-white text-xs py-1 px-3 rounded"
          :class="emailListType === 'to' ? 'bg-primary cursor-not-allowed' : 'bg-primarySoft hover:bg-primary/80'"
          :disabled="emailListType === 'to'" @click="setEmailListType('to')">
          Received
        </button>
        <button class="text-white text-xs py-1 px-3 rounded"
          :class="emailListType === 'cc' ? 'bg-primary cursor-not-allowed' : 'bg-primarySoft hover:bg-primary/80'"
          :disabled="emailListType === 'cc'" @click="setEmailListType('cc')">
          Copied
        </button>
      </div>
    </div>

    <!-- Emails container -->
    <div class="relative h-[75%] w-full py-1 space-y-2 flex flex-col">

      <LoadingSpinner v-if="isEmailListLoading" />
      <BannerServerError v-if="fetchEmailsError.status" :message="fetchEmailsError.message" />

      <!-- Message if emails not exist -->
      <div v-if="!isEmailListLoading && !fetchEmailsError.status && emailList.length === 0"
        class="h-full w-full flex flex-col items-center justify-center text-lg font-roboto text-gray-500">
        <div class="w-1/6 mx-auto opacity-50">
          <img src="../assets/img/emails-not-result-gray.png" alt="email-not-result">
        </div>
        <span v-if="emailListType == 'from'">No sent emails</span>
        <span v-if="emailListType == 'to'">No received emails</span>
        <span v-if="emailListType == 'cc' && emailList.length == 0">No copied emails</span>
      </div>

      <!-- Email List -->
      <div v-if="!isEmailListLoading && emailList.length > 0"
        class="w-full overflow-y-auto overflow-x-hidden flex-grow space-y-2 custom-scrollbar">
        <!-- Email item -->
        <div class="w-full relative bg-graySoft rounded-md py-2 px-2 flex space-x-2 cursor-pointer"
          @click="showEmailDetail(email.id)" v-for="email in highlightedEmails" :key="email.id">
          <!-- Icon -->
          <div class="w-10 pt-1 px-1 flex-shrink-0 lg:hidden xl:block">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
              <path
                d="M64 112c-8.8 0-16 7.2-16 16l0 22.1L220.5 291.7c20.7 17 50.4 17 71.1 0L464 150.1l0-22.1c0-8.8-7.2-16-16-16L64 112zM48 212.2L48 384c0 8.8 7.2 16 16 16l384 0c8.8 0 16-7.2 16-16l0-171.8L322 328.8c-38.4 31.5-93.7 31.5-132 0L48 212.2zM0 128C0 92.7 28.7 64 64 64l384 0c35.3 0 64 28.7 64 64l0 256c0 35.3-28.7 64-64 64L64 448c-35.3 0-64-28.7-64-64L0 128z" />
            </svg>
          </div>

          <!-- content -->
          <div class="w-full text-sm pr-2">
            <p class="block max-h-5">
              <span class="font-bold mr-1">From:</span>
              <span v-if="email?.from && email?.to[0] !== ''" class="truncate whitespace-nowrap overflow-hidden"
                v-html="email.from"></span>
              <span v-else class="text-xs">N/A</span>
            </p>
            <p class="block max-h-5">
              <span class="font-bold mr-1">To:</span>
              <span v-if="email?.to && email?.to[0] !== ''" class="truncate whitespace-nowrap overflow-hidden"
                v-html="email.to"></span>
              <span v-else class="text-xs">N/A</span>
            </p>
            <p class="block max-h-5">
              <span class="font-bold mr-1">Subject:</span>
              <span v-if="email?.subject && email?.subject[0] !== ''" class="truncate whitespace-nowrap overflow-hidden"
                v-html="email.subject"></span>
              <span v-else class="text-xs">N/A</span>
            </p>
          </div>

          <!-- date -->
          <span v-html="email.date" class="absolute top-2 right-4 text-xs text-gray-600 lg:hidden xl:block"></span>

          <div class="absolute top-0 right-0 h-full w-4 bg-graySoft border-r-8 border-grayExtraSoft"></div>
        </div>

      </div>
    </div>

    <Pagination v-if="!fetchEmailsError.status && emailList.length > 0" :currentPage="pageNumber"
      :totalPages="totalPages" :pageSize="pageSize" @prevPage="setPreviousPage" @nextPage="setNextPage"
      @pageSizeChange="setPageSize" />
  </div>
</template>
