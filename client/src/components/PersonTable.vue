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

library.add(fas)

const { setSelectedItemType } = useItemSelectedStore()

const personStore = usePersonStore()
const { persons, pageNumber, pageSize, totalPage, sortOrder, sortField } = storeToRefs(personStore)

const { fetchPersons, sortPersonsByField, setSelectedPersonEmail } = usePersonStore()

const { setEmailSearchParams } = useEmailViewerStore()

const searchTypeStore = useSearchTypeStore()
const { searchTerm } = storeToRefs(searchTypeStore)


const showPersonDetail = (personEmail: string) => {

  const cleanedEmail: string = cleanEmail(personEmail)

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
  return persons.value.map(person => {
    return {
      ...person,
      email: highlightText(person.email, term),
      name: highlightText(person.name, term),
    };
  });
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
              <th @click="sortPersonsByField(header.field)" v-for="header in tableHeaders" :key="header.field"
                class="px-2 py-2 text-left cursor-pointer whitespace-nowrap">
                {{ header.label }}
                <font-awesome-icon icon="arrow-up" v-if="sortField == header.field && sortOrder == 'asc'" />
                <font-awesome-icon icon="arrow-down" v-if="sortField == header.field && sortOrder == 'desc'" />
                <span v-if="sortField != header.field"> ↕ </span>
              </th>
            </tr>
          </thead>
          <tbody class="text-gray-600">
            <tr v-for="(person, index) in highlightedPersons" :key="person.id"
              class="border-t hover:bg-gray-50 cursor-pointer" @click="showPersonDetail(person.email)">
              <td class=" px-2 py-2 align-center">{{ index + 1 }}</td>
              <td class="px-2 py-2 align-top text-nowrap">
                <font-awesome-icon icon="user" />
                <span v-html="person.email"></span>
              </td>
              <td class="px-2 py-2 align-top " v-html="person.name"></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>

  <Pagination :currentPage="pageNumber" :totalPages="totalPage" :pageSize="pageSize" />

</template>
