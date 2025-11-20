<script setup lang="ts">
import { computed } from 'vue'

interface BackgroundItemProps {
  src: string
  selectedId: string
  id: string
}

const props = defineProps<BackgroundItemProps>()

const backgroundItemClasses = computed(() => ({
  'background-item': true,
  'background-item-selected': props.selectedId === props.id,
}))

const emit = defineEmits<{
  (e: 'select'): void
}>()

const handleSelect = () => {
  emit('select')
}
</script>

<template>
  <div :class="backgroundItemClasses" @click="handleSelect">
    <img :src="props.src" class="background-item-image" />
  </div>
</template>

<style scoped>
.background-item {
  width: 91px;
  height: 64px;
  position: relative;
  border-radius: 8px;
  cursor: pointer;
  overflow: hidden;
}
.background-item-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 8px;
}
.background-item-selected {
  border: 2px solid #635bff;
}
.background-item-selected::after {
  content: '';
  position: absolute;
  inset: 0;
  background-image: url('@/assets/icons/done-tick.svg');
  background-position: center;
  background-repeat: no-repeat;
  background-size: 30px 30px;
}
</style>
