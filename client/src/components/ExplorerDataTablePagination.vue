<script setup lang="ts">

import { defineProps, ref, defineEmits, onMounted } from 'vue';

interface Props {
  currentPage: number;
  totalPages: number;
  pageSize: number;
}
const props = defineProps<Props>();

const currentPageSize = ref(0);
const emit = defineEmits(['pageSizeChange', 'prevPage', 'nextPage']);

const handlerPageSizeChange = (event: Event) => {
  const newSize = (event.target as HTMLSelectElement).value;
  currentPageSize.value = parseInt(newSize)
  emit('pageSizeChange', currentPageSize.value)
};

const sizes = [5, 10, 20, 30, 40, 50];

onMounted(() => {
  currentPageSize.value = props.pageSize;
});

</script>

<template>
  <div class="w-full py-2 flex items-center justify-between px-4 sm:px-6 border-t border-gray-200">
    <div class="flex justify-between w-screen sm:flex sm:flex-1 sm:items-center sm:justify-between">

      <!-- Pagination information -->
      <div>
        <p class="text-sm text-gray-700 sm:flex space-x-1">
          <span>Page</span>
          <span class="font-extrabold text-primary">{{ currentPage }}</span>
          <span>of</span>
          <span class="font-extrabold text-primary">{{ totalPages }}</span>
          <span class="mx-4">|</span>
          <span class="hidden sm:block">Page size:</span>
          <select v-model="currentPageSize" @change="handlerPageSizeChange"
            class="hidden sm:inline-block text-sm text-gray-700 bg-none ml-1 bg-transparent cursor-pointer">
            <option v-for="size in sizes" :key="size" :value="size">{{ size }}</option>
          </select>
        </p>
      </div>

      <!-- Pagination buttons -->
      <div>
        <nav class="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">

          <!-- Previous button -->
          <div :onClick="() => $emit('prevPage')">
            <button
              class="relative inline-flex items-center rounded-l-md px-2 py-1 disabled:opacity-50 disabled:cursor-not-allowed"
              :disabled="currentPage <= 1" @click.prevent="$emit('prevPage')">
              <span class="sr-only">Previous</span>
              <svg class="size-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd"
                  d="M11.78 5.22a.75.75 0 0 1 0 1.06L8.06 10l3.72 3.72a.75.75 0 1 1-1.06 1.06l-4.25-4.25a.75.75 0 0 1 0-1.06l4.25-4.25a.75.75 0 0 1 1.06 0Z"
                  clip-rule="evenodd" />
              </svg>
            </button>
          </div>

          <!-- Current page -->
          <span aria-current="page"
            class="relative z-10 inline-flex items-center justify-center bg-primary px-4 py-1 text-sm font-semibold text-white focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primaryVariant text-center self-center w-14">
            {{ currentPage }}
          </span>

          <!-- Next button -->
          <div :onClick="() => $emit('nextPage')">
            <button
              class="relative inline-flex items-center rounded-r-md px-2 py-1 disabled:opacity-50 disabled:cursor-not-allowed"
              :disabled="currentPage == totalPages" @click.prevent="$emit('nextPage')">
              <span class="sr-only">Next</span>
              <svg class="size-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd"
                  d="M8.22 5.22a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L11.94 10 8.22 6.28a.75.75 0 0 1 0-1.06Z"
                  clip-rule="evenodd" />
              </svg>
            </button>
          </div>
        </nav>
      </div>
    </div>
  </div>
</template>
