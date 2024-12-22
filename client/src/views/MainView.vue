<script setup lang="ts">
import ExplorerContainer from '@/components/ExplorerContainer.vue'
import ViewerContainer from '@/components/ViewerContainer.vue'
</script>

<template>
  <div class="lg:h-full md:overflow-y-hidden w-full flex flex-col lg:flex-row">
    <div class="lg:w-7/12">
      <ExplorerContainer />
    </div>
    <div class="lg:w-5/12">
      <ViewerContainer />
    </div>
  </div>
</template>
