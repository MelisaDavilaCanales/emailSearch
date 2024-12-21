<script setup lang="ts">

import { useEmailStore } from '@/stores/useEmailStore'
import { storeToRefs } from 'pinia'

const emailStore = useEmailStore()

const { emailDetails } = storeToRefs(emailStore)


const handleEmailClick = (emailAddress: string) => {
  console.log('Email clicked:', emailAddress)
}

</script>

<template>
  <div class="font-roboto space-y-3 h-full flex flex-col pb-2">
    <!-- emailDeailMainHeader -->
    <div class="flex space-x-4">

      <div class="w-full">
        <div class="flex justify-between">
          <div class="">
            <p> <span class="font-bold"> Message ID: </span>{{ emailDetails?.message_id }}</p>
            <p> <span class="font-bold"> Date: </span>{{ emailDetails?.date }} </p>
            <p class="text-sm"> <span class="font-bold"> From: </span>{{ emailDetails?.from }} </p>
          </div>
          <div class="-mt-1 mr-2 w-16">
            <img src="../assets/img/email-png.png" alt="Email Icon">
          </div>
        </div>

        <p class="text-sm mb-1" v-if="emailDetails?.subject && emailDetails?.subject.trim() !== ''">
          <span class="font-bold"> Subject: </span> {{ emailDetails?.subject }}
        </p>

        <div class="text-sm flex ">
          <span class="font-bold flex pt-1 mr-1"> To: </span>
          <div class="max-h-36 overflow-x-visible overflow-y-auto custom-scrollbar"
            v-if="emailDetails?.toArray && emailDetails?.toArray.length > 1">
            <span
              class="cursor-pointer px-1 py-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
              v-for="(emailAddress, index) in emailDetails?.toArray" :key="index"
              @click="handleEmailClick(emailAddress)">
              {{ emailAddress }}
              <span v-if="index < emailDetails?.toArray.length - 1">, </span>
            </span>
          </div>
          <span v-else class="text-sm flex pt-1 ml-1">N/A</span>
        </div>
      </div>

    </div>

    <hr class="border-t border-grayDark">

    <div class="ml-2">

      <p class="text-sm" v-if="emailDetails?.cc && emailDetails?.cc.trim() !== ''">
        <span class="font-bold"> Cc: </span>{{ emailDetails?.cc }}
      </p>

      <p class="text-sm" v-if="emailDetails?.bcc && emailDetails?.bcc.trim() !== ''">
        <span class="font-bold"> Bcc: </span>{{ emailDetails?.bcc }}
      </p>

      <p class="text-sm" v-if="emailDetails?.x_folder && emailDetails?.x_folder.trim() !== ''">
        <span class="font-bold"> Folder: </span>{{ emailDetails?.x_folder }}
      </p>

      <p class="text-sm" v-if="emailDetails?.x_origin && emailDetails?.x_origin.trim() !== ''">
        <span class="font-bold"> Origin: </span>{{ emailDetails?.x_origin }}
      </p>

      <p class="text-sm" v-if="emailDetails?.x_file_name && emailDetails?.x_file_name.trim() !== ''">
        <span class="font-bold"> File Name: </span>{{ emailDetails?.x_file_name }}
      </p>

      <!-- <p class="text-sm" v-if="emailDetails?.content_type && emailDetails?.content_type.trim() !== ''">
        <span class="font-bold"> Content Type: </span>{{ emailDetails?.content_type }}
      </p>

      <p class="text-sm"
        v-if="emailDetails?.content_transfer_encoding && emailDetails?.content_transfer_encoding.trim() !== ''">
        <span class="font-bold"> Content Transfer Encoding: </span>{{ emailDetails?.content_transfer_encoding }}
      </p> -->

    </div>

    <div class="flex-1 overflow-y-auto m-1 custom-scrollbar">
      <p class="font-bold sticky top-0 bg-grayExtraSoft ">Contenido:</p>
      <p class="text-sm p-1"> {{ emailDetails?.content || 'No content available' }} </p>
    </div>
  </div>
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
