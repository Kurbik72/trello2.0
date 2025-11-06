<script setup lang="ts">
import InputText from 'primevue/inputtext'
import IconField from 'primevue/iconfield'
import InputIcon from 'primevue/inputicon'
import { useAttrs } from 'vue'

const attrs = useAttrs()
</script>

<template>
  <IconField class="search">
    <InputIcon class="pi pi-search" />
    <InputText type="text" v-bind="attrs" />
  </IconField>
</template>

<style scoped>
.p-inputtext {
  border-color: transparent;
  background: #f4f5f7;
  font-family: 'Inter';
  font-size: 14px;
  border-radius: 8px;
  padding: 8px 8px 8px;
}
.p-inputtext:enabled:hover {
  border-color: rgb(145, 145, 145);
}
.p-inputtext:enabled:focus {
  border-color: rgb(145, 145, 145);
  outline: none;
}
</style>
