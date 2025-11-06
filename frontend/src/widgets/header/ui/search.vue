<script setup lang="ts">
import InputText from 'primevue/inputtext'
import IconField from 'primevue/iconfield'
import InputIcon from 'primevue/inputicon'
</script>

<template>
  <IconField class="search">
    <InputIcon class="pi pi-search" />
    <InputText type="text" placeholder="Search" />
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
  height: 40px;
}
.p-inputtext:enabled:hover {
  border-color: rgb(145, 145, 145);
}
.p-inputtext:enabled:focus {
  border-color: rgb(145, 145, 145);
  outline: none;
}
</style>
