<template>
    <div class="vintage-container">
        <div class="parchment-bg">
            <NuxtLink 
                class="read-manuscript-btn"
                to="/"
            > Inicio
            </NuxtLink>

            <h1 class="vintage-title">Listado de Artículos</h1>

            <!-- Filtros de búsqueda y fecha -->
            <div class="filter-section">
                <input 
                    type="text" 
                    v-model="searchQuery" 
                    placeholder="Buscar artículos..." 
                    class="search-input"
                    @keyup.enter="searchArticles"
                />
                <input 
                    type="date" 
                    v-model="fromDate" 
                    class="date-input" 
                />
                <input 
                    type="date" 
                    v-model="toDate" 
                    class="date-input" 
                />
                <button class="search-btn" @click="searchArticles">
                    Aplicar filtro
                </button>
                <button class="search-btn" @click="cleanFilters">
                    Limpiar filtro
                </button>
            </div>
            
            <div v-if="loading" class="loading-vintage">
                Cargando artículos
            </div>
            
            <div v-else>
                <div v-if="articles.length === 0" class="no-articles-vintage">
                    No hay artículos disponibles
                </div>
                
                <!-- Articulos -->
                <div v-else class="articles-collection">
                    <div v-for="(article) in articles" :key="article.id" class="manuscript-card">
                        <div class="manuscript-title">{{ article.title }}</div>
                        <div class="date-vintage">{{ article.created_at.slice(0, 10) }}</div>
                        <div class="manuscript-content">{{ article.content }}</div>
                        <NuxtLink
                            class="read-manuscript-btn"
                            :to="{name: 'articleForm', query: { id: article.id }}"
                        > Leer artículo completo
                        </NuxtLink>
                    </div>
                </div>

                <!-- Paginación -->
                <div class="pagination-vintage">
                    <button 
                        class="pagination-btn"
                        :disabled="currentPage <= 1"
                        @click="goToPage(currentPage - 1)"
                    >
                        <
                    </button>
                    
                    <span class="pagination-info">
                        Página {{ currentPage }} de {{ totalPages }}
                    </span>
                    
                    <button 
                        class="pagination-btn"
                        :disabled="currentPage >= totalPages"
                        @click="goToPage(currentPage + 1)"
                    >
                        >
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useRoute, useRouter } from '#app'
const route = useRoute()
const router = useRouter()

// Parámetros reactivos sincronizados con la URL
const currentPage = ref(parseInt(route.query.page) || 1)
const limit = ref(5)
const searchQuery = ref(route.query.title || '')
const fromDate = ref(route.query.from || '')
const toDate = ref(route.query.to || '')

const offset = computed(() => (currentPage.value - 1) * limit.value)

// Estado para los artículos y la carga
const articles = ref([])
const totalArticles = ref(0)
const totalPages = computed(() => Math.ceil(totalArticles.value / limit.value))
const loading = ref(false)

// Función para obtener artículos según los filtros actuales
const fetchArticles = async () => {
  loading.value = true
  try {
    const params = new URLSearchParams({
      page: currentPage.value,
      limit: limit.value,
      offset: offset.value,
      title: searchQuery.value,
      from: fromDate.value,
      to: toDate.value
    })
    const res = await fetch(`http://localhost:8080/articles/articles?${params}`)
    const json = await res.json()
    articles.value = json.data || []
    totalArticles.value = json.total || 0
  } catch (e) {
    articles.value = []
    totalArticles.value = 0
  } finally {
    loading.value = false
  }
}

// Sincronizar filtros con la URL y recargar artículos
const goToPage = async (page) => {
  if (page < 1) return
  currentPage.value = page
  await router.push({
    query: { ...route.query, page }
  })
}

const searchArticles = async () => {
  await router.push({
    query: {
      ...route.query,
      title: searchQuery.value,
      from: fromDate.value,
      to: toDate.value,
      page: 1
    }
  })
}

const cleanFilters = async () => {
  searchQuery.value = ''
  fromDate.value = ''
  toDate.value = ''
  currentPage.value = 1
  await router.push({
    query: { page: 1 }
  })
}

// Watch para recargar artículos cuando cambian los filtros en la URL
watch(
  () => [route.query.page, route.query.title, route.query.from, route.query.to],
  ([newPage, newTitle, newFrom, newTo]) => {
    currentPage.value = parseInt(newPage) || 1
    searchQuery.value = newTitle || ''
    fromDate.value = newFrom || ''
    toDate.value = newTo || ''
    fetchArticles()
  },
  { immediate: true }
)
</script>
