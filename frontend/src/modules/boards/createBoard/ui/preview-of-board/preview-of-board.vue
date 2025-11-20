<script setup lang="ts">
import previewColumns from './preview-columns.vue'
import { getDefaultBackground } from '../../api/get-default-backgrounds'
import { onMounted } from 'vue'

const {} = getDefaultBackground()

onMounted(async () => {
  await execute()
})
</script>

<template>
  <div class="preview-of-board">
    <div v-if="isLoading">Loading...</div>
    <div v-else-if="error">Error: {{ error.message }}</div>
    <div v-else>
      <preview-columns />
      <div v-if="defaultBackgrounds">
        <div v-for="bg in defaultBackgrounds" :key="bg.src">
          <img :src="bg.src" alt="Background" />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.preview-of-board {
}
</style>
