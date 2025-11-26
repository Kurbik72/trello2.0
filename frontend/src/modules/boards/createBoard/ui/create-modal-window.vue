<script setup lang="ts">
import ModalWindow from '@/shared/ui-kit/modal-window/modal-window.vue'
import PreviewOfBoard from './preview-of-board/preview-of-board.vue'
import BoardBackground from './board-background/board-background.vue'
import inputText from '@/shared/ui-kit/input/input.vue'
import actionButton from '@/shared/ui-kit/button/button.vue'

import { computed, reactive, watch } from 'vue'
import { getDefaultBackground, type DefaultBackground } from '../api/get-default-backgrounds'
const modelValue = defineModel<boolean>()

const form = reactive({
  boardBackgroundId: '',
  boardTitle: '',
})

const { data: defaultBackgroundsData, execute: getDefaultBackgroundExecute } =
  getDefaultBackground()
watch(modelValue, async (newModalValue) => {
  if (newModalValue) {
    await getDefaultBackgroundExecute()
  }
})

watch(
  defaultBackgroundsData,
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
  const selectedBackground = defaultBackgroundsData.value?.find(
    (background) => background.id === form.boardBackgroundId,
  )
  return selectedBackground?.src ?? defaultBackgroundsData.value?.[0]?.src ?? ''
})

const handleCancel = () => {
  modelValue.value = false
}
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
        v-if="defaultBackgroundsData"
        :default-backgrounds="defaultBackgroundsData"
        @select="handleSelect"
        :selectedId="form.boardBackgroundId"
      />
      <div class="action-buttons">
        <action-button severity="secondary" label="Cancel" @click="handleCancel" />
        <action-button theme="purple" label="Create" />
      </div>
    </template>
  </modal-window>
</template>

<style scoped>
.input-text {
  width: 100%;
}
.action-buttons {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}
</style>
