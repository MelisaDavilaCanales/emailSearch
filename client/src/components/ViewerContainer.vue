<script setup lang="ts">
import { storeToRefs } from "pinia"

import PersonDetail from '@/components/PersonDetail.vue'
import EmailDetail from '@/components/EmailDetail.vue'

import { useItemSelectedStore } from '@/stores/useItemSelectedStore'


const itemSelectedStore = useItemSelectedStore()
const { isEmailSelected, isPersonSelected, isItemSelected } = storeToRefs(itemSelectedStore)

</script>

<template>
  <div class="h-screen flex flex-col items-center justify-center md:h-full  py-4 px-4 space-y-2">
    <div :class="{ 'h-[95%] overflow-hidden w-full ': isItemSelected }"
      class="h-[90%] w-full bg-grayExtraSoft rounded-md p-4">
      <PersonDetail v-if="isPersonSelected" />
      <EmailDetail v-if="isEmailSelected" />
      <div v-if="!isPersonSelected && !isEmailSelected"
        class="h-full grayExtraSoft w-full bg-grayExtraSoft flex flex-col justify-center items-center">
        <img src="../assets/img/kitty-hi.png" alt="" class="w-28 md:w-40 pb-4">
        <!-- <div class="h-36 w-36 border-2 border-green-500 "></div> -->
        <h4 class="font-roboto text-2xl text-grayDark ">Visualizer</h4>
        <p class="font-roboto  text-grayDark">[Select an item to display it in detail]</p>
      </div>
    </div>
  </div>
</template>
