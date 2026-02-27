# Port Killer 🔪

Una herramienta de terminal (TUI) interactiva para visualizar y gestionar procesos que ocupan puertos en tu sistema, sin necesidad de recordar comandos complejos como `lsof` o `netstat`.

Construido con ❤️ usando **Go** y las librerías de **Charm**.

## 🚀 Características

- **Visualización en Tiempo Real**: Lista interactiva de todos los procesos ocupando puertos (TCP/UDP).
- **Filtro Inteligente**: Busca instantáneamente por puerto, nombre de proceso o PID.
- **Detalles Expandidos**: Ver información detallada de cada socket y proceso.
- **Gestión de Procesos**: Mata procesos directamente desde la interfaz con confirmación de seguridad.
- **Señales Personalizadas**: Elige entre `SIGTERM` (terminación elegante) o `SIGKILL` (forzado).
- **Estética Premium**: Interfaz colorida y moderna diseñada con Lip Gloss.

## 🛠️ Stack Tecnológico

- [Go](https://go.dev/) (1.22+)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — Motor de TUI.
- [Bubbles](https://github.com/charmbracelet/bubbles) — Componentes de tabla e inputs.
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — Estilizador de terminal.
- [Huh](https://github.com/charmbracelet/huh) — Formularios interactivos.

## 📦 Instalación

### Requisitos previos
- Go instalado en el sistema.

### Desde el código fuente
1. Clona este repositorio o descarga el código.
2. Compila el binario:
   ```bash
   go build -o port-killer
   ```
3. (Opcional) Instálalo globalmente:
   ```bash
   go install
   ```

> [!TIP]
> Si después de instalarlo globalmente el comando `port-killer` no es reconocido, asegúrate de que el directorio bin de Go esté en tu `PATH`. Puedes agregarlo con:
> ```bash
> export PATH=$PATH:$(go env GOPATH)/bin
> ```
> Agrega esa línea a tu archivo de configuración de shell (ej. `~/.zshrc` o `~/.bashrc`) para que sea permanente.

## 🎮 Uso

Para iniciar la herramienta simplemente ejecuta:
```bash
port-killer
```

> [!NOTE]
> Si estás ejecutando desde la carpeta del proyecto sin instalar globalmente, usa `./port-killer`.

### Atajos de teclado

| Tecla | Acción |
|---|---|
| `↑` `↓` | Navegar entre procesos |
| `Enter` | Ver detalles del proceso |
| `K` | Matar el proceso seleccionado |
| `F` | Abrir filtro interactivo |
| `R` | Refrescar la lista |
| `Esc` | Volver atrás / Cancelar |
| `Q` | Salir de la aplicación |

## 📸 Screenshots (Mockup)

```text
┌─────────────────────────────────────────────────────────────┐
│  ⚡ Port Killer                          Refresh: r  Quit: q │
├────────┬──────────┬────────┬────────────┬────────────────────┤
│  PORT  │ PROTOCOL │  PID   │  PROCESS   │       STATUS       │
├────────┼──────────┼────────┼────────────┼────────────────────┤
│  3000  │   TCP    │ 18432  │  node      │    🟢 LISTENING    │
│  5432  │   TCP    │  892   │  postgres  │    🟢 LISTENING    │
│  8080  │   TCP    │ 21050  │  python3   │    🟢 LISTENING    │
└────────┴──────────┴────────┴────────────┴────────────────────┘
```

## ⚖️ Licencia

Este proyecto está bajo la Licencia MIT. ¡Siéntete libre de contribuir!
