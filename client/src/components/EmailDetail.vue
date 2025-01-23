<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import EmailDetailMailList from '@/components/EmailDetailMailList.vue'
import EmailDetailHeaders from '@/components/EmailDetailHeaders.vue'

import { type IHighlightedEmailDetail } from '@/types/index'
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

const metaDataFields = [
  { field: 'x_folder', tag: 'Folder' },
  { field: 'x_origin', tag: 'Origin' },
  { field: 'x_file_name', tag: 'File Name' },
]

const highlightedEmailDetail = computed<IHighlightedEmailDetail>(() => {
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
  <div class="font-roboto space-y-3 overflow-x-hidden h-full flex flex-col pb-2 text-sm">
    <div class="flex flex-col space-y-2 flex-1 overflow-y-auto">
      <!-- Email main headers -->
      <EmailDetailHeaders :message_id="highlightedEmailDetail.message_id" :date="highlightedEmailDetail.date"
        :from="highlightedEmailDetail.from" :subject="highlightedEmailDetail.subject"
        @showPersonDetail="showPersonDetail" />

      <!-- Email Body -->
      <div class="flex-1 overflow-y-auto m-1 custom-scrollbar">
        <!-- Received emails (TO) -->
        <div class="flex">
          <span class="font-bold flex mr-1">To:</span>
          <div v-if="emailDetail?.toArray && emailDetail?.toArray.length > 0 && emailDetail?.toArray[0] !== ''">
            <EmailDetailMailList :emailList="emailDetail?.toArray" :highlightedEmails="highlightedEmailDetail.toArray"
              :isAllSentEmailsVisible="isAllSentEmailsVisible" @showPersonDetail="showPersonDetail"
              @toggleShowAllEmails="toggleShowAllSentEmails" />
          </div>
          <span v-else class=" ml-1">N/A</span>
        </div>

        <!-- Copied emails (CC) -->
        <div class="flex"
          v-if="emailDetail?.ccArray && highlightedEmailDetail?.ccArray.length > 0 && highlightedEmailDetail?.ccArray[0] !== ''">
          <span class="font-bold flex mr-1">Cc:</span>
          <div class="overflow-x-hidden ">
            <EmailDetailMailList :emailList="emailDetail?.ccArray" :highlightedEmails="highlightedEmailDetail.ccArray"
              :isAllSentEmailsVisible="isAllCopiedEmailsVisible" @showPersonDetail="showPersonDetail"
              @toggleShowAllEmails="toggleShowAllCopiedEmails" />
          </div>
        </div>

        <hr class="border-t mt-2 border-grayDark w-[98%]">

        <!-- Email metadata -->
        <div class="mt-2">
          <p v-for="{ field, tag } in metaDataFields" :key="field" class="">
            <span class="font-bold mr-2">{{ tag }}:</span>
            <span v-html="highlightedEmailDetail?.[field as keyof IHighlightedEmailDetail]"></span>
          </p>
        </div>

        <!-- Email content -->
        <div class="custom-scrollbar mt-2">
          <p class="font-bold sticky top-0 bg-grayExtraSoft z-10">Contenido:</p>
          <p class="pr-2" v-html="highlightedEmailDetail?.content || 'No content available'"></p>
        </div>
      </div>
    </div>
  </div>
</template>
