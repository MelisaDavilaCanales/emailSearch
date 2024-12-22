<script setup lang="ts">

import { storeToRefs } from 'pinia'
import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { useEmailTableStore } from '@/stores/useEmailTableStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import Pagination from '@/components/ExplorerDataTablePagination.vue'


const emailViewerStore = useEmailViewerStore()

const { setNextPage, setPreviousPage } = emailViewerStore
const { fetchEmail } = useEmailTableStore()
const { emailList, pageSize, pageNumber, totalPages } = storeToRefs(emailViewerStore)

const { setSelectedItemType } = useItemSelectedStore()


const showEmailDetail = (emailId: string) => {
  fetchEmail(emailId)
  setSelectedItemType('email')
}

</script>

<template>
  <div class="h-full w-full bg-grayExtraSoft rounded-md space-y-4 px-4 pt-4">
    <!-- person's data -->
    <div class="flex space-x-2">
      <div class="w-14  flex items-center">
        <img src="../assets/img/person.png" alt="" class="">
      </div>
      <div class="flex flex-col justify-center w-full">
        <p class="text-sm" v-if="emailList.length > 0">{{ emailList[0].from }}</p>
        <p class="text-sm text-gray-500" v-else>No data available</p>
        <hr class="border-t border-grayDark mt-2">
      </div>
    </div>
    <!-- button's section -->
    <div class="flex justify-end space-x-2 pr-4">
      <button class="btn-secondary">
        Enviados
      </button>
      <button class="btn-secondary">
        Recibidos
      </button>
    </div>
    <!-- emails's person -->
    <div class="relative h-[65%] py-2 space-y-2 flex flex-col">
      <!-- Contenedor de correos con scroll -->
      <div class="overflow-y-auto overflow-x-hidden flex-grow space-y-2 custom-scrollbar">
        <!-- Emails -->
        <div class="relative bg-graySoft rounded-md p-2 flex space-x-2 cursor-pointer"
          @click="showEmailDetail(email.id)" v-for="email in emailList" :key="email.id">
          <div class="w-12 px-1">
            <img src="../assets/img/email-png.png" alt="">
          </div>
          <div class="text-sm pr-2">
            <p><span class="font-bold mr-1">Date:</span> {{ email.day + " " + email.time }}</p>
            <p class="block max-h-5 ">
              <span class="font-bold mr-1">To:</span>
              <span v-if="email?.to && email?.to[0] !== ''" class="truncate whitespace-nowrap overflow-hidden">{{
                email.to }}
              </span>
              <span v-else class="text-xs">N/A</span>
            </p>
            <p class="block max-h-5 ">
              <span class="font-bold mr-1">Subject:</span>
              <span v-if="email?.subject && email?.subject[0] !== ''"
                class="truncate whitespace-nowrap overflow-hidden">{{
                  email.subject }}</span>
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
