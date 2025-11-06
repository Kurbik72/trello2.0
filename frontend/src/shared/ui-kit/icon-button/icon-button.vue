<script setup lang="ts">
import { computed } from 'vue'

interface IconButtonProps {
  icon: string
  isActive: boolean
  tooltip?: string
}

const props = defineProps<IconButtonProps>()

const emit = defineEmits<{
  (e: 'toggle'): void
}>()

const handleToggle = () => {
  emit('toggle')
}
const iconButtonClasses = computed(() => ({
  'icon-button': true,
  'icon-active': props.isActive,
}))
</script>

<template>
  <div :class="iconButtonClasses" @click="handleToggle" v-tooltip.bottom="props.tooltip">
    <i :class="props.icon" :style="{ color: '#6B7785' }"></i>
  </div>
</template>

<style scoped>
.icon-button {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border-radius: 8px;
  background: #f4f5f7;
  cursor: pointer;
  transition: background 0.5s ease;
}
</style>

<style>
.p-tooltip {
  font-family: 'Inter';
  font-size: 14px;
  font-style: normal;
  font-weight: 200;
}
</style>
