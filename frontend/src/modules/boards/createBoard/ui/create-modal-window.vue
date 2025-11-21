<script setup lang="ts">
import ModalWindow from '@/shared/ui-kit/modal-window/modal-window.vue'
import PreviewOfBoard from './preview-of-board/preview-of-board.vue'
import BoardBackground from './board-background/board-background.vue'
import inputText from '@/shared/ui-kit/input/input.vue'
import { computed, onMounted, reactive, watch } from 'vue'
import { getDefaultBackground, type DefaultBackground } from '../api/get-default-backgrounds'
const modelValue = defineModel<boolean>()

const form = reactive({
  boardBackgroundId: '',
  boardTitle: '',
})

const { execute: getDefaultBackgrounds, data: defaultBackgrounds } = getDefaultBackground()
onMounted(async () => {
  await getDefaultBackgrounds()
})

watch(
  defaultBackgrounds,
  (backgrounds) => {
    if (backgrounds && backgrounds.length > 0 && !form.boardBackgroundId) {
      form.boardBackgroundId = backgrounds[0]?.id ?? ''
    }
  },
  { immediate: true },
)

const handleSelect = (selectedBackground: DefaultBackground) => {
  form.boardBackgroundId = selectedBackground.id
}
const previewSrc = computed(() => {
  if (!defaultBackgrounds.value) return ''
  const selectedBackground = defaultBackgrounds.value.find(
    (background) => background.id === form.boardBackgroundId,
  )
  return selectedBackground?.src ?? defaultBackgrounds.value[0]?.src ?? ''
})
</script>

<template>
  <modal-window v-model="modelValue" header="Create New Board" :style="{ width: '448px' }">
    <template #content>
      <preview-of-board :src="previewSrc" />
      <input-text
        type="text"
        placeholder="e.g., Q1 Marketing Plan"
        label="Board title"
        class="input-text"
        v-model="form.boardTitle"
      />
      <board-background
        v-if="defaultBackgrounds"
        :default-backgrounds
        @select="handleSelect"
        :selectedId="form.boardBackgroundId"
      />
    </template>
  </modal-window>
</template>

<style scoped>
.input-text {
  width: 100%;
}
</style>
