<script setup lang="ts">
interface Props {
  message: string;
}

defineProps<Props>();
</script>

<template>
  <div class="w-full h-full py-12 mb-4 bg-grayExtraSoft">
    <div class="'w-full h-full flex flex-col justify-center items-center  py-12'">
      <div class="w-20 pb-2 opacity-70">
        <img src="../assets/img/error.png" alt="" />
      </div>
      <h4 class="font-roboto text-2xl text-grayDark ">
        <slot name="error-message">
          {{ message }}
        </slot>
      </h4>
      <p class="font-roboto  text-grayDark">
        [ Something went wrong, please try again latter ]
      </p>
    </div>
  </div>
</template>
