<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

// import BannerServerError from '@/components/BannerServerError.vue'

import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { usePersonStore } from '@/stores/usePersonStore'
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'

import { useHighlight } from '@/composables/useHighlight'

const { emailDetail, isAllSentEmailsVisible, isAllCopiedEmailsVisible } = storeToRefs(useEmailViewerStore())
const { setEmailSearchParams, setFetchEmailsListByDefault, toggleShowAllSentEmails, toggleShowAllCopiedEmails } = useEmailViewerStore()

const { setSelectedItemType } = useItemSelectedStore()

const { setSelectedPersonEmail } = usePersonStore()

const { searchTerm } = storeToRefs(useSearchTypeStore())

const { highlightText, removeHighlightTags } = useHighlight()

const showPersonDetail = (personEmail: string) => {
  const cleanedEmail: string = removeHighlightTags(personEmail)

  setFetchEmailsListByDefault(true)
  setEmailSearchParams('from', cleanedEmail)
  setSelectedPersonEmail(cleanedEmail)
  setSelectedItemType('person')
}

const highlightedEmailDetail = computed(() => {
  const term = searchTerm.value || '';

  return {
    id: highlightText(emailDetail.value?.message_id || '', term),
    message_id: highlightText(emailDetail.value?.message_id || '', term),
    date: highlightText(emailDetail.value?.date || '', term),
    from: highlightText(emailDetail.value?.from || '', term),
    subject: highlightText(emailDetail.value?.subject || '', term),
    toArray: emailDetail.value?.toArray?.map(to => highlightText(to, term)) || [],
    ccArray: emailDetail.value?.ccArray?.map(cc => highlightText(cc, term)) || [],
    x_folder: highlightText(emailDetail.value?.x_folder || '', term),
    x_origin: highlightText(emailDetail.value?.x_origin || '', term),
    x_file_name: highlightText(emailDetail.value?.x_file_name || '', term),
    content: highlightText(emailDetail.value?.content || '', term),
  };
});

</script>

