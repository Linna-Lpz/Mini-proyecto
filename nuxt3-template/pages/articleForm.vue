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
                            <form @submit.prevent="handleSubmit">
                                <input v-model="author" placeholder="Tu nombre" required />
                                <textarea v-model="text" placeholder="Comentario" required></textarea>
                                <button type="submit">Enviar comentario</button>
                            </form>
                        </div>
                        <div class="comments-section">
                            <div v-if="!article.comments || article.comments.length === 0" class="no-articles-vintage">
                                Aún no hay comentarios ¡Se el primero en comentar!
                            </div>
                            <div v-else>
                                <div 
                                    class="comment-card"
                                    v-for="(comment) in article.comments"
                                    :key="comment.id"
                                >
                                    <div class="comment-header">
                                        <span class="comment-author">{{ comment.author }}</span>
                                        <span class="comment-date">{{ comment.created_at }}</span>
                                    </div>
                                    <div class="comment-body">{{ comment.text }}</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>


<script lang="ts" setup>
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'
import { usePostComment } from '../composables/postComment'

const route = useRoute()
const id = route.query.id as string
const { data, pending: loading } = await useFetch(`http://localhost:8080/articles/${id}`)
const article = computed(() => data.value ?? null)

const { postComment } = usePostComment()

const author = ref('')
const text = ref('')

const handleSubmit = async () => {
  const created_at = new Date().toISOString().slice(0, 10)
  await postComment(id, { author: author.value, text: text.value, created_at })
  author.value = ''
  text.value = ''
}
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

.comments-section {
    margin-top: 30px;
}

.comment-card {
    background: #fdfcf7;
    border: 1px solid #e0d6b8;
    border-radius: 8px;
    padding: 18px 20px;
    margin-bottom: 16px;
    box-shadow: 0 2px 8px rgba(139, 118, 76, 0.08);
}

.comment-header {
    display: flex;
    justify-content: space-between;
    margin-bottom: 6px;
    font-size: 15px;
    color: #5d4e2a;
}

.comment-author {
    font-weight: 600;
}

.comment-date {
    font-size: 13px;
    color: #a89b7c;
}

.comment-body {
    color: #4a3c1d;
    font-size: 16px;
    line-height: 1.4;
}

form {
    display: flex;
    flex-direction: column;
    gap: 12px;
    background: #fdfcf7;
    border: 1px solid #e0d6b8;
    border-radius: 8px;
    padding: 20px 24px;
    margin-bottom: 24px;
    box-shadow: 0 2px 8px rgba(139, 118, 76, 0.08);
}

form input,
form textarea {
    font-size: 16px;
    padding: 10px;
    border: 1px solid #d4c4a0;
    border-radius: 6px;
    background: #fffdfa;
    color: #4a3c1d;
    resize: none;
}

form textarea {
    min-height: 60px;
}

form button {
    align-self: flex-end;
    background: #8b764c;
    color: #faf8f3;
    padding: 8px 18px;
    border: none;
    border-radius: 6px;
    font-weight: 600;
    cursor: pointer;
    box-shadow: 0 2px 8px rgba(139, 118, 76, 0.12);
    transition: background 0.2s;
}

form button:hover {
    background: #a48d5f;
}
</style>