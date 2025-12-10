<script setup lang="ts">
import ModalWindow from '@/shared/ui-kit/modal-window/modal-window.vue'
import PreviewOfBoard from './preview-of-board/preview-of-board.vue'
import BoardBackground from './board-background/board-background.vue'
import inputText from '@/shared/ui-kit/input/input.vue'
import actionButton from '@/shared/ui-kit/button/button.vue'
import { computed, reactive, watch } from 'vue'
import { getDefaultBackground, type DefaultBackground } from '../api/get-default-backgrounds'
import { useBoardsStore } from '@/shared/stores/boards'
import { useRoute } from 'vue-router'
import * as v from 'valibot'
import { useFormValidation } from '@/shared/composables/use-form-validation'

const modelValue = defineModel<boolean>()

const form = reactive({
  boardBackgroundId: '',
  boardTitle: '',
})

const FormSchema = v.object({
  boardTitle: v.pipe(
    v.string('Введите строку'),
    v.nonEmpty('Поле обязательно для заполнения'),
    v.minLength(3, 'Минимум 3 символа'),
  ),
  boardBackgroundId: v.pipe(v.string('Введите строку'), v.nonEmpty('Выберите фон')),
})

const { errors, validate, clearErrors } = useFormValidation(FormSchema)

const { data: defaultBackgroundsData, execute: getDefaultBackgroundExecute } =
  getDefaultBackground()
watch(modelValue, () => {
  if (modelValue.value) {
    getDefaultBackgroundExecute()
    clearErrors()
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

const route = useRoute()
const boardStore = useBoardsStore()

const handleCreate = async () => {
  if (!validate(form)) {
    return
  }

  await boardStore.saveBoard({
    title: form.boardTitle,
    backgroundId: form.boardBackgroundId,
    user_id: route.query.user_id as string,
  })

  modelValue.value = false
  form.boardTitle = ''
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
        :error="errors.boardTitle"
      />
      <board-background
        v-if="defaultBackgroundsData"
        :default-backgrounds="defaultBackgroundsData"
        @select="handleSelect"
        :selectedId="form.boardBackgroundId"
      />
      <div class="action-buttons">
        <action-button severity="secondary" label="Cancel" @click="handleCancel" />
        <action-button theme="purple" label="Create" @click="handleCreate" />
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