<template>
  <div class="font-roboto space-y-3 overflow-x-hidden h-full flex flex-col pb-2">
    <!-- <BannerServerError v-if="fetchEmailsError.status" :message="fetchEmailsError.message" /> -->

    <div class="flex flex-col space-y-2 flex-1 overflow-y-auto">
      <!-- Email main headers -->
      <div class="w-full overflow-x-hidden flex-none">
        <div>
          <div class="flex justify-between">
            <div>
              <p>
                <span class="font-bold">Message ID:</span>
                <span v-html="highlightedEmailDetail?.message_id"></span>
              </p>
              <p>
                <span class="font-bold">Date:</span>
                <span v-html="highlightedEmailDetail?.date"></span>
              </p>
              <p class="text-sm">
                <span class="font-bold">From:</span>
                <span
                  class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
                  @click="highlightedEmailDetail?.from && showPersonDetail(highlightedEmailDetail.from)"
                  v-html="highlightedEmailDetail?.from">
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

          <p class="text-sm mb-1" v-if="emailDetail?.subject && highlightedEmailDetail?.subject.trim() !== ''">
            <span class="font-bold">Subject:</span> {{ highlightedEmailDetail?.subject }}
          </p>
        </div>
      </div>

      <!-- Email content -->
      <div class="flex-1 overflow-y-auto m-1 custom-scrollbar">

        <!-- Received emails (TO) -->
        <div class="text-sm flex">
          <span class="font-bold flex mr-1">To:</span>
          <div class=""
            v-if="emailDetail?.toArray && highlightedEmailDetail?.toArray.length > 0 && highlightedEmailDetail?.toArray[0] !== ''">
            <!-- First items -->
            <span
              class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
              v-for="(emailAddress, index) in highlightedEmailDetail?.toArray.slice(0, 7)" :key="index"
              @click="emailAddress && showPersonDetail(emailAddress)">
              <span v-html="emailAddress"></span>
              <span v-if="index < highlightedEmailDetail?.toArray.length - 1">,</span>
            </span>
            <!-- Option to show remaining items -->
            <span @click="toggleShowAllSentEmails(true)" v-if="highlightedEmailDetail.toArray.length > 7"
              class="opacity-60 cursor-pointer hover:opacity-100" :class="{ 'hidden': isAllSentEmailsVisible }">...
              See more</span>
            <!-- Remaining items -->
            <span v-if="isAllSentEmailsVisible && highlightedEmailDetail?.toArray.length > 7">
              <div
                class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
                v-for="(emailAddress, index) in highlightedEmailDetail?.toArray.slice(7,)" :key="index"
                @click="emailAddress && showPersonDetail(emailAddress)">
                <span v-html="emailAddress"></span>
                <span v-if="index < highlightedEmailDetail?.toArray.length - 8">,</span>
              </div>
            </span>
            <!-- Option to hide remaining items -->
            <span @click="toggleShowAllSentEmails(false)"
              v-if="isAllSentEmailsVisible && highlightedEmailDetail.toArray.length > 7"
              class="opacity-60 cursor-pointer hover:opacity-100">
              See less</span>
          </div>
          <span v-else class="text-sm flex pt-1 ml-1">N/A</span>
        </div>

        <!-- Copied emails (CC) -->
        <div class="text-sm flex"
          v-if="emailDetail?.ccArray && highlightedEmailDetail?.ccArray.length > 0 && highlightedEmailDetail?.ccArray[0] !== ''">
          <span class="font-bold flex mr-1">Cc:</span>
          <div class="overflow-x-hidden ">
            <div class=""
              v-if="emailDetail?.ccArray && highlightedEmailDetail?.ccArray.length > 0 && highlightedEmailDetail?.ccArray[0] !== ''">
              <!-- First items -->
              <span
                class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
                v-for="(emailAddress, index) in highlightedEmailDetail?.ccArray.slice(0, 7)" :key="index"
                @click="emailAddress && showPersonDetail(emailAddress)">
                <span v-html="emailAddress"></span>
                <span v-if="index < highlightedEmailDetail?.ccArray.length - 1">,</span>
              </span>
              <!-- Option to show remaining items -->
              <span @click="toggleShowAllCopiedEmails(true)" v-if="highlightedEmailDetail.ccArray.length > 7"
                class="opacity-60 cursor-pointer hover:opacity-100" :class="{ 'hidden': isAllCopiedEmailsVisible }">...
                See more</span>
              <!-- Remaining items -->
              <span v-if="isAllCopiedEmailsVisible && highlightedEmailDetail?.ccArray.length > 7">
                <div
                  class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
                  v-for="(emailAddress, index) in highlightedEmailDetail?.ccArray.slice(7,)" :key="index"
                  @click="emailAddress && showPersonDetail(emailAddress)">
                  <span v-html="emailAddress"></span>
                  <span v-if="index < highlightedEmailDetail?.ccArray.length - 8">,</span>
                </div>
              </span>
              <!-- Option to hide remaining items -->
              <span @click="toggleShowAllCopiedEmails(false)"
                v-if="isAllCopiedEmailsVisible && highlightedEmailDetail.ccArray.length > 7"
                class="opacity-60 cursor-pointer hover:opacity-100">
                See less</span>
            </div>
          </div>
        </div>

        <hr class="border-t mt-2 border-grayDark w-[98%]">

        <!-- Email metadata -->
        <div class="mt-2">
          <p class="text-sm" v-if="emailDetail?.x_folder && highlightedEmailDetail?.x_folder.trim() !== ''">
            <span class="font-bold">Folder:</span>
            <span v-html="highlightedEmailDetail?.x_folder"> </span>
          </p>
          <p class="text-sm" v-if="emailDetail?.x_origin && highlightedEmailDetail?.x_origin.trim() !== ''">
            <span class="font-bold">Origin:</span>
            <span v-html="highlightedEmailDetail?.x_origin"> </span>
          </p>
          <p class="text-sm" v-if="emailDetail?.x_file_name && highlightedEmailDetail?.x_file_name.trim() !== ''">
            <span class="font-bold">File Name:</span>
            <span v-html="highlightedEmailDetail?.x_file_name"> </span>
          </p>
        </div>

        <!-- Email content -->
        <div class="custom-scrollbar">
          <p class="font-bold sticky top-0 bg-grayExtraSoft z-10">Contenido:</p>
          <p class="text-sm py-1 pr-2" v-html="highlightedEmailDetail?.content || 'No content available'"></p>
        </div>
      </div>
    </div>
  </div>
</template>
