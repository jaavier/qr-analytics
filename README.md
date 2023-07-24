# Idea
La idea es que para 1 url se puedan generar códigos QR que se pueden poner en distintos lugares y poder medir cuál de ellos tiene
más impacto, a que hora y si es posible, desde que latitud y longitud

Cada URL tiene estas propiedades:

- Address: Av Los Leones 11
- Additional Instructions: Dentro de la universidad, 2do piso
- Analytics


type Analytics struct {
    Views int
    Referer string
    Latitude float64
    Longitude float64
}