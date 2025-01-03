<script setup lang="ts">
import { storeToRefs } from "pinia"

import EmailTable from './EmailTable.vue'
import PersonTable from './PersonTable.vue'

import { useSearchTypeStore } from '@/stores/useSearchTypeStore'

const { isEmailSearchActive, isPersonSearchActive } = storeToRefs(useSearchTypeStore())

</script>

<template>
  <div class="w-full h-[97%]">
    <EmailTable v-if="isEmailSearchActive" />
    <PersonTable v-if="isPersonSearchActive" />
  </div>
</template>
