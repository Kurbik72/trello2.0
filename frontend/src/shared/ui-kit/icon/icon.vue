<script setup lang="ts">
import { iconIds } from '@/shared/constants/iconIds'
import { computed, defineAsyncComponent, useAttrs } from 'vue'

interface IconProps {
  width: number
  height: number
  iconId?: iconIds
  primeVueIcon?: string
}

const props = defineProps<IconProps>()

const sizeOfIcon = computed(() => ({
  width: `${props.width}px`,
  height: `${props.height}px`,
}))
const icon = computed(() =>
  props.iconId ? defineAsyncComponent(() => import(`@/assets/icons/${props.iconId}.svg`)) : 'i',
)
const attrs = useAttrs()
// TODO: add map for primevue icons for view like this: <i name="user" />
const styleIcon = computed(() => ['icon', `${props.primeVueIcon}`])
</script>

<template>
  <component :is="icon" :class="styleIcon" :attrs="attrs" />
</template>

<style scoped>
.icon {
  width: v-bind('sizeOfIcon.width');
  height: v-bind('sizeOfIcon.height');
  overflow: hidden;
}
</style>
