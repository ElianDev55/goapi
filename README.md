## 🚀 ¿Qué es GoAPI CLI?

GoAPI CLI es una herramienta de línea de comandos que te permite **crear y escalar proyectos backend en Go** de manera rápida, utilizando el framework **Gin** y el ORM **GORM**.

Esta herramienta está diseñada para que puedas concentrarte en el desarrollo de funcionalidades sin preocuparte por configurar desde cero la estructura de tu aplicación.

---

## 💠 ¿Cómo se usa?

### 1. **Crear un nuevo proyecto**

```bash
goapi new mi-backend
```

Esto generará una estructura de carpetas organizada con todo lo necesario para comenzar: `main.go`, módulos en `internal/`, conexión a base de datos en `pkg/db/`, archivos de entorno, etc.

---

### 2. **Agregar un módulo (por ejemplo, user)**

```bash
goapi generate m/module user
```

Esto creará automáticamente todo lo relacionado con el módulo `user`:

- Modelo de base de datos
- Repositorio
- Servicio
- Endpoint
- Rutas

Todo quedará organizado en `internal/user/`.

---

### 3. **Crear archivos o carpetas personalizados**

```bash
goapi generate logger --f        # Crea el archivo logger.go
goapi generate utils --d         # Crea la carpeta utils/
goapi create utils/logger.go     # Crea el archivo si no existe
```

---

### 4. **Ver ayuda y comandos disponibles**

```bash
goapi --help
```

Para ver la versión actual del CLI:

```bash
goapi --version
```

---

## 📂 ¿Qué incluye el proyecto generado?

Una estructura organizada lista para producción:

- `main.go`: punto de entrada de la app.
- `internal/`: módulos divididos por entidad (user, product, etc.).
- `pkg/`: paquetes reutilizables como la base de datos.
- `tools/seed/`: scripts para poblar la base de datos.
- `.env`: configuración de entorno.
- `go.mod`: dependencias del proyecto.

Cada módulo generado tiene:

- `model.go`
- `repository.go`
- `service.go`
- `endpoint.go`
- `router.go`

---

### ✅ Así de simple: Genera tu proyecto → Crea módulos → ¡Empieza a programar!
