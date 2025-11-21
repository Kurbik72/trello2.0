<script setup lang="ts">
import buttonForOpenModalWindow from './button-for-open-modal-window.vue'
import { useToggle } from '@/shared/composables/use-toggle'
import CreateModalWindow from './create-modal-window.vue'
const { isActive, setActive } = useToggle()
</script>

<template>
  <div class="header-for-boards">
    <h1 class="header-title">All Boards</h1>
    <button-for-open-modal-window @toggle="setActive" />
    <create-modal-window v-model="isActive" />
  </div>
</template>

<style scoped>
.header-for-boards {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'Inter';
}
.header-title {
  color: #1f2937;
  font-size: 24px;
  font-style: normal;
  font-weight: 600;
  line-height: 32px;
}
</style>
