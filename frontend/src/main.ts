import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import 'primeicons/primeicons.css'

import Material from '@primeuix/themes/material'

import App from './App.vue'
import router from './app/router'
import './assets/fonts/fonts.css'

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

app.mount('#app')
