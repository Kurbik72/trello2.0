<script setup lang="ts">
import Logo from '@/shared/ui-kit/logo/logo.vue'

const columns = [
  {
    width: 115,
    height: 120,
    lines: 3,
  },
  {
    width: 115,
    height: 80,
    lines: 2,
  },
  {
    width: 115,
    height: 150,
    lines: 5,
  },
]
</script>

<template>
  <div class="preview-columns">
    <logo :with-title="false" />
    <div class="preview-columns-content">
      <div
        v-for="column of columns"
        :key="column.height"
        class="preview-columns-content-item"
        :style="{ width: column.width + 'px', height: column.height + 'px' }"
      >
        <div v-for="line in column.lines" :key="line" class="preview-columns-content-item-line" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.preview-columns-content {
  display: flex;
  gap: 16px;
  justify-content: center;
  margin-top: 24px;
}
.preview-columns-content-item {
  background: #d2d3d4;
  border-radius: 8px;
  padding: 16px;
}
.preview-columns-content-item-line {
  background: #fff;
  height: 10px;
  width: 100%;
  margin-bottom: 10px;
  border-radius: 2px;
}
</style>
