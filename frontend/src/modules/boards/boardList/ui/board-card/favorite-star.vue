<script setup lang="ts">
import { computed } from 'vue'

interface FavoriteStarProps {
  isFavorite: boolean
}

const props = defineProps<FavoriteStarProps>()

const emit = defineEmits<{
  (e: 'toggle'): void
}>()

const handleToggle = () => {
  emit('toggle')
}

const iconClasses = computed(() => ({
  'pi pi-star-fill': true,
  'star-active': props.isFavorite,
  'star-inactive': !props.isFavorite,
}))
</script>

<template>
  <div class="favorite-star" @click="handleToggle">
    <i :class="iconClasses" />
  </div>
</template>

<style scoped>
.favorite-star {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.favorite-star {
  background: #f4f5f7;
  border-radius: 6px;
}
.star-active {
  color: #eab308;
  -webkit-text-stroke: 0;
}

.star-inactive {
  color: transparent;
  -webkit-text-stroke: 1.5px #eab308;
}
</style>
