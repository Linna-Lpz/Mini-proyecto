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
                            <form v-if="isLoggedIn" @submit.prevent="handleSubmit">
                                <p><strong>Autor:</strong> {{ author }}</p>
                                <textarea v-model="text" placeholder="Comentario" required></textarea>
                                <div>
                                    <button type="submit">Enviar comentario</button>
                                </div>
                            </form>
                            <form v-else>
                                <p>Para comentar, por favor <NuxtLink to="/login">inicia sesión</NuxtLink> o <NuxtLink to="/register">regístrate</NuxtLink>.</p>
                            </form>
                        </div>
                        <div class="filter-section">
                            <input 
                                type="text" 
                                v-model="searchAuthor" 
                                placeholder="Buscar por autor..." 
                                class="search-input"
                                @keyup.enter="fetchComments"
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
                            <button class="search-btn" @click="fetchComments">
                                Aplicar filtro
                            </button>
                            <button class="search-btn" @click="cleanFilters">
                                Limpiar filtro
                            </button>
                        </div>
                        <div class="comments-section">
                            <div v-if="comments.length === 0" class="no-articles-vintage">
                                Aún no hay comentarios ¡Se el primero en comentar!
                            </div>
                            <div v-else>
                                <div 
                                    class="comment-card"
                                    v-for="(comment) in comments"
                                    :key="comment.id"
                                >
                                    <div class="comment-header">
                                        <span class="comment-author">{{ comment.author_name || 'Anónimo' }}</span>
                                        <span class="comment-date">{{ comment.created_at.slice(0, 10) }}</span>
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

import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { usePostComment } from '../composables/postComment'

const route = useRoute()
const id = route.query.id as string

// @ts-ignore: useFetch es global en Nuxt 3, pero TypeScript puede no reconocerlo sin tipos de Nuxt
const { data, pending: loading } = useFetch(`http://localhost:8080/articles/${id}`)
const article = computed(() => data.value ?? null)

const { postComment } = usePostComment()

const authorId = ref('')
const author = ref('')
const text = ref('')
const comments = ref<CommentWithAuthor[]>([])
const searchAuthor = ref('')
const fromDate = ref('')
const toDate = ref('')

const handleSubmit = async () => {
  await postComment(id, { text: text.value })
  author.value = ''
  text.value = ''
}

const isLoggedIn = ref(false)

onMounted(() => {
    const token = localStorage.getItem('token')
    if (token) {
        try {
            const payload = JSON.parse(atob(token.split('.')[1])) // Decodificar el token JWT
            author.value = payload.name // Asignar el nombre del usuario al campo de autor
            authorId.value = payload.user_id // Asignar el ID del usuario al campo de autorId
            isLoggedIn.value = true
        } catch (error) {
            isLoggedIn.value = false
        }
    }

    fetchComments()
})

// Función para obtener comentarios según los filtros actuales
const fetchComments = async () => {
  const params = new URLSearchParams()
  console.log("Fetching comments with params:", params.toString())
  if (searchAuthor.value) params.append("author", searchAuthor.value)
  if (fromDate.value) params.append("from", fromDate.value)
  if (toDate.value) params.append("to", toDate.value)
  console.log("añadidos: ", params.toString())

  const res = await fetch(`http://localhost:8080/articles/${id}/comments?${params}`)
  console.log("Response from comments API:", res)
  const json = await res.json()
  console.log("JSON response from comments API:", json)
  comments.value = json.comments || []
  console.log("Comentarios obtenidos:", comments.value)
}

const cleanFilters = async () => {
  searchAuthor.value = ''
  fromDate.value = ''
  toDate.value = '' 
  await fetchComments()
}

interface CommentWithAuthor {
  id: string;
  article_id: string;
  author_id: string;
  text: string;
  created_at: string;
  author_name?: string;
}

</script>