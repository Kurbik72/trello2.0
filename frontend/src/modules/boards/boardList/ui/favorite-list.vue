<script setup lang="ts">
import BoardCard from './board-card/board-card.vue'
import { useBoardsStore } from '@/shared/stores/boards'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const boardsStore = useBoardsStore()

const favoriteBoards = computed(() => {
  return boardsStore.boards.filter((board) => board.isFavorite)
})

const toggleFavorite = async (boardId: string) => {
  const toggleFavoriteStar = boardsStore.boards.find((board) => board.id === boardId)
  if (toggleFavoriteStar) {
    toggleFavoriteStar.isFavorite = !toggleFavoriteStar.isFavorite
  }
  await boardsStore.favoriteBoardList(boardId, route.query.user_id as string)
}
</script>

<template>
  <div class="favorite-list">
    <h1 class="favorite-list-title">
      <i class="pi pi-star-fill" :style="{ color: '#EAB308' }" />Favorite List
    </h1>
    <div class="board-cards">
      <board-card
        v-for="board in favoriteBoards"
        :key="board.id"
        :id="board.id"
        :title="board.title"
        :isFavorite="board.isFavorite"
        :linkToBoard="board.linkToBoard"
        :backgroundSrc="board.backgroundSrc"
        @toggleFavorite="toggleFavorite(board.id)"
      />
    </div>
  </div>
</template>

<style scoped>
.favorite-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.favorite-list-title {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #374151;
  font-family: 'Inter';
  font-size: 20px;
  font-style: normal;
  font-weight: 500;
  line-height: 28px;
}
.board-cards {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
}
</style>
