<script setup lang="ts">

import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { usePersonStore } from '@/stores/usePersonStore'
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'

import Pagination from '@/components/ExplorerDataTablePagination.vue'


const emailViewerStore = useEmailViewerStore()
const personStore = usePersonStore()

const { fetchEmail, setEmailListType, setNextPage, setPreviousPage } = useEmailViewerStore()
const { emailList, pageSize, pageNumber, totalPages, emailListType } = storeToRefs(emailViewerStore)
const { selectedPersonEmail } = storeToRefs(personStore)

const { setSelectedItemType } = useItemSelectedStore()

const searchTypeStore = useSearchTypeStore()
const { searchTerm } = storeToRefs(searchTypeStore)

const showEmailDetail = (emailId: string) => {
  fetchEmail(emailId)
  setSelectedItemType('email')
}

const highlightText = (text: string, term: string): string => {
  if (!term) return text;
  const regex = new RegExp(`(${term})`, 'gi');
  return text.replace(regex, '<span class="highlight">$1</span>');
};

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
  <div class="h-full w-full bg-grayExtraSoft rounded-md px-4 pt-4">
    <!-- person's data -->
    <div class="flex space-x-2">
      <div class="w-14  flex items-center">
        <img src="../assets/img/person.png" alt="" class="">
      </div>
      <div class="flex flex-col justify-center w-full">
        <p class="text-sm">{{ selectedPersonEmail }}</p>
        <hr class="border-t border-grayDark mt-2">
      </div>
    </div>
    <!-- button's section -->
    <div class="flex justify-between space-x-2 my-1 pr-4">
      <div class=" flex items-end pt-6">
        <p class="text-primaryMiddle text-sm font-semibold" v-if="emailListType == 'from' && emailList.length > 0"> -
          Correos Enviados - </p>
        <p class="text-primaryMiddle text-sm font-semibold" v-if="emailListType == 'to' && emailList.length > 0"> -
          Correos Recibidos - </p>
        <p class="text-sm text-gray-500" v-if="emailList.length <= 0">No data available</p>
      </div>

      <div class="space-x-2">
        <button class="text-white text-xs py-1 px-3 rounded"
          :class="emailListType === 'from' ? 'bg-primary cursor-not-allowed' : 'bg-primarySoft hover:bg-primaryDark'"
          :disabled="emailListType === 'from'" @click="setEmailListType('from')">
          Enviados
        </button>
        <button class="text-white text-xs py-1 px-3 rounded"
          :class="emailListType === 'to' ? 'bg-primary cursor-not-allowed' : 'bg-primarySoft hover:bg-primaryDark'"
          :disabled="emailListType === 'to'" @click="setEmailListType('to')">
          Recibidos
        </button>
      </div>
    </div>

    <!-- emails's person -->
    <div class="relative h-[75%] w-full py-1 space-y-2 flex flex-col">
      <!-- Contenedor de correos con scroll -->
      <div class="w-full overflow-y-auto overflow-x-hidden flex-grow space-y-2 custom-scrollbar">
        <!-- Emails -->
        <div class="w-full relative bg-graySoft rounded-md py-2 px-2 flex space-x-2 cursor-pointer"
          @click="showEmailDetail(email.id)" v-for="email in highlightedEmails" :key="email.id">
          <div class="w-12 px-1 ">
            <img src="../assets/img/email-png.png" alt="">
          </div>
          <div class="w-full text-sm pr-2">
            <div class="flex justify-between w-full">
              <p class="block max-h-5 ">
                <span class="font-bold mr-1">From:</span>
                <span v-if="email?.from && email?.to[0] !== ''" class="truncate whitespace-nowrap overflow-hidden"
                  v-html="email.from"></span>
                <span v-else class="text-xs">N/A</span>
              </p>
              <span v-html="email.date"></span>
            </div>
            <p class="block max-h-5 ">
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
          <div class="absolute top-0 right-0 h-full w-4 bg-graySoft border-r-8 border-grayExtraSoft"></div>
        </div>
      </div>
    </div>

    <!-- Franja gris fija -->

    <Pagination :currentPage="pageNumber" :totalPages="totalPages" :pageSize="pageSize" @prevPage="setPreviousPage"
      @nextPage="setNextPage" />
  </div>
</template>


<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 10px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f9fafb;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background-color: #d1d5db;
  border-radius: 10px;
}
</style>
