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
                    <h1 class="manuscript-title">{{ article.article.title}}</h1>
                    <div class="manuscript-content-form">{{ article.article.content }}</div>
                    <div>
                        <div>
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

// @ts-ignore: useFetch es global en Nuxt 3, pero TypeScript puede no reconocerlo sin tipos de Nuxt
const { data, pending: loading } = useFetch(`http://localhost:8080/articles/${id}`)
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