<template>
  <div class="register-container">
    <NuxtLink 
        class="read-manuscript-btn"
        to="/"
    > Volver al inicio
    </NuxtLink>
    <h2>Registrarse</h2>
    <form @submit.prevent="onRegister">
      <div>
        <label for="name">Nombre:</label>
        <input v-model="name" type="text" id="name" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <input v-model="email" type="email" id="email" required />
      </div>
      <div>
        <label for="password">Contraseña:</label>
        <input v-model="password" type="password" id="password" required />
      </div>
      <button type="submit">Registrarse</button>
      <p v-if="error" class="error-msg">{{ error }}</p>
      <p v-if="success" class="success-msg">{{ success }}</p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '../composables/useAuth'

const name = ref('')
const email = ref('')
const password = ref('')
const error = ref('')
const success = ref('')
const { register } = useAuth()

const onRegister = async () => {
  error.value = ''
  success.value = ''
  try {
    await register(name.value, email.value, password.value)
    success.value = 'Usuario registrado correctamente. Ahora puedes iniciar sesión.'
    name.value = ''
    email.value = ''
    password.value = ''
  } catch (e: any) {
    error.value = e.message || 'Error al registrar usuario'
  }
}
</script>