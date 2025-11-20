<script setup lang="ts">
import ModalWindow from '@/shared/ui-kit/modal-window/modal-window.vue'
import PreviewOfBoard from './preview-of-board/preview-of-board.vue'
import BoardBackground from './board-background/board-background.vue'
import { computed, onMounted, reactive, watch } from 'vue'
import { getDefaultBackground, type DefaultBackground } from '../api/get-default-backgrounds'
const modelValue = defineModel<boolean>()

const selectedId = reactive({
  boardBackgroundId: '',
})

const { execute: getDefaultBackgrounds, data: defaultBackgrounds } = getDefaultBackground()
onMounted(async () => {
  await getDefaultBackgrounds()
})

watch(
  defaultBackgrounds,
  (backgrounds) => {
    if (backgrounds && backgrounds.length > 0 && !selectedId.boardBackgroundId) {
      selectedId.boardBackgroundId = backgrounds[0]?.id ?? ''
    }
  },
  { immediate: true },
)

const handleSelect = (selectedBackground: DefaultBackground) => {
  selectedId.boardBackgroundId = selectedBackground.id
}
const previewSrc = computed(() => {
  if (!defaultBackgrounds.value) return ''
  const selectedBackground = defaultBackgrounds.value.find(
    (background) => background.id === selectedId.boardBackgroundId,
  )
  return selectedBackground?.src ?? defaultBackgrounds.value[0]?.src ?? ''
})
</script>

<template>
  <modal-window v-model="modelValue" header="Create New Board" :style="{ width: '448px' }">
    <template #content>
      <preview-of-board :src="previewSrc" />
      <board-background
        v-if="defaultBackgrounds"
        :default-backgrounds
        @select="handleSelect"
        :selectedId="selectedId.boardBackgroundId"
      />
    </template>
  </modal-window>
</template>
