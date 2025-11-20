<script setup lang="ts">
import BackgroundItem from './background-item.vue'

interface DefaultBackground {
  id: string
  src: string
}

interface BoardBackgroundProps {
  defaultBackgrounds: DefaultBackground[]
  selectedId: string
}

const props = defineProps<BoardBackgroundProps>()

const emit = defineEmits<{
  select: [background: DefaultBackground]
}>()

const handleSelect = (selectedBackground: DefaultBackground) => {
  emit('select', selectedBackground)
}
</script>

<template>
  <div class="board-background">
    <span class="board-background-title">Background</span>
    <div class="board-background-items">
      <background-item
        class="board-background-item"
        v-for="background in defaultBackgrounds"
        :key="background.id"
        :src="background.src"
        :selected-id
        :id="background.id"
        @select="handleSelect(background)"
      />
    </div>
  </div>
</template>

<style scoped>
.board-background {
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 12px;
}
.board-background-title {
  color: #334155;
  font-family: 'Inter';
  font-size: 14px;
  font-style: normal;
  font-weight: 500;
  line-height: 20px;
}
.board-background-items {
  display: flex;
  gap: 12px;
}
</style>
