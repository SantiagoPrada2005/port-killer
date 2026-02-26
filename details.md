# Port Killer 🔪

Una herramienta interactiva para visualizar y gestionar procesos que ocupan puertos en tu sistema, sin necesidad de recordar comandos como `lsof` o `netstat`.

---

## Stack Tecnológico

**Lenguaje y runtime**
- Go 1.22+

**Librerías de Charm**
- **Bubble Tea** — motor principal de la TUI, maneja el loop de eventos y el estado de la aplicación
- **Bubbles** — componentes de tabla (para listar puertos) y spinner (mientras carga la información)
- **Lip Gloss** — estilos visuales: colores por estado, bordes, layout general
- **Huh** — formularios de confirmación antes de matar un proceso

**Sistema operativo**
- `os/exec` — ejecuta comandos del sistema (`lsof`, `ss`, `netstat`) para obtener la información de puertos
- `syscall` — envía señales al proceso (`SIGTERM`, `SIGKILL`)
- Compatible con macOS y Linux

**Almacenamiento**
- Sin base de datos. Toda la información es en tiempo real desde el sistema operativo.

**Distribución**
- Binario único compilado con `go build`
- Instalable con `go install` o descargando el binario desde GitHub Releases

---

## Pantallas y Menú de Opciones

---

### 🏠 Pantalla Principal — Lista de Puertos

Al abrir la app, muestra una tabla interactiva con todos los puertos en uso:

```
┌─────────────────────────────────────────────────────────────┐
│  ⚡ Port Killer                          Refresh: r  Quit: q │
├────────┬──────────┬────────┬────────────┬────────────────────┤
│  PORT  │ PROTOCOL │  PID   │  PROCESS   │       STATUS       │
├────────┼──────────┼────────┼────────────┼────────────────────┤
│  3000  │   TCP    │ 18432  │  node      │    🟢 LISTENING    │
│  5432  │   TCP    │  892   │  postgres  │    🟢 LISTENING    │
│  8080  │   TCP    │ 21050  │  python3   │    🟢 LISTENING    │
│  6379  │   TCP    │  901   │  redis     │    🟢 LISTENING    │
│  9000  │   TCP    │ 33201  │  php-fpm   │    🟡 ESTABLISHED  │
└────────┴──────────┴────────┴────────────┴────────────────────┘

  ↑↓ Navegar   Enter Detalles   K Matar   F Filtrar   S Ordenar
```

**Atajos disponibles:**

| Tecla | Acción |
|---|---|
| `↑` `↓` | Navegar entre procesos |
| `Enter` | Ver detalles del proceso |
| `K` | Matar el proceso seleccionado |
| `F` | Abrir filtro por puerto, nombre o PID |
| `S` | Ordenar por columna |
| `R` | Refrescar la lista |
| `Q` | Salir |

---

### 🔍 Pantalla de Detalles

Al presionar `Enter` sobre un proceso, muestra información expandida:

```
┌─────────────────────────────────────┐
│  Detalles del Proceso               │
├─────────────────────────────────────┤
│  Proceso   : node                   │
│  PID       : 18432                  │
│  Puerto    : 3000                   │
│  Protocolo : TCP                    │
│  Estado    : LISTENING              │
│  Usuario   : carlos                 │
│  Comando   : node server.js         │
│  Directorio: /home/carlos/proyecto  │
│  Iniciado  : hace 2 horas           │
├─────────────────────────────────────┤
│  [ K ] Matar    [ Esc ] Volver      │
└─────────────────────────────────────┘
```

---

### ⚠️ Confirmación antes de matar

Al presionar `K`, Huh muestra un formulario de confirmación para evitar accidentes:

```
┌──────────────────────────────────────────┐
│                                          │
│   ¿Matar el proceso "node" (PID 18432)?  │
│   Puerto: 3000                           │
│                                          │
│   Tipo de señal:                         │
│   ○ SIGTERM  (terminación elegante)      │
│   ● SIGKILL  (forzar cierre)             │
│                                          │
│   [ Confirmar ]       [ Cancelar ]       │
└──────────────────────────────────────────┘
```

---

### 🔎 Filtro Interactivo

Al presionar `F`, aparece un input para filtrar en tiempo real:

```
  Filtrar: 3_
  ┌────────────────────────────────────────┐
  │  3000  TCP  18432  node   LISTENING    │
  │  3306  TCP   910   mysql  LISTENING    │
  └────────────────────────────────────────┘
  Filtrando por: puerto, proceso o PID
```

---

### ✅ Feedback tras matar un proceso

Mensaje de resultado que aparece brevemente antes de volver a la lista:

```
  ✓ Proceso "node" (PID 18432) terminado correctamente.
    Puerto 3000 ahora está libre.
```

O en caso de error:

```
  ✗ No se pudo terminar el proceso. Permiso denegado.
    Intenta ejecutar port-killer con sudo.
```

---

## Estructura del Proyecto

```
port-killer/
├── main.go
├── internal/
│   ├── ports/
│   │   └── scanner.go      # lógica para leer puertos del sistema
│   ├── process/
│   │   └── killer.go       # lógica para matar procesos
│   └── ui/
│       ├── table.go        # pantalla principal
│       ├── detail.go       # pantalla de detalles
│       ├── filter.go       # componente de filtro
│       ├── confirm.go      # formulario de confirmación
│       └── styles.go       # todos los estilos de Lip Gloss
└── README.md
```

---

Tiene todo lo necesario para practicar las principales librerías de Charm en un solo proyecto, y es algo que usarías en tu día a día como desarrollador. ¿Lo arrancamos?