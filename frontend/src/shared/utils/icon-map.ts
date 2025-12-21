import { customIcon, primeVueIcons } from '@/shared/constants/iconIds'
import type { CustomIconProps, PrimeVueIconProps } from '@/shared/ui-kit/icon/icon.vue'

export const iconMap = (icon: CustomIconProps | PrimeVueIconProps): string => {
  return icon.isCustomIcon ? customIcon[icon.name] : primeVueIcons[icon.name]
}
