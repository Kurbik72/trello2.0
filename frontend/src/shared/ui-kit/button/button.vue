<script setup lang="ts">
import Button from 'primevue/button'
import { computed, useAttrs } from 'vue'

interface ButtonProps {
  theme?: 'purple'
}

const props = defineProps<ButtonProps>()
const attrs = useAttrs()

const buttonClasses = computed(() => ({
  'button-purple': props.theme === 'purple',
}))
</script>

<template>
  <Button v-bind="attrs" :class="buttonClasses" />
</template>

<style scoped>
.p-button {
  border: none;
  font-family: 'Inter';
  font-size: 16px;
  padding: 8px 16px;
  font-weight: 500;
  border-radius: 5px;
}
.p-button:hover {
  border: none !important;
}
.p-button.button-purple {
  background: #635bff;
  color: #fff;
  border: none;
  font-family: 'Inter';
}
.p-button.button-purple:hover {
  background: #453ddd;
}
</style>
