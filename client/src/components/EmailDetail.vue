<script setup lang="ts">

import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { usePersonStore } from '@/stores/usePersonStore'
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'
import { useHighlight } from '@/composables/useHighlight'

const emailStore = useEmailViewerStore()
const { emailDetail } = storeToRefs(emailStore)
const { setEmailSearchParams, resetEmailExistence, setFetchEmailsListByDefault } = useEmailViewerStore()

const { setSelectedItemType } = useItemSelectedStore()

const { setSelectedPersonEmail } = usePersonStore()

const searchTypeStore = useSearchTypeStore()
const { searchTerm } = storeToRefs(searchTypeStore)

const { highlightText } = useHighlight()

import { useCleanTerm } from '@/composables/useCleanTerm'
const { cleanEmail } = useCleanTerm()



const showPersonDetail = (personEmail: string) => {

  const cleanedEmail: string = cleanEmail(personEmail)

  resetEmailExistence()
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
    <!-- Contenedor general con altura máxima y desplazamiento habilitado -->
    <div class="flex flex-col space-y-2 flex-1 overflow-y-auto">

      <!-- Datos principales del correo -->
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

      <!-- Contenido del correo -->
      <div class="flex-1 overflow-y-auto m-1 custom-scrollbar">
        <!-- Datos de destinatarios -->
        <div class="text-sm flex">
          <span class="font-bold flex mr-1">To:</span>
          <div class="max-h-36 overflow-x-hidden overflow-y-auto custom-scrollbar"
            v-if="emailDetail?.toArray && highlightedEmailDetail?.toArray.length > 0 && highlightedEmailDetail?.toArray[0] !== ''">
            <span
              class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
              v-for="(emailAddress, index) in highlightedEmailDetail?.toArray" :key="index"
              @click="emailAddress && showPersonDetail(emailAddress)">
              <span v-html="emailAddress"></span>
              <span v-if="index < highlightedEmailDetail?.toArray.length - 1">,</span>
            </span>
          </div>
          <span v-else class="text-sm flex pt-1 ml-1">N/A</span>
        </div>

        <!-- Datos de copia -->
        <div class="text-sm flex"
          v-if="emailDetail?.ccArray && highlightedEmailDetail?.ccArray.length > 0 && highlightedEmailDetail?.ccArray[0] !== ''">
          <span class="font-bold flex mr-1">Cc:</span>
          <div class="max-h-36 overflow-x-hidden overflow-y-auto custom-scrollbar">
            <span
              class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
              v-for="(emailAddress, index) in highlightedEmailDetail?.ccArray" :key="index"
              @click="showPersonDetail(emailAddress)">
              <span v-html="emailAddress"></span>
              <span v-if="index < highlightedEmailDetail?.ccArray.length - 1">,</span>
            </span>
          </div>
        </div>

        <hr class="border-t mt-2 border-grayDark">

        <!-- Metadata -->
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

        <!-- Contenido -->
        <div>
          <p class="font-bold sticky top-0 bg-grayExtraSoft z-10">Contenido:</p>
          <p class="text-sm py-1" v-html="highlightedEmailDetail?.content || 'No content available'"></p>
        </div>
      </div>
    </div>
  </div>
</template>



<style scoped>
.highlight {
  background-color: yellow;
  font-weight: bold;
}

.custom-scrollbar::-webkit-scrollbar {
  width: 8px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f9fafb;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background-color: #d1d5db;
  border-radius: 10px;
}
</style>
