<script setup lang="ts">

import { storeToRefs } from 'pinia'
import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { useItemSelectedStore } from '@/stores/useItemSelectedStore'
import { usePersonStore } from '@/stores/usePersonStore'
import { useToast } from "vue-toastification";

const emailStore = useEmailViewerStore()

const { emailDetail } = storeToRefs(emailStore)

const { setSelectedItemType } = useItemSelectedStore()
const { setSelectedPersonEmail } = usePersonStore()
const { setEmailSearchTerm, setEmailSearchField, } = useEmailViewerStore()


const showPersonDetail = (personEmail: string) => {
  setEmailSearchTerm(personEmail)
  setEmailSearchField('from')

  //AGREGAR TOAST
  const toast = useToast()
  toast.success("Person selected")

  setSelectedPersonEmail(personEmail)

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
              <p><span class="font-bold">Message ID:</span> {{ emailDetail?.message_id }}</p>
              <p><span class="font-bold">Date:</span> {{ emailDetail?.date }}</p>
              <p class="text-sm">
                <span class="font-bold">From:</span>
                <span
                  class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
                  @click="emailDetail?.from && showPersonDetail(emailDetail.from)"> {{ emailDetail?.from }}
                </span>
              </p>
            </div>
            <div class="-mt-1 mr-2 w-16">
              <img src="../assets/img/email-png.png" alt="Email Icon">
            </div>
          </div>

          <p class="text-sm mb-1" v-if="emailDetail?.subject && emailDetail?.subject.trim() !== ''">
            <span class="font-bold">Subject:</span> {{ emailDetail?.subject }}
          </p>
        </div>
      </div>

      <!-- Contenido del correo -->
      <div class="flex-1 overflow-y-auto m-1 custom-scrollbar">
        <!-- Datos de destinatarios -->
        <div class="text-sm flex">
          <span class="font-bold flex mr-1">To:</span>
          <div class="max-h-36 overflow-x-hidden overflow-y-auto custom-scrollbar"
            v-if="emailDetail?.toArray && emailDetail?.toArray.length > 0 && emailDetail?.toArray[0] !== ''">
            <span
              class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
              v-for="(emailAddress, index) in emailDetail?.toArray" :key="index"
              @click="emailAddress && showPersonDetail(emailAddress)">
              {{ emailAddress }}
              <span v-if="index < emailDetail?.toArray.length - 1">, </span>
            </span>
          </div>
          <span v-else class="text-sm flex pt-1 ml-1">N/A</span>
        </div>

        <!-- Datos de copia -->
        <div class="text-sm flex"
          v-if="emailDetail?.ccArray && emailDetail?.ccArray.length > 0 && emailDetail?.ccArray[0] !== ''">
          <span class="font-bold flex mr-1">Cc:</span>
          <div class="max-h-36 overflow-x-hidden overflow-y-auto custom-scrollbar">
            <span
              class="cursor-pointer px-1 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-600 hover:bg-gray-200 inline-block whitespace-nowrap"
              v-for="(emailAddress, index) in emailDetail?.ccArray" :key="index"
              @click="showPersonDetail(emailAddress)">
              {{ emailAddress }}
              <span v-if="index < emailDetail?.ccArray.length - 1">, </span>
            </span>
          </div>
        </div>

        <hr class="border-t mt-2 border-grayDark">

        <!-- Metadata -->
        <div class="mt-2">
          <p class="text-sm" v-if="emailDetail?.x_folder && emailDetail?.x_folder.trim() !== ''">
            <span class="font-bold">Folder:</span> {{ emailDetail?.x_folder }}
          </p>
          <p class="text-sm" v-if="emailDetail?.x_origin && emailDetail?.x_origin.trim() !== ''">
            <span class="font-bold">Origin:</span> {{ emailDetail?.x_origin }}
          </p>
          <p class="text-sm" v-if="emailDetail?.x_file_name && emailDetail?.x_file_name.trim() !== ''">
            <span class="font-bold">File Name:</span> {{ emailDetail?.x_file_name }}
          </p>
        </div>

        <!-- Contenido -->
        <div>
          <p class="font-bold sticky top-0 bg-grayExtraSoft z-10">Contenido:</p>
          <p class="text-sm py-1">{{ emailDetail?.content || 'No content available' }}</p>
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
