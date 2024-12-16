<script setup lang="ts">
</script>

<template>
  <div class="w-full flex flex-col justify-center items-center md:flex-row md:justify-around">
    <input class="w-[90%] border-2 px-2 md:w-4/6" type="text" placeholder="Buscar">
    <div class="space-x-2">
      <label for="">Search by:</label>
      <input class="searchType" type="radio"> <label for="">Emails</label>
      <input class="searchType" type="radio"> <label for="">Persons</label>
    </div>
  </div>
</template>
