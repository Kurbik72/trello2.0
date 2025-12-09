<script setup lang="ts">
import BoardCard from './board-card/board-card.vue'
import { useBoardsStore } from '@/shared/stores/boards'
import { computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'

interface Props {
  title: string
  icon: string
  iconColor?: string
  showOnlyFavorites?: boolean
  loadBoards?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  iconColor: '',
  showOnlyFavorites: false,
  loadBoards: false,
})

const route = useRoute()
const boardsStore = useBoardsStore()

onMounted(async () => {
  if (props.loadBoards) {
    await boardsStore.getBoardsList((route.query.user_id as string) ?? '')
  }
})

const toggleFavorite = async (boardId: string) => {
  await boardsStore.favoriteBoardList(boardId, route.query.user_id as string)
}

const displayBoards = computed(() => {
  if (props.showOnlyFavorites) {
    return boardsStore.boards.filter((board) => board.isFavorite)
  }
  return boardsStore.boards
})
</script>

<template>
  <div class="board-list">
    <h1 class="board-list-title"><i :class="icon" :style="{ color: iconColor }" />{{ title }}</h1>
    <div class="board-cards">
      <board-card
        v-for="board in displayBoards"
        :key="board.id"
        :id="board.id"
        :title="board.title"
        :isFavorite="board.isFavorite"
        :linkToBoard="board.linkToBoard"
        :backgroundSrc="board.backgroundSrc"
        @toggleFavorite="() => toggleFavorite(board.id)"
      />
    </div>
  </div>
</template>

<style scoped>
.board-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.board-list-title {
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
