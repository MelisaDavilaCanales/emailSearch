<script setup lang="ts">
import { storeToRefs } from "pinia"
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'
import EmailTable from './EmailTable.vue'
import PersonTable from './PersonTable.vue'

const store = useSearchTypeStore()

const { isEmailSearchActive, isPersonSearchActive } = storeToRefs(store)
</script>

<template>
  <div class="w-full h-[97%]">
    <EmailTable v-if="isEmailSearchActive" />
    <PersonTable v-if="isPersonSearchActive" />
  </div>
</template>
