<script setup lang="ts">
import previewColumns from './preview-columns.vue'

interface PreviewOfBoardProps {}
const props = defineProps<PreviewOfBoardProps>()
</script>

<template>
  <div class="preview-of-board">
    <preview-columns />
  </div>
</template>

<style scoped>
.preview-of-board {
}
</style>
