<script setup lang="ts">
import { storeToRefs } from "pinia";

import ExplorerSearcher from '@/components/ExplorerSearcher.vue'
import ExplorerFieldSelector from '@/components/ExplorerFieldSelector.vue'
import ExplorerDataTable from '@/components/ExplorerDataTable.vue'
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'

const { existsSearchData } = useSearchTypeStore()
const searchTypeStore = useSearchTypeStore()
const { searchType } = storeToRefs(searchTypeStore)
</script>

<template>
  <div class="lg:h-full w-full mx-auto pt-4 px-4 overflow-y-hidden flex flex-col space-y-3 items-start justify-start">
    <ExplorerSearcher></ExplorerSearcher>
    <ExplorerFieldSelector v-if="searchType == 'emails'"></ExplorerFieldSelector>
    <div :class="{ 'w-full max-w-full': existsSearchData, 'w-full max-w-full flex-1': !existsSearchData }"
      class="max-h-[85%]">
      <ExplorerDataTable></ExplorerDataTable>
    </div>
  </div>
</template>
