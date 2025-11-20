<script setup lang="ts">
import ModalWindow from '@/shared/ui-kit/modal-window/modal-window.vue'
import PreviewOfBoard from './preview-of-board/preview-of-board.vue'
import BoardBackground from './board-background/board-background.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultBackground, type DefaultBackground } from '../api/get-default-backgrounds'
const modelValue = defineModel<boolean>()

const selectedId = reactive({
  boardBackgroundId: '',
})

const { execute: getDefaultBackgrounds, data: defaultBackgrounds } = getDefaultBackground()
onMounted(async () => {
  await getDefaultBackgrounds()
})

const handleSelect = (selectedBackground: DefaultBackground) => {
  selectedId.boardBackgroundId = selectedBackground.id
}
</script>

<template>
  <modal-window v-model="modelValue" header="Create New Board" :style="{ width: '448px' }">
    <template #content>
      <preview-of-board />
      <board-background
        v-if="defaultBackgrounds"
        :default-backgrounds
        @select="handleSelect"
        :selectedId="selectedId.boardBackgroundId"
      />
    </template>
  </modal-window>
</template>
