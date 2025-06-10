<template>
    <div class="vintage-container">
        <div class="parchment-bg">
            <NuxtLink 
                class="read-manuscript-btn"
                to="/"
            > Inicio
            </NuxtLink>

            <h1 class="vintage-title">Listado de Artículos</h1>
            
            <div v-if="loading" class="loading-vintage">
                Cargando artículos
            </div>
            
            <div v-else>
                <div v-if="articles.length === 0" class="no-articles-vintage">
                    No hay artículos disponibles
                </div>
                
                <div v-else class="articles-collection">
                    <div v-for="(article) in articles" :key="article._id" class="manuscript-card">
                        <div class="manuscript-title">{{ article.title }}</div>
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
const route = useRoute()
const router = useRouter()

// Parámetros de paginación
const currentPage = ref(parseInt(route.query.page) || 1)
const limit = ref(5) // Artículos por página

// Computed para calcular offset
const offset = computed(() => (currentPage.value - 1) * limit.value)

// Fetch con paginación
const { data, pending: loading, refresh } = await useFetch('http://localhost:8080/articles/', {
    query: {
        page: currentPage,
        limit: limit,
        offset: offset
    }
})

const articles = computed(() => data.value?.data ?? [])
const totalArticles = computed(() => data.value?.total ?? 0)
const totalPages = computed(() => Math.ceil(totalArticles.value / limit.value))

// Función para navegar a una página específica
const goToPage = async (page) => {
    if (page < 1) {
        return
    }
    
    currentPage.value = page
    
    // Actualizar URL
    await router.push({
        query: { ...route.query, page: page }
    })
    
    // Refrescar datos
    await refresh()
}

// Watch para cambios en la query de la URL
watch(() => route.query.page, (newPage) => {
    const page = parseInt(newPage) || 1
    if (page !== currentPage.value) {
        currentPage.value = page
    }
})
</script>
