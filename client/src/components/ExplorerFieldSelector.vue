<script setup lang="ts">

import { useSearchTypeStore } from '@/stores/useSearchTypeStore';
import { storeToRefs } from 'pinia';
import { ref, computed } from 'vue';

const searchTypeStore = useSearchTypeStore();
const { setSearchFieldActive } = useSearchTypeStore();

const { searchType } = storeToRefs(searchTypeStore);

const searchEmailOptions = [
  { value: '_all', label: '[all fields]' },
  { value: 'message_id', label: 'Message ID' },
  { value: 'date', label: 'Date' },
  { value: 'from', label: 'From' },
  { value: 'to', label: 'To' },
  { value: 'subject', label: 'Subject' },
  { value: 'cc', label: 'CC' },
  { value: 'bcc', label: 'BCC' },
  { value: 'x_folder', label: 'Folder' },
  { value: 'x_origin', label: 'Origin' },
  { value: 'x_file_name', label: 'File name' },
  { value: 'content', label: 'Content' },
];

const searchPersonaOptions = [
  { value: '_all', label: '[all fields]' },
  { value: 'name', label: 'Name' },
  { value: 'email', label: 'Email' },
];


const options = computed(() => {
  if (searchType.value === 'emails') {
    return searchEmailOptions;
  } else if (searchType.value === 'persons') {
    return searchPersonaOptions;
  }
  return [];
});

const selectedOption = ref<string>('');

function handleSelectChange(event: Event) {
  const target = event.target as HTMLSelectElement;
  const value = target.value;
  selectedOption.value = value;
  if (searchType.value === 'emails') {
    setSearchFieldActive(value);
  } else if (searchType.value === 'persons') {
    setSearchFieldActive(value);
  }
}
</script>

<template>
  <div class="p-1 w-full flex justify-center font-roboto">
    <div class="flex space-x-2 items-center">
      <p class="text-sm text-gray-800">Search for matches only in the field:</p>
      <select
        class="text-sm border  border-primarySoft p-1 rounded focus:outline-none focus:ring-1 focus:ring-primarySoft"
        @change="handleSelectChange">
        <option value="" disabled selected>Select field</option>
        <option v-for="option in options" :key="option.value" :value="option.value">
          {{ option.label }}
        </option>
      </select>
    </div>
  </div>
</template>
