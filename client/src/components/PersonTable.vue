<script setup lang="ts">

import { onBeforeMount } from 'vue'
import { storeToRefs } from 'pinia'

import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { usePersonStore } from '@/stores/usePersonStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'


import Pagination from '@/components/ExplorerDataTablePagination.vue'

const { setSelectedItemType } = useItemSelectedStore()

const personStore = usePersonStore()
const { persons, pageNumber, pageSize, tatalPages, } = storeToRefs(personStore)

const { fetchPersons, sortPersonsByField } = usePersonStore()

const { setEmailSearchTerm, setEmailSearchField, } = useEmailViewerStore()


const showPersonDetail = (personEmail: string) => {
  setEmailSearchTerm(personEmail)
  setEmailSearchField('from')

  setSelectedItemType('person')
}

onBeforeMount(async () => {
  fetchPersons()
});

</script>

<template>
  <main class="h-full overflow-auto table-container rounded-md">
    <div class="table-wrapper h-full flex items-center">
      <div class="overflow-y-auto max-h-full w-full border rounded-lg custom-scrollbar">
        <table ref="dataTable" class="min-w-full table-auto border-collapse bg-white text-sm">
          <thead class="bg-gray-100 sticky top-0 z-10">
            <tr>
              <th class="px-2 py-2 text-center cursor-pointer">#</th>
              <th class="pl-2 py-2 text-left cursor-pointer"></th>
              <th @click="sortPersonsByField('email')" class="px-2 py-2 text-left cursor-pointer">Email ↕</th>
              <th @click="sortPersonsByField('name')" class="px-2 py-2 text-left cursor-pointer">Name ↕</th>
            </tr>
          </thead>
          <tbody class="text-gray-600">
            <tr v-for="(person, index) in persons" :key="person.id" class="border-t hover:bg-gray-50 cursor-pointer"
              @click="showPersonDetail(person.email)">
              <td class=" px-2 py-2 align-center">{{ index + 1 }}</td>
              <td class="pl-2 py-2align-top">
                <span class="block max-w-6">
                  <img class="" src="../assets/img/person.png" alt="">
                </span>
              </td>
              <td class="px-2 py-2 align-top">{{ person.email }}</td>
              <td class="px-2 py-2 align-top ">{{ person.name }} </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>

  <Pagination :currentPage="pageNumber" :totalPages="tatalPages" :pageSize="pageSize" />

</template>
