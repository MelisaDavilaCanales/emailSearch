<script setup lang="ts">
import ExplorerContainer from '@/components/ExplorerContainer.vue'
import ViewerContainer from '@/components/ViewerContainer.vue'
</script>

<template>
  <div class="lg:h-full md:overflow-y-hidden w-full flex flex-col lg:flex-row">
    <div class="lg:w-3/5">
      <ExplorerContainer />
    </div>
    <div class="lg:w-2/5">
      <ViewerContainer />
    </div>
  </div>
</template>
