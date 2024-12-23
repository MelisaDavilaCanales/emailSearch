<script setup lang="ts">
import { ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'

const { toggleSearchType } = useSearchTypeStore()

const storeUserType = useSearchTypeStore()
const { searchType } = storeToRefs(storeUserType)

const selectedOption = ref(searchType)

watch(selectedOption, (newValue) => {
  toggleSearchType(newValue)
})
</script>

<template>
  <div class="w-full flex flex-col justify-center items-center md:flex-row md:justify-around">

    <div
      class="w-2/3 flex items-center max-w-3xl mx-auto bg-white border border-primarySoft rounded-lg p-1 focus:ring-1 focus:ring-primarySoft">
      <div class="px-1">
        <img class="w-6" src="../assets/img/search.png" alt="">
      </div>

      <input type="text" placeholder="Buscar notas"
        class="flex-1 text-sm bg-transparent text-gray-700 placeholder-gray-400 focus:outline-none" />

      <button
        class="ml-3 bg-primaryMiddle text-white text-xs py-1 px-3 rounded hover:bg-primaryDark focus:outline-none">
        Buscar
      </button>
    </div>


    <div
      class="w-1/3 flex items-center justify-center ml-3 space-x-4 bg-primaryExtraSoft px-1 py-2 rounded-lg font-light text-xs">
      <span class="font-normal text-sm text-gray-800 leading-none ">Search by:</span>

      <div class="flex items-center space-x-2">
        <input type="radio" id="emails" value="emails" v-model="selectedOption" />
        <label for="emails" class="font-normal text-sm text-gray-800cursor-pointer leading-none">
          Emails
        </label>
      </div>

      <div class="flex items-center space-x-2">
        <input type="radio" id="persons" value="persons" v-model="selectedOption" />
        <label for="persons" class="font-normal text-sm text-gray-800 cursor-pointer leading-none">
          Persons
        </label>
      </div>
    </div>

  </div>
</template>
