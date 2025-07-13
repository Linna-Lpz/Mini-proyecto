<template>  
  <div class="vintage-container">
    <div class="parchment-bg-home">
      <h1 class="vintage-title">Bienvenido al Archivo de Artículos</h1>
      <p class="intro-vintage">
        Explora nuestra colección de artículos y descubre un mundo interesante
        adentrándote en temas variados. Desde historia hasta ciencia.
      </p>
      <nuxt-link to="/articlesList" class="read-manuscript-btn">
        Comenzar a leer
      </nuxt-link>

      <p class="intro-vintage">
        Se parte de nuestra comunidad escritora. Regístrate o inicia sesión para publicar tus
        propios artículos y compartir tus conocimientos con el mundo.
      </p>
      <nuxt-link to="/register" class="read-manuscript-btn">
        Registrarse
      </nuxt-link>
      
      <nuxt-link to="/login" class="read-manuscript-btn">
        Iniciar sesión
      </nuxt-link>
    </div>
  </div>
</template>

<script>
</script>