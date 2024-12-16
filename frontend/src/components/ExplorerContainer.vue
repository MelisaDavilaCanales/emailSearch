<script setup lang="ts">


import ExplorerSearcher from '@/components/ExplorerSearcher.vue'
import ExplorerFieldSelector from '@/components/ExplorerFieldSelector.vue'
import ExplorerDataTable from '@/components/ExplorerDataTable.vue'
import ExplorerDataTablePagination from '@/components/ExplorerDataTablePagination.vue'

import { useSearchTypeStore } from '@/stores/useSearchTypeStore'
const { existsSearchData } = useSearchTypeStore()
</script>

<template>
  <div class="md:h-full w-full p-4 overflow-y-hidden flex flex-col space-y-3 items-start justify-start md:pl-6 md:pr-0">
    <ExplorerSearcher></ExplorerSearcher>
    <ExplorerFieldSelector></ExplorerFieldSelector>
    <div :class="{ 'w-full max-w-full': existsSearchData, 'w-full max-w-full flex-1': !existsSearchData }"
      class="max-h-[80%]">
      <ExplorerDataTable></ExplorerDataTable>
      <ExplorerDataTablePagination v-if="existsSearchData"></ExplorerDataTablePagination>
    </div>
  </div>
</template>