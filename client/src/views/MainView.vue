<script setup lang="ts">
import { storeToRefs } from 'pinia'

import ExplorerContainer from '@/components/ExplorerContainer.vue'
import ViewerContainer from '@/components/ViewerContainer.vue'
import BannerServerError from '@/components/BannerServerError.vue'

import { useEmailTableStore } from '@/stores/useEmailTableStore'
import { useEmailViewerStore } from '@/stores/useEmailViewerStore'
import { usePersonStore } from '@/stores/usePersonStore'

const { serverError: emailTableServerError } = storeToRefs(useEmailTableStore())
const { serverError: emailViewerServerError } = storeToRefs(useEmailViewerStore())
const { serverError: personStoreServerError } = storeToRefs(usePersonStore())

</script>

<template>

  <BannerServerError v-if="emailTableServerError.status" :message="emailTableServerError.message" />
  <BannerServerError v-if="emailViewerServerError.status" :message="emailViewerServerError.message" />
  <BannerServerError v-if="personStoreServerError.status" :message="personStoreServerError.message" />

  <main v-if="!emailTableServerError.status && !emailViewerServerError.status && !personStoreServerError.status"
    class="lg:h-full lg:overflow-y-hidden w-full flex flex-col lg:flex-row">
    <div class="lg:w-7/12">
      <ExplorerContainer />
    </div>
    <div class="lg:w-5/12">
      <ViewerContainer />
    </div>
  </main>
</template>
