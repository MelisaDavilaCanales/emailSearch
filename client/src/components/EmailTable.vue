<script setup lang="ts">

import { useEmailStore } from '@/stores/useEmailStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { onBeforeMount } from 'vue'

const { fetchEmails, fetchEmail, emails } = useEmailStore()
const { setSelectedItemType } = useItemSelectedStore()

const selectedEmail = (emailId: string) => {
  fetchEmail(emailId)
  setSelectedItemType('email')
}

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
              <th class="pl-3 py-2 text-top cursor-pointer">#</th>
              <th class="px-2 py-2 text-left cursor-pointer">Date ↕</th>
              <th class="px-2 py-2 text-left cursor-pointer">From ↕</th>
              <th class="px-2 py-2 text-left cursor-pointer">To ↕</th>
              <th class="px-2 py-2 text-left cursor-pointer">Subject ↕</th>
            </tr>
          </thead>
          <tbody class="text-gray-600">
            <tr @click="selectedEmail(email.id)" v-for="(email, index) in emails" :key="email.id"
              class="border-t hover:bg-gray-50 cursor-pointer">
              <td class="pl-3 py-2 align-top">{{ index + 1 }}</td>
              <td class="px-2 py-2 align-top whitespace-nowrap">{{ email.day }} {{ email.time }}</td>
              <td class="px-2 py-2 align-top">{{ email.from }}</td>
              <td class="px-2 py-2 align-top">
                <span v-if="email?.toArray && email?.toArray.length > 1"
                  class="my-1 block max-h-32 overflow-y-auto custom-scrollbar">
                  <span v-for="(emailAddress, index) in email.toArray" :key="index">
                    {{ emailAddress }}<br />
                  </span>
                </span>
                <span v-else class="text-xs flex pt-1 ml-1">N/A</span>
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
