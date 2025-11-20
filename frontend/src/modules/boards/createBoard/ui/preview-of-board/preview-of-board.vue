<script setup lang="ts">
import previewColumns from './preview-columns.vue'

import { computed } from 'vue'

interface PreviewOfBoardProps {
  src: string
}

const props = defineProps<PreviewOfBoardProps>()

const previewImage = computed(() => `url(${props.src})`)
</script>

<template>
  <div class="preview-of-board">
    <preview-columns />
  </div>
</template>

<style scoped>
.preview-of-board {
  width: 400px;
  height: 225px;
  padding: 16px;
  background-image: v-bind(previewImage);
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
}
</style>
