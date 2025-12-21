<script setup lang="ts">
import { customIcon, primeVueIcons } from '@/shared/constants/iconIds'
import { iconMap } from '@/shared/utils/icon-map'
import { computed, defineAsyncComponent, useAttrs } from 'vue'

type CustomIcon = keyof typeof customIcon

type PrimeVueIcon = keyof typeof primeVueIcons

interface CommonProps<T> {
  width: number
  height: number
  name: T
}

export type CustomIconProps = CommonProps<CustomIcon> & {
  isCustomIcon: true
  isPrimeVueIcon?: false
}

export type PrimeVueIconProps = CommonProps<PrimeVueIcon> & {
  isPrimeVueIcon: true
  isCustomIcon?: false
}
type IconProps = CustomIconProps | PrimeVueIconProps

const props = defineProps<IconProps>()

const sizeOfIcon = computed(() => ({
  width: `${props.width}px`,
  height: `${props.height}px`,
}))
const icon = computed(() =>
  props.isCustomIcon
    ? defineAsyncComponent(() => import(`@/assets/icons/${iconMap(props)}.svg`))
    : 'i',
)
const attrs = useAttrs()
const styleIcon = computed(() => ['icon', iconMap(props)])
</script>

<template>
  <component :is="icon" :class="styleIcon" v-bind="attrs" />
</template>

<style scoped>
.icon {
  font-size: v-bind('sizeOfIcon.width');
  display: flex;
  align-items: center;
  justify-content: center;
  width: v-bind('sizeOfIcon.width');
  height: v-bind('sizeOfIcon.height');
  overflow: hidden;
}
</style>
