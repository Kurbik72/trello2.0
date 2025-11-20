import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import 'primeicons/primeicons.css'

import Material from '@primeuix/themes/material'
import { VueQueryPlugin } from '@tanstack/vue-query'
import App from './App.vue'
import router from './app/router'
import './assets/fonts/fonts.css'
import Tooltip from 'primevue/tooltip'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(PrimeVue, {
  theme: {
    preset: Material,
    options: {
      darkModeSelector: false,
    },
  },
})
app.use(VueQueryPlugin)
app.directive('tooltip', Tooltip)

app.mount('#app')
