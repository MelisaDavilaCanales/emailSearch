<script setup lang="ts">
import ExplorerContainer from '@/components/ExplorerContainer.vue'
import ViewerContainer from '@/components/ViewerContainer.vue'
</script>

<template>
  <div class="md:h-full md:overflow-y-hidden w-full flex flex-col md:flex-row">
    <div class="md:w-3/5">
      <ExplorerContainer />
    </div>
    <div class="md:w-2/5">
      <ViewerContainer />
    </div>
  </div>
</template>