<script setup lang="ts">
interface BoardCardProps {
  title: string
  linkToBoard: string
  isFavorite: boolean
}

const props = defineProps<BoardCardProps>()
</script>

<template>
  <div class="board-card">
    <h1 class="board-card-title">Board Card</h1>
  </div>
</template>

<style scoped></style>
