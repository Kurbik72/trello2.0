<script setup lang="ts">
import Dialog from 'primevue/dialog'
import { useAttrs } from 'vue'

const modelValue = defineModel<boolean>()

const attrs = useAttrs()
</script>

<template>
  <Dialog v-model:visible="modelValue" modal :dismissableMask="true" v-bind="attrs">
    <slot name="content" />
  </Dialog>
</template>

<style>
.p-dialog-mask {
  background: rgba(0, 0, 0, 0.5) !important;
  backdrop-filter: blur(2px) !important;
}

.p-dialog-header {
  font-family: 'Inter' !important;
  font-size: 24px;
  font-style: normal;
  font-weight: 600;
  line-height: 32px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
</style>
