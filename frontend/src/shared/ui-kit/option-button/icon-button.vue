<script setup lang="ts">
import { computed } from 'vue'
import { useToggle } from '@/shared/composables/use-toggle'

interface IconButtonProps {
  icon: string
  isActive: boolean
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
  <div :class="iconButtonClasses" @click="handleToggle">
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
