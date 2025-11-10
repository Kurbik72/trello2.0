<script setup lang="ts">
import HeaderForBoards from '@/modules/boards/createBoard/ui/header-for-boards.vue'
import FavoriteList from '@/modules/boards/boardList/ui/favorite-list.vue'
import OtherBoards from '@/modules/boards/boardList/ui/other-boards.vue'
</script>

<template>
  <div class="boards">
    <header-for-boards />
    <favorite-list />
    <other-boards />
  </div>
</template>

<style scoped>
.boards {
  display: flex;
  flex-direction: column;
  gap: 24px;
  padding: 32px 24px;
}
</style>
