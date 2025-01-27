<script setup lang="ts">
import { onBeforeMount, computed } from 'vue'
import { storeToRefs } from 'pinia'

import Pagination from '@/components/ExplorerDataTablePagination.vue'
import BannerDataNotFound from '@/components/BannerDataNotFound.vue'
import BannerServerError from '@/components/BannerServerError.vue'
import LoadingSpinner from '@/components/LoadingSpinner.vue'

import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { usePersonStore } from '@/stores/usePersonStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'

import { library } from '@fortawesome/fontawesome-svg-core'
import { faUser } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { useHighlight } from '@/composables/useHighlight'

const {
  personList,
  isPersonsLoading,
  pageNumber,
  pageSize,
  totalPage,
  sortOrder,
  sortField,
  fetchPersonsError,
} = storeToRefs(usePersonStore())
const { fetchPersons, sortPersonsByField, setSelectedPersonEmail, setNextPage, setPreviousPage, setPageSize } =
  usePersonStore()

const { setSelectedItemType } = useItemSelectedStore()

const { setEmailSearchParams, setFetchEmailsListByDefault } = useEmailViewerStore()

const { searchTerm } = storeToRefs(useSearchTypeStore())

const { highlightText, removeHighlightTags } = useHighlight()

library.add(faUser)

const showPersonDetail = (personEmail: string) => {
  const cleanedEmail: string = removeHighlightTags(personEmail)

  setFetchEmailsListByDefault(true)
  setEmailSearchParams('from', cleanedEmail)
  setSelectedPersonEmail(cleanedEmail)
  setSelectedItemType('person')
}

const handleRowClick = (personEmail: string) => {
  showPersonDetail(personEmail)

  const section = document.getElementById('viewer');
  if (section) {
    section.scrollIntoView({ behavior: 'smooth' });
  }
}

const tableHeaders = [
  { field: 'email', label: 'Email' },
  { field: 'name', label: 'Name' },
]

const highlightedPersons = computed(() => {
  const term = searchTerm.value || ''
  return personList.value.map((person) => {
    return {
      ...person,
      email: highlightText(person.email, term),
      name: highlightText(person.name, term),
    }
  })
})

onBeforeMount(async () => {
  fetchPersons()
})
</script>

<template>
  <div class="mt-9 h-[93%] w-full">
    <LoadingSpinner v-if="isPersonsLoading" />
    <BannerServerError v-if="fetchPersonsError.status" :message="fetchPersonsError.message" />
    <BannerDataNotFound v-if="!isPersonsLoading && personList.length === 0 && !fetchPersonsError.status" />

    <div v-if="!isPersonsLoading && personList.length !== 0" class="h-full table-container rounded-md">
      <div class="h-full flex flex-col items-start overflow-hidden border rounded-lg">
        <!-- table container -->
        <div class="flex-grow overflow-y-auto custom-scrollbar">
          <table class="w-full overflow-x-auto min-w-full table-fixed text-sm">
            <thead class="bg-gray-100 border-b sticky top-0 z-10">
              <tr>
                <th class="w-7 px-2 py-2 text-center cursor-pointer">#</th>
                <th @click="sortPersonsByField(header.field)" v-for="header in tableHeaders" :key="header.field"
                  class="overflow-x-hidden px-2 py-2 text-left cursor-pointer whitespace-nowrap">
                  <div class="inline-flex items-center justify-center space-x-1">
                    <span>{{ header.label }}</span>

                    <span v-if="sortField == header.field && sortOrder == 'asc'">
                      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512" class="size-3">
                        <path fill="#000000"
                          d="M214.6 41.4c-12.5-12.5-32.8-12.5-45.3 0l-160 160c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0L160 141.2 160 448c0 17.7 14.3 32 32 32s32-14.3 32-32l0-306.7L329.4 246.6c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3l-160-160z" />
                      </svg>
                    </span>

                    <span v-if="sortField == header.field && sortOrder == 'desc'">
                      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512" class="size-3">
                        <path fill="#000000"
                          d="M169.4 470.6c12.5 12.5 32.8 12.5 45.3 0l160-160c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L224 370.8 224 64c0-17.7-14.3-32-32-32s-32 14.3-32 32l0 306.7L54.6 265.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l160 160z" />
                      </svg>
                    </span>

                    <span v-if="sortField != header.field"> ↕ </span>
                  </div>
                </th>
              </tr>
            </thead>
            <tbody class="text-gray-600">
              <tr v-for="(person, index) in highlightedPersons" :key="person.id"
                class="border-b hover:bg-gray-100 cursor-pointer" @click="handleRowClick(person.email)">
                <td class="px-2 pl-3 pr-2 align-center align-top">{{ index + 1 }}</td>
                <td class="overflow-x-hidden break-words px-2 py-2 align-top text-nowrap space-x-1">
                  <font-awesome-icon icon="user" />
                  <span v-html="person.email"></span>
                </td>
                <td class="overflow-x-hidden break-words px-2 py-2 align-top" v-html="person.name"></td>
              </tr>
            </tbody>
          </table>
        </div>

        <Pagination :currentPage="pageNumber" :totalPages="totalPage" :pageSize="pageSize" @prevPage="setPreviousPage"
          @nextPage="setNextPage" @pageSizeChange="setPageSize" />
      </div>
    </div>
  </div>
</template>
