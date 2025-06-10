<template>  
  <div class="vintage-container">
    <div class="parchment-bg">
      <h1 class="vintage-title">Bienvenido al Archivo de Artículos</h1>
      <p class="intro-vintage">
        Explora nuestra colección de artículos y descubre un mundo interesante.
      </p>
      <nuxt-link to="/articlesList" class="read-manuscript-btn">
        Ir a la lista de artículos
      </nuxt-link>
    </div>
  </div>
</template>

<script>
</script>

<style scoped>
.vintage-container {
  min-height: 89vh;
  padding: 40px 20px;
  background: linear-gradient(135deg, #f4f1e8 0%, #e8dcc0 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.parchment-bg {
  max-width: 600px;
  margin: 0 auto;
  background: #faf8f3;
  border: 2px solid #8b764c;
  border-radius: 12px;
  padding: 50px 40px;
  box-shadow: 0 8px 25px rgba(139, 118, 76, 0.2);
  text-align: center;
}

.vintage-title {
  font-size: 38px;
  color: #4a3c1d;
  margin-bottom: 24px;
  letter-spacing: 1px;
}

.intro-vintage {
  font-size: 20px;
  color: #6b5b3a;
  margin-bottom: 36px;
  line-height: 1.5;
}

.read-manuscript-btn {
  display: inline-block;
  background: #8b764c;
  color: #faf8f3;
  padding: 12px 28px;
  border-radius: 6px;
  text-decoration: none;
  font-weight: 600;
  font-size: 18px;
  box-shadow: 0 2px 8px rgba(139, 118, 76, 0.2);
  transition: background 0.2s;
}

.read-manuscript-btn:hover {
  background: #a48d5f;
  text-decoration: none;
}
</style>