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
    <input class="w-[90%] border-2 px-2 md:w-4/6" type="text" placeholder="Buscar" />
    <div class="space-x-2">
      <label for="">Search by:</label>

      <input type="radio" id="emails" value="emails" v-model="selectedOption" />
      <label for="emails">Emails</label>

      <input type="radio" id="persons" value="persons" v-model="selectedOption" />
      <label for="persons">Persons</label>
    </div>
  </div>
</template>
