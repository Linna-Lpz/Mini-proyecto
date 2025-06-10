export default defineNuxtConfig({
  compatibilityDate: '2023-10-01',
  typescript: {
    strict: true
  },
  app: {
    head: {
      title: 'Nuxt3 Test App',
      meta: [
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Template de prueba para postulantes' }
      ]
    }
  },
  css : ['~/assets/css/main.css'],
})