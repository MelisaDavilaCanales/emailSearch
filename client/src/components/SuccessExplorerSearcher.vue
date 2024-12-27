<script setup lang="ts">
import { ref, watch } from 'vue';
import { storeToRefs } from 'pinia';
import { useSearchTypeStore } from '@/stores/useSearchTypeStore';
import { useEmailTableStore } from '@/stores/useEmailTableStore';
import { usePersonStore } from '@/stores/usePersonStore';

const { toggleSearchType, setSearchFieldActive } = useSearchTypeStore();
const { setEmailSearchParams } = useEmailTableStore();
const { setPersonSortField, setPersonSearchTerm } = usePersonStore();

const searchTypeStore = useSearchTypeStore();
const { searchType, searchFieldActive } = storeToRefs(searchTypeStore);
const emailTableStore = useEmailTableStore();
const { searchTerm } = storeToRefs(emailTableStore);

const selectedSearchTypeOption = ref(searchType);
const searchContent = ref<string>('');

const searchHandler = () => {
  if (searchContent.value.trim() === '') return;

  searchTerm.value = searchContent.value;

  if (searchType.value === 'emails') {
    setEmailSearchParams(searchFieldActive.value, searchContent.value);
  } else if (searchType.value === 'persons') {
    setPersonSortField(searchFieldActive.value);
    setPersonSearchTerm(searchContent.value);
  }
};

const handlerClearSearchField = () => {
  setSearchFieldActive('');
  searchContent.value = '';
  searchTerm.value = '';
};

watch(searchContent, (newValue) => {
  searchTerm.value = newValue;
});

watch(selectedSearchTypeOption, (newValue) => {
  toggleSearchType(newValue);
});
</script>

<template>
  <div class="w-full flex flex-col justify-center items-center md:flex-row md:justify-around">
    <div
      class="w-2/3 flex items-center max-w-3xl mx-auto bg-white border border-primarySoft rounded-lg p-1 focus:ring-1 focus:ring-primarySoft">
      <div class="px-1">
        <img class="w-6" src="../assets/img/search.png" alt="Search Icon">
      </div>

      <div v-if="searchFieldActive != ''" class="bg-primarySoft text-white text-xs ml-1 mr-3 py-1 px-2 rounded">
        {{ searchFieldActive }}
        <span @click="handlerClearSearchField"
          class="cursor-pointer text-xs opacity-50 hover:opacity-100 hover:font-bold ml-1"> x
        </span>
      </div>

      <input type="text" v-model="searchContent"
        placeholder="field:term <e.g. from:john.zufferli@enron.com or name:john>"
        class="flex-1 text-sm bg-transparent text-gray-700 placeholder-gray-400 focus:outline-none"
        @keyup.enter="searchHandler" />

      <button class="ml-3 bg-primaryMiddle text-white text-xs py-1 px-3 rounded hover:bg-primaryDark focus:outline-none"
        @click="searchHandler">
        Buscar
      </button>
    </div>

    <div
      class="w-1/3 flex items-center justify-center ml-3 space-x-4 bg-primaryExtraSoft px-1 py-2 rounded-lg font-light text-xs">
      <span class="font-normal text-sm text-gray-800 leading-none">Search by:</span>

      <div class="flex items-center space-x-2">
        <input type="radio" id="emails" value="emails" v-model="selectedSearchTypeOption" />
        <label for="emails" class="font-normal text-sm text-gray-800cursor-pointer leading-none">
          Emails
        </label>
      </div>

      <div class="flex items-center space-x-2">
        <input type="radio" id="persons" value="persons" v-model="selectedSearchTypeOption" />
        <label for="persons" class="font-normal text-sm text-gray-800 cursor-pointer leading-none">
          Persons
        </label>
      </div>
    </div>
  </div>
</template>

<style>
.highlight {
  background-color: yellow;
  font-weight: bold;
}
</style>
