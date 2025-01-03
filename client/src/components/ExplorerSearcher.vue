<script setup lang="ts">
import { ref, watch } from 'vue'
import { storeToRefs } from 'pinia'

import { useSearchTypeStore } from '@/stores/useSearchTypeStore'
import { useEmailTableStore } from '@/stores/useEmailTableStore'
import { usePersonStore } from '@/stores/usePersonStore'

const { searchType, searchFieldActive, searchTerm } = storeToRefs(useSearchTypeStore())
const { toggleSearchType, setSearchFieldActive } = useSearchTypeStore()

const { setEmailSearchParams } = useEmailTableStore()

const { setPersonSearchParams } = usePersonStore()

const selectedSearchTypeOption = ref(searchType) // Preserve the state across page refreshes.
const currentSearchTerm = ref<string>('');

const previousSearchTerm = ref<string>('');
const previousSearchFieldActive = ref<string>('');
const previousSearchType = ref<string>('');

const searchHandler = () => {
  // Return if search parameters haven't changed.
  if (currentSearchTerm.value === previousSearchTerm.value
    && searchFieldActive.value === previousSearchFieldActive.value
    && searchType.value === previousSearchType.value) {
    return
  }

  previousSearchTerm.value = currentSearchTerm.value
  previousSearchFieldActive.value = searchFieldActive.value
  previousSearchType.value = searchType.value

  if (searchType.value === 'emails') {
    setEmailSearchParams(searchFieldActive.value, currentSearchTerm.value)
  } else if (searchType.value === 'persons') {
    setPersonSearchParams(searchFieldActive.value, currentSearchTerm.value)
  }
}

const handlerClearSearchField = () => {
  setSearchFieldActive('');
}

// Allows other components to highlight the active search term.
watch(currentSearchTerm, (newValue) => {
  searchTerm.value = newValue;
})

watch(selectedSearchTypeOption, (newValue) => {
  toggleSearchType(newValue)
})
</script>

<template>
  <div class="w-full flex flex-col justify-center items-center md:flex-row md:justify-around mx-auto">
    <!-- Search bar -->
    <div
      class="w-full lg:w-2/3 flex items-center max-w-3xl bg-white border border-primarySoft rounded-lg p-1 focus:ring-1 focus:ring-primarySoft">
      <div class="px-1 w-5 mr-1">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
          <path fill="#504e4e"
            d="M416 208c0 45.9-14.9 88.3-40 122.7L502.6 457.4c12.5 12.5 12.5 32.8 0 45.3s-32.8 12.5-45.3 0L330.7 376c-34.4 25.2-76.8 40-122.7 40C93.1 416 0 322.9 0 208S93.1 0 208 0S416 93.1 416 208zM208 352a144 144 0 1 0 0-288 144 144 0 1 0 0 288z" />
        </svg>
      </div>
      <!-- Active search field display -->
      <div v-if="searchFieldActive != ''" class="bg-primarySoft text-white text-xs ml-1 mr-3 py-1 px-2 rounded">
        {{ searchFieldActive }}
        <span @click="handlerClearSearchField"
          class="cursor-pointer text-xs opacity-50 hover:opacity-100 hover:font-bold ml-1"> x
        </span>
      </div>
      <!-- Input field for search term -->
      <input type="text" v-model="currentSearchTerm" placeholder=""
        class="flex-1 text-sm bg-transparent text-gray-700 placeholder-gray-400 focus:outline-none"
        @keyup.enter="searchHandler" />
      <button class="ml-3 bg-primaryMiddle text-white text-xs py-1 px-3 rounded hover:bg-primaryDark focus:outline-none"
        @click="searchHandler">
        Search
      </button>
    </div>

    <!-- Search type selector -->
    <div
      class="lg:w-1/3 flex items-center justify-center mt-2 lg:mt-0 ml-3 space-x-2 xl:space-x-4 bg-primaryExtraSoft px-2 py-2 rounded-lg font-light text-xs">
      <span class="font-normal text-nowrap text-gray-800 leading-none xl:text-sm">
        Search by:
      </span>
      <div class="flex items-center space-x-2">
        <input type="radio" id="emails" value="emails" v-model="selectedSearchTypeOption" />
        <label for="emails" class="font-normal xl:text-sm text-gray-800cursor-pointer leading-none">
          Emails
        </label>
      </div>
      <div class="flex items-center space-x-2">
        <input type="radio" id="persons" value="persons" v-model="selectedSearchTypeOption" />
        <label for="persons" class="font-normal xl:text-sm text-gray-800 cursor-pointer leading-none">
          Persons
        </label>
      </div>
    </div>

  </div>
</template>
