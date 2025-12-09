<script setup lang="ts">
import HeaderForBoards from '@/modules/boards/createBoard/ui/header-for-boards.vue'
import BoardList from '@/modules/boards/boardList/ui/board-list.vue'
</script>

<template>
  <div class="boards">
    <header-for-boards />
    <board-list
      title="Favorite List"
      icon="pi pi-star-fill"
      icon-color="#EAB308"
      :show-only-favorites="true"
    />
    <board-list title="All Boards" icon="pi pi-align-justify" :load-boards="true" />
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
