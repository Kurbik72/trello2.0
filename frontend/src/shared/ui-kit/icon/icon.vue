<script setup lang="ts">
import { iconIds } from '@/shared/constants/iconIds'
import { defineAsyncComponent, shallowRef, toRefs, watchEffect } from 'vue'

interface IconProps {
  width: number
  height: number
  iconId: iconIds
}

const props = defineProps<IconProps>()

const { width, height } = toRefs(props)

const iconPath = shallowRef()

watchEffect(() => {
  iconPath.value = defineAsyncComponent(() => import(`@/assets/icons/${props.iconId}.svg`))
})
</script>

<template>
  <component :is="iconPath" class="icon" />
</template>

<style scoped>
.icon {
  width: v-bind('width + "px"');
  height: v-bind('height + "px"');
  overflow: hidden;
}
</style>
