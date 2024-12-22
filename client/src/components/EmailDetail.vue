<script setup lang="ts">

import { useEmailTableStore } from '@/stores/useEmailTableStore'
import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { storeToRefs } from 'pinia'

const emailStore = useEmailTableStore()

const { emailDetails } = storeToRefs(emailStore)

const { setSelectedItemType } = useItemSelectedStore()
const { setEmailSearchTerm, setEmailSearchField, } = useEmailViewerStore()


const showPersonDetail = (personEmail: string) => {
  setEmailSearchTerm(personEmail)
  setEmailSearchField('from')

  alert('show person detail: ' + personEmail)

  setSelectedItemType('person')
}

</script>

<template>
  <div class="font-roboto space-y-3 overflow-x-hidden h-full flex flex-col pb-2">
    <!-- Contenedor general con altura máxima y desplazamiento habilitado -->
    <div class="flex flex-col space-y-2 flex-1 overflow-y-auto">

      <!-- Datos principales del correo -->
      <div class="w-full overflow-x-hidden flex-none">
        <div>
          <div class="flex justify-between">
            <div>
              <p><span class="font-bold">Message ID:</span> {{ emailDetails?.message_id }}</p>
              <p><span class="font-bold">Date:</span> {{ emailDetails?.date }}</p>
              <p class="text-sm">
                <span class="font-bold">From:</span>
                <span
                  class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
                  @click="emailDetails?.from && showPersonDetail(emailDetails.from)"> {{ emailDetails?.from }}
                </span>
              </p>
            </div>
            <div class="-mt-1 mr-2 w-16">
              <img src="../assets/img/email-png.png" alt="Email Icon">
            </div>
          </div>

          <p class="text-sm mb-1" v-if="emailDetails?.subject && emailDetails?.subject.trim() !== ''">
            <span class="font-bold">Subject:</span> {{ emailDetails?.subject }}
          </p>
        </div>
      </div>

      <!-- Contenido del correo -->
      <div class="flex-1 overflow-y-auto m-1 custom-scrollbar">
        <!-- Datos de destinatarios -->
        <div class="text-sm flex">
          <span class="font-bold flex mr-1">To:</span>
          <div class="max-h-36 overflow-x-hidden overflow-y-auto custom-scrollbar"
            v-if="emailDetails?.toArray && emailDetails?.toArray.length > 0 && emailDetails?.toArray[0] !== ''">
            <span
              class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
              v-for="(emailAddress, index) in emailDetails?.toArray" :key="index"
              @click="emailAddress && showPersonDetail(emailAddress)">
              {{ emailAddress }}
              <span v-if="index < emailDetails?.toArray.length - 1">, </span>
            </span>
          </div>
          <span v-else class="text-sm flex pt-1 ml-1">N/A</span>
        </div>

        <!-- Datos de copia -->
        <div class="text-sm flex"
          v-if="emailDetails?.ccArray && emailDetails?.ccArray.length > 0 && emailDetails?.ccArray[0] !== ''">
          <span class="font-bold flex mr-1">Cc:</span>
          <div class="max-h-36 overflow-x-hidden overflow-y-auto custom-scrollbar">
            <span
              class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
              v-for="(emailAddress, index) in emailDetails?.ccArray" :key="index"
              @click="showPersonDetail(emailAddress)">
              {{ emailAddress }}
              <span v-if="index < emailDetails?.ccArray.length - 1">, </span>
            </span>
          </div>
        </div>

        <hr class="border-t mt-2 border-grayDark">

        <!-- Metadata -->
        <div class="mt-2">
          <p class="text-sm" v-if="emailDetails?.x_folder && emailDetails?.x_folder.trim() !== ''">
            <span class="font-bold">Folder:</span> {{ emailDetails?.x_folder }}
          </p>
          <p class="text-sm" v-if="emailDetails?.x_origin && emailDetails?.x_origin.trim() !== ''">
            <span class="font-bold">Origin:</span> {{ emailDetails?.x_origin }}
          </p>
          <p class="text-sm" v-if="emailDetails?.x_file_name && emailDetails?.x_file_name.trim() !== ''">
            <span class="font-bold">File Name:</span> {{ emailDetails?.x_file_name }}
          </p>
        </div>

        <!-- Contenido -->
        <div>
          <p class="font-bold sticky top-0 bg-grayExtraSoft z-10">Contenido:</p>
          <p class="text-sm py-1">{{ emailDetails?.content || 'No content available' }}</p>
        </div>
      </div>
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
