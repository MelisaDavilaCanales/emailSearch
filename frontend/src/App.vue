<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="w-full h-screen flex flex-col">
    <header class="sticky top-0 z-50 bg-white shadow-md">
      <nav class="w-full py-2 bg-primaryDark text-white flex justify-center">
        <RouterLink to="/">EmailSearch</RouterLink>
      </nav>
    </header>
    <RouterView class="w-full flex-grow" />
  </div>
</template>
