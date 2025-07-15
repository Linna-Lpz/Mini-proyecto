# Listado de artículos
## Descripción
Pequeño proyecto donde un usuario puede ver una lista de artículos, entrar en cada uno de ellos y dejar un comentario.

## Estructura del Proyecto

```
├── README.md
├── go-template-structured/          # Backend API en Go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── controllers/                 # Controladores de la API
│   │   ├── controlArticle.go
│   │   ├── controlComment.go
│   ├── helpers/                     # Token
│   │   ├── token.go                 
│   ├── middleware/                  # Middleware
│   │   ├── auth.go    
│   ├── models/                      # Modelos de datos
│   │   ├── article.go
│   │   ├── comment.go
│   │   └── user.go
│   ├── routes/                      # Configuración de rutas
│   │   └── routes.go
│   ├── services/                    # Servicios (Base de datos)
│   │    └── mongo.go
│   └── utils/                       
│       └── filter.go
└── nuxt3-template/                  # Frontend en Nuxt 3
    ├── nuxt.config.ts
    ├── package.json
    ├── components/
    ├── composables/
    ├── pages/
    ├── plugins/
    └── public/
```

## Tecnologías Utilizadas

### Backend (Go)
- **Go**: Lenguaje de programación principal
- **MongoDB**: Base de datos NoSQL
- **Gin** (presumiblemente): Framework web para Go

### Frontend (Nuxt 3)
- **Nuxt 3**: Framework de Vue.js
- **TypeScript**: Tipado estático
- **Vue 3**: Framework de JavaScript reactivo

## Instalación y Configuración

### Backend (Go)

1. Navega al directorio del backend:
```bash
cd go-template-structured
```

2. Instala las dependencias:
```bash
go mod tidy
```

3. Ejecuta el servidor:
```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080` (o el puerto configurado).

### Frontend (Nuxt 3)

1. Navega al directorio del frontend:
```bash
cd nuxt3-template
```

2. Instala las dependencias:
```bash
npm install
```

3. Ejecuta el servidor de desarrollo:
```bash
npm run dev
```

El frontend estará disponible en `http://localhost:3000`.

## Base de Datos

El proyecto utiliza MongoDB como base de datos. Asegúrate de que MongoDB esté ejecutándose antes de iniciar el backend.

### Modelos disponibles:
- **Article**: Gestión de artículos
- **Comment**: Sistema de comentarios de cada artículo
