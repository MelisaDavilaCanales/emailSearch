<script setup lang="ts">

import { usePersonStore } from '@/stores/usePersonStore'
import { onBeforeMount } from 'vue'

const { fetchPersons, persons } = usePersonStore()

onBeforeMount(async () => {
  fetchPersons()
});
</script>

<template>
  <main class="h-full overflow-auto table-container rounded-md">
    <div class="table-wrapper h-full flex items-center">
      <div class="overflow-y-auto max-h-full w-full border rounded-lg custom-scrollbar">
        <table ref="dataTable" class="min-w-full table-auto border-collapse bg-white text-sm">
          <thead class="bg-gray-100 sticky top-0 z-10">
            <tr>
              <th class="px-2 py-2 text-center cursor-pointer">#</th>
              <th class="pl-2 py-2 text-left cursor-pointer"></th>
              <th class="px-2 py-2 text-left cursor-pointer">Email ↕</th>
              <th class="px-2 py-2 text-left cursor-pointer">Person ↕</th>
            </tr>
          </thead>
          <tbody class="text-gray-600">
            <tr v-for="(person, index) in persons" :key="person.id" class="border-t hover:bg-gray-50 cursor-pointer">
              <td class="px-2 py-2 align-center">{{ index }}</td>
              <td class="pl-2 py-2align-top">
                <img class="w-12" src="../assets/img/person.png" alt="">
              </td>
              <td class="px-2 py-2 align-top">{{ person.email }}</td>
              <td class="px-2 py-2 align-top ">{{ person.name }} </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>
