<script setup lang="ts">

import { useEmailTableStore } from '@/stores/useEmailTableStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { onBeforeMount } from 'vue'
import { storeToRefs } from 'pinia'

import Pagination from '@/components/ExplorerDataTablePagination.vue'

const emailStore = useEmailTableStore()
const { emailList, pageNumber, pageSize, tatalPages, } = storeToRefs(emailStore)

const { fetchEmails, fetchEmail, sortEmailsByField, setNextPage, setPreviousPage } = useEmailTableStore()
const { setSelectedItemType } = useItemSelectedStore()


const showEmailDetail = (emailId: string) => {
  fetchEmail(emailId)
  setSelectedItemType('email')
}

onBeforeMount(async () => {
  fetchEmails()
});
</script>

<template>
  <main class="h-full overflow-auto table-container rounded-md">
    <div class="table-wrapper h-full flex items-center">
      <div class="overflow-y-auto overflow-x-hidden max-h-full w-full border rounded-lg custom-scrollbar">
        <table ref="dataTable" class="min-w-full table-auto border-collapse bg-white text-sm">
          <thead class="bg-gray-100 sticky top-0 z-10">
            <tr>
              <th class="pl-3 py-2 text-top cursor-pointer whitespace-nowrap">#</th>
              <th @click="sortEmailsByField('date')" class="px-2 py-2 text-left cursor-pointer whitespace-nowrap">Date ↕
              </th>
              <th @click="sortEmailsByField('from')" class="px-2 py-2 text-left cursor-pointer whitespace-nowrap">From ↕
              </th>
              <th @click="sortEmailsByField('to')" class="px-2 py-2 text-left cursor-pointer whitespace-nowrap">To ↕
              </th>
              <th @click="sortEmailsByField('subject')" class="px-2 py-2 text-left cursor-pointer whitespace-nowrap">
                Subject ↕</th>
            </tr>
          </thead>
          <tbody class="text-gray-600">
            <tr @click="showEmailDetail(email.id)" v-for="(email, index) in emailList" :key="email.id"
              class="border-t hover:bg-gray-50 cursor-pointer">
              <td class="pl-3 py-2 align-top">{{ index + 1 }}</td>
              <td class="px-2 py-2 align-top whitespace-nowrap">{{ email.day }} {{ email.time }}</td>
              <td class="px-2 py-2 align-top">{{ email.from }}</td>
              <td class="px-2 py-2 align-top">
                <span v-if="email?.toArray && email?.toArray.length > 0 && email?.toArray[0] !== ''"
                  class="my-1 block max-h-32 overflow-x-hidden overflow-y-auto custom-scrollbar">
                  <span v-for="(emailAddress, index) in email.toArray" :key="index">
                    {{ emailAddress }}<br />
                  </span>
                </span>
                <span v-else class="text-xs flex pt-1 ml-1">N/A</span>
              </td>
              <td class="px-2 py-2 align-top">{{ email.subject }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>

  <!-- <Pagination :currentPage="pageNumber" :totalPages="tatalPages" :pageSize="pageSize" /> -->

  <Pagination :currentPage="pageNumber" :totalPages="tatalPages" :pageSize="pageSize" @prevPage="setPreviousPage"
    @nextPage="setNextPage" />

</template>

<style scoped>
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
