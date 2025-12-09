<script setup lang="ts">
import BoardCard from './board-card/board-card.vue'
import { useBoardsStore } from '@/shared/stores/boards'
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const boardsStore = useBoardsStore()

onMounted(async () => {
  await boardsStore.getBoardsList((route.query.user_id as string) ?? '')
})
const toggleFavorite = async (boardId: string) => {
  const toggleFavoriteStar = boardsStore.boards.find((board) => board.id === boardId)
  if (toggleFavoriteStar) {
    toggleFavoriteStar.isFavorite = !toggleFavoriteStar.isFavorite
  }
}
</script>

<template>
  <div class="other-boards">
    <h1 class="other-boards-title"><i class="pi pi-align-justify" />Other Boards</h1>
    <div class="board-cards">
      <board-card
        v-for="board in boardsStore.boards"
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
.other-boards {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.other-boards-title {
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
