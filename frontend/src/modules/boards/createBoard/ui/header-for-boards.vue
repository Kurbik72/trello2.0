<script setup lang="ts">
import CreateBoard from './create-board.vue'
</script>

<template>
  <div class="header-for-boards">
    <h1 class="header-title">All Boards</h1>
    <create-board />
  </div>
</template>

<style scoped>
.header-for-boards {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'Inter';
}
.header-title {
  color: #1f2937;
  font-size: 24px;
  font-style: normal;
  font-weight: 600;
  line-height: 32px;
}
</style>
