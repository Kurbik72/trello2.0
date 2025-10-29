<script setup lang="ts">
import { iconIds } from '@/shared/constants/IconIds'
import { defineAsyncComponent, shallowRef, watchEffect } from 'vue'

interface IconProps {
  width: number
  height: number
  iconId: iconIds
}

const props = defineProps<IconProps>()

const iconPath = shallowRef()

watchEffect(() => {
  iconPath.value = defineAsyncComponent(() => import(`@/assets/icons/${props.iconId}.svg`))
})
</script>

<template>
  <component :is="iconPath" :width="props.width" :height="props.height" />
</template>
