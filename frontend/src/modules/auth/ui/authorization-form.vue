<script setup lang="ts">
import { computed } from 'vue'

interface AuthorizationFormProps {
  maxWidth: string
}

const props = defineProps<AuthorizationFormProps>()

const maxWidth = computed(() => `${props.maxWidth}px`)
</script>

<template>
  <form class="authorization-form">
    <div class="header">
      <slot name="header" />
    </div>
    <div class="field-list" v-if="$slots['field-list']">
      <slot name="field-list" />
    </div>
    <div class="submit-content" v-if="$slots['submit-content']">
      <slot name="submit-content" />
    </div>
  </form>
</template>

<style scoped>
.authorization-form {
  font-family: 'Inter';
  max-width: v-bind(maxWidth);
  background-color: #ffff;
  border-radius: 8px;
  box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.1);
  padding: 24px;
}
.header,
.field-list,
.submit-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 12px;
}
</style>
