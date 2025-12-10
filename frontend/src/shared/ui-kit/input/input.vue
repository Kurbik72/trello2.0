<script setup lang="ts">
import Message from 'primevue/message'
import InputText from 'primevue/inputtext'
import { useAttrs, computed } from 'vue'

const attrs = useAttrs()

const props = defineProps<{
  label: string
  error?: string
}>()

const inputClasses = computed(() => ({
  'p-error': props.error,
}))
</script>

<template>
  <div class="input">
    <label for="input" class="input-label">{{ label }}</label>
    <InputText v-bind="attrs" :class="inputClasses" />
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
.p-inputtext.p-error {
  border-color: #e24c4c;
}
.p-inputtext.p-error:enabled:hover,
.p-inputtext.p-error:enabled:focus {
  border-color: #e24c4c;
}
.input-label {
  font-family: 'Inter';
  color: #334155;
  font-size: 14px;
  font-style: normal;
  font-weight: 500;
}
</style>

<style>
.p-message-sm .p-message-text {
  font-family: 'Inter';
  font-style: normal;
  font-weight: 400;
  font-size: 12px;
}
</style>
