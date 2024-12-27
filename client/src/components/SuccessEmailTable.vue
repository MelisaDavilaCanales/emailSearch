<script setup lang="ts">
import { useEmailTableStore } from '@/stores/useEmailTableStore';
import { useSearchTypeStore } from '@/stores/useSearchTypeStore';
import { onBeforeMount, computed } from 'vue';
import { storeToRefs } from 'pinia';

const emailStore = useEmailTableStore();
const { emailList } = storeToRefs(emailStore);
const { fetchEmails } = useEmailTableStore();

const searchTypeStore = useSearchTypeStore();
const { searchTerm } = storeToRefs(searchTypeStore);

onBeforeMount(() => {
  fetchEmails();
});

const highlightText = (text: string, term: string): string => {
  if (!term) return text;
  const regex = new RegExp(`(${term})`, 'gi');
  return text.replace(regex, '<span class="highlight">$1</span>');
};

const highlightedEmails = computed(() => {
  const term = searchTerm.value || '';
  return emailList.value.map(email => {
    return {
      ...email,
      from: highlightText(email.from, term),
      subject: highlightText(email.subject, term),
      toArray: email.toArray?.map(to => highlightText(to, term)) || [],
    };
  });
});
</script>

<template>
  <main class="h-full overflow-auto table-container rounded-md">
    <div class="table-wrapper h-full flex items-center">
      <div class="overflow-y-auto overflow-x-hidden max-h-full w-full border rounded-lg custom-scrollbar">
        <table class="min-w-full table-auto border-collapse bg-white text-sm">
          <thead class="bg-gray-100 sticky top-0 z-10">
            <tr>
              <th class="pl-3 py-2">#</th>
              <th class="px-2 py-2">From</th>
              <th class="px-2 py-2">To</th>
              <th class="px-2 py-2">Subject</th>
            </tr>
          </thead>
          <tbody class="text-gray-600">
            <tr v-for="(email, index) in highlightedEmails" :key="index" class="border-t">
              <td class="pl-3 py-2">{{ index + 1 }}</td>
              <td class="px-2 py-2" v-html="email.from"></td>
              <td class="px-2 py-2">
                <span v-for="to in email.toArray" :key="to" v-html="to" class="block"></span>
              </td>
              <td class="px-2 py-2" v-html="email.subject"></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<style scoped>
.highlight {
  background-color: yellow;
  font-weight: bold;
}

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
