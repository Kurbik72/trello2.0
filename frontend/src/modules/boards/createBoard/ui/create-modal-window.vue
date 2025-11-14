<script setup lang="ts">
import ModalWindow from '@/shared/ui-kit/modal-window/modal-window.vue'
import PreviewOfBoard from './preview-of-board/preview-of-board.vue'

const modelValue = defineModel<boolean>()
</script>

<template>
  <modal-window v-model="modelValue" header="Create New Board" :style="{ width: '448px' }">
    <template #content><preview-of-board /></template>
  </modal-window>
</template>
