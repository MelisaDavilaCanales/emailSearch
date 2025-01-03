<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { storeToRefs } from 'pinia';

import { useSearchTypeStore } from '@/stores/useSearchTypeStore';

const { setSearchFieldActive } = useSearchTypeStore();
const { searchFieldActive } = storeToRefs(useSearchTypeStore());

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

const selectedOption = ref<string>('');

onMounted(() => {
  // Set the initial value of the select based on the value of searchFieldActive
  if (selectedOption.value === '') {
    selectedOption.value = searchFieldActive.value;
  }
});

function handleSearchFieldChange(event: Event) {
  const target = event.target as HTMLSelectElement;
  const value = target.value;
  selectedOption.value = value;
  setSearchFieldActive(value);
}
</script>

<template>
  <div class="p-1 w-full flex justify-center font-roboto">
    <div class="flex space-x-2 items-center">
      <p class="text-sm text-gray-800">Search for matches only in the field:</p>
      <select
        class="text-sm border border-primarySoft p-1 rounded focus:outline-none focus:ring-1 focus:ring-primarySoft"
        v-model="selectedOption" @change="handleSearchFieldChange">
        <option value="" disabled>Select field</option>
        <option v-for="option in searchEmailOptions" :key="option.value" :value="option.value">
          {{ option.label }}
        </option>
      </select>
    </div>
  </div>
</template>
