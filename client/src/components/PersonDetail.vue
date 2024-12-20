<!-- <script setup lang="ts">

import { useItemSelectedStore } from '@/stores/useItemSelectedStore'

const { isItemSelected } = useItemSelectedStore()

</script> -->

<template>
    <div class="h-full w-full bg-grayExtraSoft rounded-md space-y-4 px-4 pt-4">
        <!-- person's data -->
        <div class="flex space-x-2">
            <div class="w-14  flex items-center">
                <img src="../assets/img/person.png" alt="" class="">
            </div>
            <div class="flex flex-col justify-center w-full">
                <h2 class="text-2xl font-bold">Maria Valentina</h2>
                <p class="text-sm">maria.valentina@gmail.com</p>
                <hr class="border-t border-grayDark mt-2">
            </div>
        </div>
        <!-- button's section -->
        <div class="flex justify-end space-x-2 pr-4">
            <button class="btn-secondary">
                Enviados
            </button>
            <button class="btn-secondary">
                Recibidos
            </button>
        </div>
        <!-- emails's person -->
        <div class="h-[65%] py-2 overflow-y-auto space-y-2">
            <div class="bg-graySoft rounded-md p-2 flex space-x-2">
                <div class="w-12">
                    <img src="../assets/img/email.png" alt="">
                </div>
                <div class="text-sm">
                    <p class=""> <span class="font-bold">Message ID:</span> 8201835.1075849629259 </p>
                    <p class=""> <span class="font-bold">To: </span> Jorge Melendez </p>
                    <p class=""> <span class="font-bold">Subject:</span> Solicitud para el registro de una propiedad en
                        SPS</p>
                </div>
            </div>
            <div class="bg-graySoft rounded-md p-2 flex space-x-2">
                <div class="w-12">
                    <img src="../assets/img/email.png" alt="">
                </div>
                <div class="text-sm">
                    <p class=""> <span class="font-bold">Message ID:</span> 8201835.1075849629259 </p>
                    <p class=""> <span class="font-bold">To: </span> Jorge Melendez </p>
                    <p class=""> <span class="font-bold">Subject:</span> Solicitud para el registro de una propiedad en
                        SPS</p>
                </div>
            </div>
            <div class="bg-graySoft rounded-md p-2 flex space-x-2">
                <div class="w-12">
                    <img src="../assets/img/email.png" alt="">
                </div>
                <div class="text-sm">
                    <p class=""> <span class="font-bold">Message ID:</span> 8201835.1075849629259 </p>
                    <p class=""> <span class="font-bold">To: </span> Jorge Melendez </p>
                    <p class=""> <span class="font-bold">Subject:</span> Solicitud para el registro de una propiedad en
                        SPS</p>
                </div>
            </div>
            <div class="bg-graySoft rounded-md p-2 flex space-x-2">
                <div class="w-12">
                    <img src="../assets/img/email.png" alt="">
                </div>
                <div class="text-sm">
                    <p class=""> <span class="font-bold">Message ID:</span> 8201835.1075849629259 </p>
                    <p class=""> <span class="font-bold">To: </span> Jorge Melendez </p>
                    <p class=""> <span class="font-bold">Subject:</span> Solicitud para el registro de una propiedad en
                        SPS</p>
                </div>
            </div>
            <div class="bg-graySoft rounded-md p-2 flex space-x-2">
                <div class="w-12">
                    <img src="../assets/img/email.png" alt="">
                </div>
                <div class="text-sm">
                    <p class=""> <span class="font-bold">Message ID:</span> 8201835.1075849629259 </p>
                    <p class=""> <span class="font-bold">To: </span> Jorge Melendez </p>
                    <p class=""> <span class="font-bold">Subject:</span> Solicitud para el registro de una propiedad en
                        SPS</p>
                </div>
            </div>
            <div class="bg-graySoft rounded-md p-2 flex space-x-2">
                <div class="w-12">
                    <img src="../assets/img/email.png" alt="">
                </div>
                <div class="text-sm">
                    <p class=""> <span class="font-bold">Message ID:</span> 8201835.1075849629259 </p>
                    <p class=""> <span class="font-bold">To: </span> Jorge Melendez </p>
                    <p class=""> <span class="font-bold">Subject:</span> Solicitud para el registro de una propiedad en
                        SPS</p>
                </div>
            </div>
            <div class="bg-graySoft rounded-md p-2 flex space-x-2">
                <div class="w-12">
                    <img src="../assets/img/email.png" alt="">
                </div>
                <div class="text-sm">
                    <p class=""> <span class="font-bold">Message ID:</span> 8201835.1075849629259 </p>
                    <p class=""> <span class="font-bold">To: </span> Jorge Melendez </p>
                    <p class=""> <span class="font-bold">Subject:</span> Solicitud para el registro de una propiedad en
                        SPS</p>
                </div>
            </div>
            <div class="bg-graySoft rounded-md p-2 flex space-x-2">
                <div class="w-12">
                    <img src="../assets/img/email.png" alt="">
                </div>
                <div class="text-sm">
                    <p class=""> <span class="font-bold">Message ID:</span> 8201835.1075849629259 </p>
                    <p clas=""> <span class="font-bold">To: </span> Jorge Melendez </p>
                    <p class=""> <span class="font-bold">Subject:</span> Solicitud para el registro de una propiedad en
                        SPS</p>
                </div>
            </div>
        </div>
        <!-- data pagination -->
        <div class="pb-4 md:pb-0">
            <div id="pagination" class="flex items-center justify-between px-4 sm:px-6">
                <div class="flex justify-between w-screen  sm:flex sm:flex-1 sm:items-center sm:justify-between">
                    <div>
                        <p class="text-sm text-gray-700 sm:flex space-x-1">
                            <span>Page</span>
                            <!-- <span class="font-extrabold text-primary">{{ "personStore.state.currentPage" }}</span> -->
                            <span class="font-extrabold text-primary">{{ 1 }}</span>
                            <span>of</span>
                            <!-- <span class="font-extrabold text-primary">{{ "personStore.state.totalPages" }}</span> -->
                            <span class="font-extrabold text-primary">{{ 2 }}</span>
                            <span class="hidden sm:block"> | Page size:</span>
                            <!-- <span class="hidden sm:block font-medium">{{ "personStore.state.pageSize" }}</span> -->
                            <span class="hidden sm:block font-medium">{{ 10 }}</span>
                        </p>
                    </div>
                    <div>
                        <nav class="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">
                            <a href="#" class="relative inline-flex items-center rounded-l-md px-2 py-1">
                                <span class="sr-only">Previous</span>
                                <svg class="size-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true"
                                    data-slot="icon">
                                    <path fill-rule="evenodd"
                                        d="M11.78 5.22a.75.75 0 0 1 0 1.06L8.06 10l3.72 3.72a.75.75 0 1 1-1.06 1.06l-4.25-4.25a.75.75 0 0 1 0-1.06l4.25-4.25a.75.75 0 0 1 1.06 0Z"
                                        clip-rule="evenodd" />
                                </svg>
                            </a>

                            <a href="#" aria-current="page"
                                class="relative z-10 inline-flex items-center justify-center bg-primary px-4 py-1 text-sm font-semibold text-white focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primaryVariant text-center self-center w-14">
                                <!-- {{ " personStore.state.currentPage" }} -->
                                {{ 1 }}
                            </a>

                            <a href="#" class="relative inline-flex items-center rounded-r-md px-2 py-1">
                                <span class="sr-only">Next</span>
                                <svg class="size-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true"
                                    data-slot="icon">
                                    <path fill-rule="evenodd"
                                        d="M8.22 5.22a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L11.94 10 8.22 6.28a.75.75 0 0 1 0-1.06Z"
                                        clip-rule="evenodd" />
                                </svg>
                            </a>
                        </nav>
                    </div>

                </div>
            </div>
        </div>
    </div>
</template>