<template>
  <div>
    <h2>Selecciona una opción:</h2>
    <div class="option">
      <input type="radio" id="opcion1" value="opcion1" v-model="selectedOption" @change="handleSelection" />
      <label for="opcion1">Opción 1</label>
    </div>
    <div class="option">
      <input type="radio" id="opcion2" value="opcion2" v-model="selectedOption" @change="handleSelection" />
      <label for="opcion2">Opción 2</label>
    </div>

    <p>Opción seleccionada: {{ selectedOption }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const selectedOption = ref<string | null>(null);

const handleSelection = () => {
  console.log('Opción seleccionada:', selectedOption.value);
};
</script>

<style scoped>
.option {
  margin-bottom: 8px;
}
</style>
