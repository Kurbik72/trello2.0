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
  opacity: 0.5;
  flex: 1 0 calc(33.333% - 12px);
}
.background-item-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.background-item-selected {
  border: 3.5px solid #635bff;
  scale: 1.1;
  opacity: 1;
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
