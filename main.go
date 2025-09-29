package main
import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    sqlc "tp/bd/sqlc"
)
func main() {
    // 1. Define el contenido HTML
    htmlContent, err1 := os.ReadFile("index.html")
    if err1 != nil {
        fmt.Println("Error al leer el archivo:", err1)
        return
    }
    // 2. Registra un manejador (handler) para la ruta raíz "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // 3. Establece la cabecera Content-Type
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        // 4. Escribe el HTML en la respuesta
        w.Write(htmlContent)
    })
    // 5. Define el puerto y muestra un mensaje
    port := ":8080"
    fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
    // 6. Inicia el servidor HTTP
    err := http.ListenAndServe(port, nil)
    if err != nil {
        fmt.Printf("Error al iniciar el servidor: %s\n", err)
    }
}
