<template>
  <div class="login-container">
    <NuxtLink 
        class="read-manuscript-btn"
        to="/"
    > Volver al inicio
    </NuxtLink>
    <h2>Iniciar sesión</h2>
    <form @submit.prevent="onLogin">
      <div>
        <label for="email">Email:</label>
        <input v-model="email" type="email" id="email" required />
      </div>
      <div>
        <label for="password">Contraseña:</label>
        <input v-model="password" type="password" id="password" required />
      </div>
      <button type="submit">Entrar</button>
      <p v-if="error" class="error-msg">{{ error }}</p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '../composables/useAuth'

const email = ref('')
const password = ref('')
const error = ref('')
const { login } = useAuth()

const onLogin = async () => {
  error.value = ''
  try {
    await login(email.value, password.value)
    // Guarda el token en: localStorage.setItem('token', res.token)
    window.location.href = '/'
  } catch (e: any) {
    error.value = e.message || 'Error al iniciar sesión'
  }
}
</script>