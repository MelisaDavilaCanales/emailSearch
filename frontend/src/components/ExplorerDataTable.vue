<script setup lang="ts">
import EmailTable from './EmailTable.vue';
import PersonTable from './PersonTable.vue';
import { useSearchTypeStore } from '@/stores/useSearchTypeStore';

const { isEmailSearchActive, isPersonSearchActive, existsSearchData } = useSearchTypeStore();

</script>

<template>
  <div :class="{ 'w-full h-full': existsSearchData, 'w-full h-[90%]': !existsSearchData }">
    <EmailTable v-if="isEmailSearchActive" />
    <PersonTable v-if="isPersonSearchActive" />
    <div v-if="!isEmailSearchActive && !isPersonSearchActive" class="h-full">
      <div
        :class="{ 'w-full h-full flex flex-col justify-center items-center': existsSearchData, 'w-full h-full flex flex-col justify-center items-center bg-grayExtraSoft py-12': !existsSearchData }">
        <div class="w-24">
          <img src="../assets/img/not-results.png" alt="" class="">
        </div>
        <h1 class="pt-2 text-2xl text-grayDark font-roboto">[ No results found ]</h1>
      </div>
    </div>
  </div>
</template>
