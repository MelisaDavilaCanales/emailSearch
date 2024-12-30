<script setup lang="ts">
import { storeToRefs } from "pinia"
import { useSearchTypeStore } from '@/stores/useSearchTypeStore'
import EmailTable from './EmailTable.vue'
import PersonTable from './PersonTable.vue'

const store = useSearchTypeStore()

const { isEmailSearchActive, isPersonSearchActive, existsSearchData } = storeToRefs(store)
</script>

<template>
  <div :class="{ 'w-full h-full': existsSearchData, 'w-full h-[90%]': !existsSearchData }">
    <!-- Si está activo 'isEmailSearchActive', mostramos la tabla de emails -->
    <EmailTable v-if="isEmailSearchActive" />

    <!-- Si está activo 'isPersonSearchActive', mostramos la tabla de personas -->
    <PersonTable v-if="isPersonSearchActive" />

    <!-- Si no hay resultados de búsqueda, mostramos un mensaje -->
    <!-- <div v-if="!existsSearchData" class="h-full">
      <div
        :class="{ 'w-full h-full flex flex-col justify-center items-center': existsSearchData, 'w-full h-full flex flex-col justify-center items-center bg-grayExtraSoft py-12': !existsSearchData }">
        <div class="w-32">
          <img src="../assets/img/kitty-cuestion.png" alt="" />
        </div>
        <h1 class="pt-2 text-2xl text-grayDark font-roboto">[ No results found ]</h1>
      </div>
    </div> -->
  </div>
</template>
