<script setup lang="ts">
import CreateBoard from './createBoard.vue'
</script>

<template>
  <div class="header-for-boards">
    <h1>All Boards</h1>
    <create-board />
  </div>
</template>

<style scoped>
.header-for-boards {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 24px;
}
</style>
