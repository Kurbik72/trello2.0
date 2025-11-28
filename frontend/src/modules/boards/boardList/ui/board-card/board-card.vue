<script setup lang="ts">
import { computed } from 'vue'
import FavoriteStar from './favorite-star.vue'

interface BoardCardProps {
  id: string
  title: string
  isFavorite: boolean
  linkToBoard: string
  backgroundSrc: string
}

const props = defineProps<BoardCardProps>()

const backgroundImage = computed(() => `url(${props.backgroundSrc})`)

const emit = defineEmits<{
  (e: 'toggleFavorite'): void
}>()

const handleToggleFavorite = () => {
  emit('toggleFavorite')
}
</script>

<template>
  <div class="board-card">
    <favorite-star
      :isFavorite="props.isFavorite"
      @toggle="handleToggleFavorite"
      class="favorite-star"
    />
    <h1 class="board-card-title">{{ props.title }}</h1>
  </div>
</template>

<style scoped>
.board-card {
  background-image: v-bind(backgroundImage);
  position: relative;
  display: flex;
  width: 350px;
  height: 160px;
  padding: 16px;
  overflow: hidden;
}
.favorite-star {
  transform: translateX(200%);
  transition: transform 0.3s ease;
  position: absolute;
  top: 16px;
  right: 16px;
}
.board-card:hover .favorite-star {
  transform: translateX(0);
}
@keyframes star-animation {
  0% {
    transform: translateX(100%);
  }
  100% {
    transform: translateX(0);
  }
}
.board-card-title {
  font-family: 'Inter';
  font-size: 18px;
  font-weight: 500;
  line-height: 24px;
  color: white;
}
</style>
