# 🏋️‍♂️ PAC-ENT: Sistema de Gestión para Gimnasios

Este proyecto es una API REST desarrollada en Go utilizando el framework **[Ent](https://entgo.io/)** y conectada a una base de datos **PostgreSQL**.
Su propósito es permitir el manejo de clientes, rutinas, membresías y entrenadores de un gimnasio.

## 📦 Estructura del proyecto 
    pac-ent/
    │
    ├── cmd/pac-ent/main.go # Punto de entrada de la aplicación
    ├── internal/
    │ └── handlers/ # Handlers para las rutas HTTP
    │ ├── clientes.go
    │ ├── entrenadores.go
    │ ├── membresias.go
    │ └── rutinas.go
    │
    ├── ent/ # Código generado por el framework Ent
    ├── ent/schema/ # Esquemas de Ent (Clientes, Rutinas, etc.)
    │
    ├── .env # Variables de entorno (conexión a DB)
    ├── go.mod # Módulo Go
    └── go.sum # Suma de dependencias
    
🧠 Tecnologías utilizadas
    Go
    Ent Framework
    PostgreSQL
    net/http

👨‍💻 Autor
    Desarrollado por Julinho.
    Github: @jumata96

📝 Licencia
    Este proyecto está bajo la licencia MIT.
    ¡Úsalo, modifícalo y mejóralo!
