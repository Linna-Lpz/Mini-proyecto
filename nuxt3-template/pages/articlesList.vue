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

<style scoped>

.vintage-container {
    min-height:100%;
    padding: 40px 20px;
    background: linear-gradient(135deg, #f4f1e8 0%, #e8dcc0 100%);
}

.parchment-bg {
    max-width: 900px;
    margin: 0 auto;
    background: #faf8f3;
    border: 2px solid #8b764c;
    border-radius: 12px;
    padding: 40px;
    box-shadow: 0 8px 25px rgba(139, 118, 76, 0.2);
}

.vintage-title {
    font-size: 40px;
    color: #4a3c1d;
    text-align: center;
    margin-bottom: 40px;
}

.loading-vintage, .no-articles-vintage {
    text-align: center;
    padding: 40px 20px;
    color: #6b5b3a;
}

.articles-collection {
    display: flex;
    flex-direction: column;
    gap: 30px;
}

.manuscript-card {
    background: #fdfcf7;
    border: 1px solid #d4c4a0;
    border-radius: 8px;
    padding: 30px;
    box-shadow: 0 4px 15px rgba(139, 118, 76, 0.1);
}

.manuscript-title {
    font-size: 24px;
    font-weight: 600;
    color: #5d4e2a;
    margin-bottom: 15px;
    letter-spacing: 0.5px;
}

.manuscript-content {
    background: rgba(255, 255, 255, 0.6);
    border: 1px solid rgba(139, 118, 76, 0.2);
    border-radius: 6px;
    padding: 20px;
    color: #4a3c1d;
    font-size: 18px;
    line-height: 20px;
    margin-bottom: 25px;
    /* Para mostrar solo una línea del contexto */
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.read-manuscript-btn {
    display: inline-block;
    background: #8b764c;
    color: #faf8f3;
    padding: 10px 20px;
    border-radius: 6px;
    text-decoration: none;
    font-weight: 600;
    box-shadow: 0 2px 8px rgba(139, 118, 76, 0.2);
}

.pagination-vintage {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 30px;
    gap: 20px;
}

.pagination-btn {
    background: #8b764c;
    color: #faf8f3;
    padding: 10px 20px;
    border: none;
    border-radius: 6px;
    font-weight: 600;
    cursor: pointer;
    box-shadow: 0 2px 8px rgba(139, 118, 76, 0.2);
    transition: background-color 0.3s ease;
}

.pagination-btn:hover:not(:disabled) {
    background: #6d5a3a;
}

.pagination-btn:disabled {
    background: #ccc;
    cursor: not-allowed;
    opacity: 0.6;
}

.pagination-info {
    color: #4a3c1d;
    font-weight: 600;
    text-align: center;
}
</style>