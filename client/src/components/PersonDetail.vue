<script setup lang="ts">

import { storeToRefs } from 'pinia'
import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import Pagination from '@/components/ExplorerDataTablePagination.vue'


const emailViewerStore = useEmailViewerStore()

const { setNextPage, setPreviousPage } = emailViewerStore
const { emailList, pageSize, pageNumber, totalPages } = storeToRefs(emailViewerStore)

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
      <div class="overflow-y-auto overflow-x-hidden flex-grow space-y-2">
        <!-- Emails -->
        <div class="relative bg-graySoft rounded-md p-2 flex space-x-2" v-for="email in emailList" :key="email.id">
          <div class="w-12">
            <img src="../assets/img/email.png" alt="">
          </div>
          <div class="text-sm pr-2">
            <p><span class="font-bold">Date:</span> {{ email.day + " " + email.time }}</p>
            <p class="block max-h-4 ">
              <span class="font-bold">To:</span>
              <span class="truncate whitespace-nowrap overflow-hidden">{{ email.to }}</span>
            </p>
            <p class="block max-h-4 ">
              <span class="font-bold">Subject:</span>
              <span class="truncate whitespace-nowrap overflow-hidden">{{ email.subject }}</span>
            </p>
          </div>
          <div class="absolute top-0 right-0 h-full w-4 bg-graySoft"></div>
        </div>
      </div>
    </div>

    <!-- Franja gris fija -->

    <Pagination :currentPage="pageNumber" :totalPages="totalPages" :pageSize="pageSize" @prevPage="setPreviousPage"
      @nextPage="setNextPage" />
  </div>
</template>
