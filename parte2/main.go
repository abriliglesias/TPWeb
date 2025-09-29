package main
import (
    "context"        // para context.Background()
    "database/sql"   // para sql.Open()
    "fmt"            // para fmt.Println()
    "log"    
    _ "github.com/lib/pq"
    sqlc "tp/bd/sqlc"
)
func main(){
    connStr := "postgresql://abril:1234@localhost:5432/bd?sslmode=disable"
    //connStr := "user=abril password=1234 dbname=bd"
    bd, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("failed to connect to DB: %v", err)
    }
    defer bd.Close()
    queries := sqlc.New(bd)
    ctx := context.Background()

    createdLibro, err := queries.CreateLibro(ctx, // Create
    sqlc.CreateLibroParams{
        Titulo: "La tia Cosima",
        Autor: "Florencia Bonelli",
        Descripcion: "Cósima es una mujer en la plenitud de la vida. Psicóloga de profesión y especializada en el tratamiento del autismo infantil, posee una fundación de enorme prestigio, donde se respira un ambiente cuidado y buen humor. Allí trabaja con perros especialmente adiestrados para ayudar a los niños con alguna condición del espectro autista. Es feliz con su trabajo, con sus amigos y sus sobrinos, a quienes dedica los pocos ratos libres de que dispone.",
        Valoracion: sql.NullInt32{Int32: 4, Valid: true,},
        Anio: int32(2020),
        GeneroPrincipal: "Romance",
    })
    if err != nil {
        log.Fatalf("failed to create libro: %v", err)
    }
    fmt.Printf("Created libro: %+v\n", createdLibro)
    
    libro, err := queries.GetLibroByID(ctx, createdLibro.ID) // Read One
    if err != nil {
        log.Fatalf("failed to get libro: %v", err)
    }
    fmt.Printf("Retrieved libro: %+v\n", libro)
}
