<script setup lang="ts">
import { useEmailTableStore } from '@/stores/useEmailTableStore'
import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'
import { onBeforeMount, computed } from 'vue'
import { storeToRefs } from 'pinia'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { useHighlight } from '@/composables/useHighlight'
const { highlightText } = useHighlight()

import Pagination from '@/components/ExplorerDataTablePagination.vue'
import DataNotFoundBanner from '@/components/DataNotFoundBanner.vue'

library.add(fas)

const emailStore = useEmailTableStore()
const { emailList, pageNumber, pageSize, totalPage, sortOrder, sortField } = storeToRefs(emailStore)
const { fetchEmails, sortEmailsByField, setNextPage, setPreviousPage } = useEmailTableStore()

const { fetchEmail } = useEmailViewerStore()

const { setSelectedItemType } = useItemSelectedStore()

const searchTypeStore = useSearchTypeStore()
const { searchTerm } = storeToRefs(searchTypeStore)

const showEmailDetail = (emailId: string) => {
  fetchEmail(emailId)
  setSelectedItemType('email')
}

onBeforeMount(async () => {
  fetchEmails()
});

const tableHeaders = [
  { field: 'date', label: 'Date' },
  { field: 'from', label: 'From' },
  { field: 'to', label: 'To' },
  { field: 'subject', label: 'Subject' },
]

const highlightedEmails = computed(() => {
  const term = searchTerm.value || '';
  return emailList.value.map(email => {
    return {
      ...email,
      date: highlightText(email.date, term),
      from: highlightText(email.from, term),
      subject: highlightText(email.subject, term),
      toArray: email.toArray?.map(to => highlightText(to, term)) || [],
    };
  });
});

</script>

<template>

  <DataNotFoundBanner v-if="emailList.length === 0" />

  <main v-if="emailList.length !== 0" class="h-full table-container rounded-md">
    <div class="h-full flex flex-col items-start overflow-hidden border rounded-lg">

      <!-- Contenedor con scroll para las filas -->
      <div class="flex-grow overflow-y-auto custom-scrollbar">
        <table class="w-full overflow-x-auto min-w-full table-fixed  text-sm">

          <thead class="bg-gray-100 border-b sticky top-0 z-10">
            <tr>
              <th class="w-7 pl-3 py-2 text-top cursor-pointer whitespace-nowrap">#</th>
              <th @click="sortEmailsByField(header.field)" v-for="header in tableHeaders" :key="header.field"
                class="px-2 py-2 text-left cursor-pointer whitespace-nowrap">
                {{ header.label }}
                <font-awesome-icon icon="arrow-up" v-if="sortField == header.field && sortOrder == 'asc'" />
                <font-awesome-icon icon="arrow-down" v-if="sortField == header.field && sortOrder == 'desc'" />
                <span v-if="sortField != header.field"> ↕ </span>
              </th>
            </tr>
          </thead>
          <tbody class="text-gray-600 ">
            <tr @click="showEmailDetail(email.id)" v-for="(email, index) in highlightedEmails" :key="email.id"
              class="border-b hover:bg-gray-50 cursor-pointer">
              <td class="overflow-x-hidden w-7 pl-3 py-2 align-top">{{ index + 1 }}</td>
              <td class="overflow-x-hidden px-2 py-2 align-top flex-nowrap" v-html="email.date"></td>
              <td class="overflow-x-hidden px-2 py-2 align-top" v-html="email.from"></td>
              <td class="overflow-x-hidden px-2 py-2 align-top">
                <span v-if="email?.toArray && email?.toArray.length > 0 && email?.toArray[0] !== ''"
                  class="my-1 block max-h-32 overflow-x-hidden">
                  <span>
                    <span v-for="(emailAddress, index) in email.toArray.slice(0, 2)" :key="index">
                      <span v-html="emailAddress + '</br>'"></span>
                    </span>
                    <span v-if="email.toArray.length > 2" class="opacity-60">... Ver mas</span>
                  </span>
                </span>
                <span v-else class="text-xs flex pt-1 ml-1">N/A</span>
              </td>
              <td class="overflow-x-hidden px-2 py-2 align-top" v-html="email.subject"></td>
            </tr>
          </tbody>
        </table>
      </div>

      <Pagination class="h-[8%] flex-shrink-0" :currentPage="pageNumber" :totalPages="totalPage" :pageSize="pageSize"
        @prevPage="setPreviousPage" @nextPage="setNextPage" />

    </div>
  </main>

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
