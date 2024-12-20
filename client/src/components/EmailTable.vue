<script setup lang="ts">

import { useEmailStore } from '@/stores/useEmailStore'
import { onBeforeMount } from 'vue'

const { fetchEmails, emails } = useEmailStore()

onBeforeMount(async () => {
  fetchEmails()
});

</script>

<template>
  <main class="h-full overflow-auto table-container rounded-md">
    <div class="table-wrapper h-full flex items-center">
      <div class="overflow-y-auto max-h-full w-full border rounded-lg custom-scrollbar">
        <table ref="dataTable" class="min-w-full table-auto border-collapse bg-white text-sm">
          <thead class="bg-gray-100 sticky top-0 z-10">
            <tr>
              <th class="px-2 py-2 text-left cursor-pointer">Date ↕</th>
              <th class="px-2 py-2 text-left cursor-pointer">From ↕</th>
              <th class="px-2 py-2 text-left cursor-pointer">To ↕</th>
              <th class="px-2 py-2 text-left cursor-pointer">Subject ↕</th>
            </tr>
          </thead>
          <tbody class="text-gray-600">
            <tr v-for="email in emails" :key="email.id" class="border-t hover:bg-gray-50">
              <td class="px-2 py-2 align-top whitespace-nowrap">{{ email.day }} {{ email.time }}</td>
              <td class="px-2 py-2 align-top">{{ email.from }}</td>
              <td class="px-2 py-2 align-top">
                <span class="my-1 block max-h-32 overflow-y-auto custom-scrollbar">
                  <span v-for="(emailAddress, index) in email.toArray" :key="index">
                    {{ emailAddress }}<br />
                  </span>
                </span>
              </td>
              <td class="px-2 py-2 align-top">{{ email.subject }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 8px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f9fafb;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background-color: #d1d5db;
  border-radius: 10px;
}
</style>
