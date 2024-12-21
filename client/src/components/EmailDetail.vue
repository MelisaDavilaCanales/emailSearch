<script setup lang="ts">

import { useEmailStore } from '@/stores/useEmailStore'
import { storeToRefs } from 'pinia'

const emailStore = useEmailStore()

const { emailDetail } = storeToRefs(emailStore)


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
            <p> <span class="font-bold"> Message ID: </span>{{ emailDetail?.message_id }}</p>
            <p> <span class="font-bold"> Date: </span>{{ emailDetail?.date }} </p>
            <p class="text-sm"> <span class="font-bold"> From: </span>{{ emailDetail?.from }} </p>
          </div>
          <div class="-mt-1 mr-2 w-16">
            <img src="../assets/img/email-png.png" alt="Email Icon">
          </div>
        </div>

        <p class="text-sm mb-1" v-if="emailDetail?.subject && emailDetail?.subject.trim() !== ''">
          <span class="font-bold"> Subject: </span> {{ emailDetail?.subject }}
        </p>

        <div class="text-sm flex ">
          <span class="font-bold flex pt-1 mr-1"> To: </span>
          <div class="max-h-36 overflow-x-visible overflow-y-auto custom-scrollbar"
            v-if="emailDetail?.toArray && emailDetail?.toArray.length > 1">
            <span
              class="cursor-pointer px-1 py-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
              v-for="(emailAddress, index) in emailDetail?.toArray" :key="index"
              @click="handleEmailClick(emailAddress)">
              {{ emailAddress }}
              <span v-if="index < emailDetail?.toArray.length - 1">, </span>
            </span>
          </div>
          <span v-else class="text-sm flex pt-1 ml-1">N/A</span>
        </div>
      </div>

    </div>

    <hr class="border-t border-grayDark">

    <div class="ml-2">

      <p class="text-sm" v-if="emailDetail?.cc && emailDetail?.cc.trim() !== ''">
        <span class="font-bold"> Cc: </span>{{ emailDetail?.cc }}
      </p>

      <p class="text-sm" v-if="emailDetail?.bcc && emailDetail?.bcc.trim() !== ''">
        <span class="font-bold"> Bcc: </span>{{ emailDetail?.bcc }}
      </p>

      <p class="text-sm" v-if="emailDetail?.x_folder && emailDetail?.x_folder.trim() !== ''">
        <span class="font-bold"> Folder: </span>{{ emailDetail?.x_folder }}
      </p>

      <p class="text-sm" v-if="emailDetail?.x_origin && emailDetail?.x_origin.trim() !== ''">
        <span class="font-bold"> Origin: </span>{{ emailDetail?.x_origin }}
      </p>

      <p class="text-sm" v-if="emailDetail?.x_file_name && emailDetail?.x_file_name.trim() !== ''">
        <span class="font-bold"> File Name: </span>{{ emailDetail?.x_file_name }}
      </p>

      <!-- <p class="text-sm" v-if="emailDetail?.content_type && emailDetail?.content_type.trim() !== ''">
        <span class="font-bold"> Content Type: </span>{{ emailDetail?.content_type }}
      </p>

      <p class="text-sm"
        v-if="emailDetail?.content_transfer_encoding && emailDetail?.content_transfer_encoding.trim() !== ''">
        <span class="font-bold"> Content Transfer Encoding: </span>{{ emailDetail?.content_transfer_encoding }}
      </p> -->

    </div>

    <div class="flex-1 overflow-y-auto m-1 custom-scrollbar">
      <p class="font-bold sticky top-0 bg-grayExtraSoft ">Contenido:</p>
      <p class="text-sm p-1"> {{ emailDetail?.content || 'No content available' }} </p>
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
