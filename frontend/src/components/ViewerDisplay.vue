<script setup lang="ts">
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import PersonDetail from '@/components/PersonDetail.vue'
import EmailDetail from '@/components/EmailDetail.vue'

const { isEmailSelected, isPersonSelected, isItemSelected } = useItemSelectedStore()

</script>

<template>
  <div :class="{ 'h-[90%] overflow-hidden w-full ': isItemSelected }"
    class="h-[90%] w-full bg-grayExtraSoft rounded-md p-4">
    <PersonDetail v-if="isPersonSelected" />
    <EmailDetail v-if="isEmailSelected" />
    <div v-if="!isPersonSelected && !isEmailSelected"
      class="h-full grayExtraSoft w-full bg-grayExtraSoft flex flex-col justify-center items-center">
      <img src="../assets/img/kitty.png" alt="" class="w-28 md:w-40 pb-4">
      <h4 class="font-roboto text-2xl text-grayDark ">Visualizer</h4>
      <p class="font-roboto  text-grayDark">[Select an item to display it in detail]</p>
    </div>
  </div>
</template>
