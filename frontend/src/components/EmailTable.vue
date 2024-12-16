<script setup lang="ts">

import { useEmailStore } from '@/stores/useEmailStore'
import { onBeforeMount } from 'vue'

const { fetchEmails, emails } = useEmailStore()

onBeforeMount(async () => {
    fetchEmails()
});

</script>

<!-- <div class="loading-container text-center">
      <p class="text-lg font-semibold mb-4 text-sky-200">Cargando datos...</p>
    </div> -->

<template>
    <main class="h-full overflow-auto table-container rounded-md">
        <div class="table-wrapper h-full flex items-center">
            <div class="overflow-y-auto max-h-full w-full border rounded-lg">
                <table ref="dataTable" class="min-w-full table-auto border-collapse bg-white text-sm">
                    <thead class="bg-gray-100 sticky top-0 z-10">
                        <tr>
                            <th class="px-4 py-2 text-left ">#</th>
                            <th class="px-4 py-2 text-left cursor-pointer">Id ↕</th>
                            <th class="px-4 py-2 text-left cursor-pointer">From ↕</th>
                            <th class="px-4 py-2 text-left cursor-pointer">To ↕</th>
                            <th class="px-4 py-2 text-left cursor-pointer">Subject ↕</th>
                            <th class="px-4 py-2 text-left cursor-pointer">Date ↕</th>
                        </tr>
                    </thead>
                    <tbody class="text-gray-600">
                        <tr v-for="(email, index) in emails" :key="email.message_id" class="border-t hover:bg-gray-50">
                            <td class="px-4 py-2">{{ index + 1 }}</td>
                            <td class="px-4 py-2">{{ email.from }}</td>
                            <td class="px-4 py-2">{{ email.to }}</td>
                            <td class="px-4 py-2">{{ email.subject }}</td>
                            <td class="px-4 py-2">{{ email.date }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </main>
</template>

<!-- <style scoped>
.table-container {
    max-width: 100%;
    overflow-x: auto;
    white-space: nowrap;
}

.table-wrapper {
    display: inline-block;
    min-width: 100%;
}
</style> -->