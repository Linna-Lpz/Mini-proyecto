<template>
    <div class="vintage-container">
        <div class="parchment-bg">
            <NuxtLink 
                class="read-manuscript-btn"
                to="/articlesList"
            > Volver
            </NuxtLink>

            <div v-if="loading" class="loading-vintage">
                Cargando artículos
            </div>
            
            <div v-else>
                <div v-if="!article" class="no-articles-vintage">
                    El artículo no existe o no se pudo cargar.
                </div>
                <div v-else class="">
                    <div class="manuscript-title">{{ article.article.title}}</div>
                    <div class="manuscript-content">{{ article.article.content }}</div>
                    <div>
                        <div class="parchment-bg">
                            <h3>Deja tu comentario</h3>
                            
                        </div>
                        <div v-if="!article.comments"> Aún no hay comentarios ¡Se el primero en comentar!</div>
                        <div v-else >
                            <div class="parchment-bg" v-for="(comment) in article.comments" :key="comment.id">
                                {{ comment.author }}: {{ comment.text }}: {{ comment.created_at }}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>


<script lang="ts" setup>
import { useRoute } from 'vue-router'
const route = useRoute()
const id = route.query.id
const { data, pending: loading } = await useFetch(`http://localhost:8080/articles/${id}`)
const article = computed(() => data.value ?? null)

const { data: commentsData } = await useFetch(`http://localhost:8080/articles/${id}/comments`)
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
</style>