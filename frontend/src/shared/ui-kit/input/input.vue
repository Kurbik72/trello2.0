<script setup lang="ts">
import type { Message } from 'primevue'
import InputText from 'primevue/inputtext'
import { useAttrs } from 'vue'

const attrs = useAttrs()

defineProps<{
  label: string
  error?: string
}>()
</script>

<template>
  <div class="input">
    <label for="input" class="input-label">{{ label }}</label>
    <InputText v-bind="attrs" />
    <Message v-if="error" severity="error" size="small" variant="simple">{{ error }}</Message>
  </div>
</template>

<style scoped>
.input {
  display: flex;
  flex-direction: column;
  gap: 6px;
}
.p-inputtext {
  border-color: transparent;
  font-family: 'Inter';
  border: 1px solid #cbd5e1;
  background: #fff;
  font-size: 14px;
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
.input-label {
  font-family: 'Inter';
  color: #334155;
  font-size: 14px;
  font-style: normal;
  font-weight: 500;
}
</style>
