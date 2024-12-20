<script setup lang="ts">
import ViewerNavigation from '@/components/ViewerNavigation.vue'
import ViewerDisplay from '@/components/ViewerDisplay.vue'
</script>

<template>
  <div class="h-screen flex flex-col items-center justify-center md:h-full  py-4 px-4 space-y-2">
    <ViewerNavigation></ViewerNavigation>
    <ViewerDisplay></ViewerDisplay>
  </div>
</template>
