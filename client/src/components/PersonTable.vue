<script setup lang="ts">

import { onBeforeMount, computed } from 'vue'
import { storeToRefs } from 'pinia'

import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { usePersonStore } from '@/stores/usePersonStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'

import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { useHighlight } from '@/composables/useHighlight'
const { highlightText } = useHighlight()

import { useCleanTerm } from '@/composables/useCleanTerm'
const { cleanEmail } = useCleanTerm()

import Pagination from '@/components/ExplorerDataTablePagination.vue'
import DataNotFoundBanner from '@/components/DataNotFoundBanner.vue'

library.add(fas)

const { setSelectedItemType } = useItemSelectedStore()

const personStore = usePersonStore()
const { personList, pageNumber, pageSize, totalPage, sortOrder, sortField } = storeToRefs(personStore)

const { fetchPersons, sortPersonsByField, setSelectedPersonEmail, setNextPage, setPreviousPage } = usePersonStore()

const { setEmailSearchParams, resetEmailExistence, setFetchEmailsListByDefault } = useEmailViewerStore()

const searchTypeStore = useSearchTypeStore()
const { searchTerm } = storeToRefs(searchTypeStore)


const showPersonDetail = (personEmail: string) => {

  const cleanedEmail: string = cleanEmail(personEmail)

  resetEmailExistence()
  setFetchEmailsListByDefault(true)
  setEmailSearchParams('from', cleanedEmail)
  setSelectedPersonEmail(cleanedEmail)
  setSelectedItemType('person')
}

onBeforeMount(async () => {
  fetchPersons()
});

const tableHeaders = [
  { field: 'email', label: 'Email' },
  { field: 'name', label: 'Name' },
]


const highlightedPersons = computed(() => {
  const term = searchTerm.value || '';
  return personList.value.map(person => {
    return {
      ...person,
      email: highlightText(person.email, term),
      name: highlightText(person.name, term),
    };
  });
});

</script>

<template>

  <div class="h-full mt-8">

    <DataNotFoundBanner v-if="personList.length === 0" />

    <main v-if="personList.length !== 0" class="h-full table-container rounded-md">
      <div class="h-full flex flex-col items-start overflow-hidden border rounded-lg">

        <div class="flex-grow overflow-y-auto custom-scrollbar">
          <table class="w-full overflow-x-auto min-w-full table-fixed  text-sm">

            <thead class="bg-gray-100 border-b sticky top-0 z-10">
              <tr>
                <th class="w-7 px-2 py-2 text-center cursor-pointer">#</th>
                <th @click="sortPersonsByField(header.field)" v-for="header in tableHeaders" :key="header.field"
                  class="overflow-x-hidden px-2 py-2 text-left cursor-pointer whitespace-nowrap">
                  {{ header.label }}
                  <font-awesome-icon icon="arrow-up" v-if="sortField == header.field && sortOrder == 'asc'" />
                  <font-awesome-icon icon="arrow-down" v-if="sortField == header.field && sortOrder == 'desc'" />
                  <span v-if="sortField != header.field"> ↕ </span>
                </th>
              </tr>
            </thead>
            <tbody class="text-gray-600">
              <tr v-for="(person, index) in highlightedPersons" :key="person.id"
                class="border-b hover:bg-gray-50 cursor-pointer" @click="showPersonDetail(person.email)">
                <td class="px-2 pl-3 pr-2 align-center align-top">{{ index + 1 }}</td>
                <td class="overflow-x-hidden break-words px-2 py-2 align-top text-nowrap space-x-1">
                  <font-awesome-icon icon="user" />
                  <span v-html="person.email"></span>
                </td>
                <td class="overflow-x-hidden break-words  px-2 py-2 align-top " v-html="person.name">
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <Pagination :currentPage="pageNumber" :totalPages="totalPage" :pageSize="pageSize" @prevPage="setPreviousPage"
          @nextPage="setNextPage" />

      </div>
    </main>

  </div>

</template>
