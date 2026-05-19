# Algueiza Flight System

Sistema de gestión de vuelos en tiempo real implementado en Go, basado en estructuras de datos propias para garantizar eficiencia en consultas y actualizaciones.

---

## 🚀 Features

- Carga de vuelos desde CSV
- Consultas por código de vuelo
- Tablero de vuelos por rango de fechas
- Cola de prioridad de vuelos
- Búsqueda de conexiones entre aeropuertos
- Eliminación por rango temporal

---

## 🧠 Estructuras de datos

- **HashMap (diccionario):** acceso O(1) por código de vuelo  
- **Árbol binario de búsqueda (ABB):** orden temporal de vuelos  
- **Heap máximo:** cola de prioridad  
- **Lista enlazada:** almacenamiento auxiliar e iteración  
- **Pila / Cola:** estructuras de soporte  

---

## ⏱️ Complejidad de operaciones

| Operación            | Descripción                          | Complejidad |
|---------------------|--------------------------------------|-------------|
| agregar_archivo      | Inserción desde CSV                  | O(V log n)  |
| ver_tablero          | Consulta por rango temporal          | O(v)        |
| info_vuelo           | Búsqueda por código                  | O(1)        |
| prioridad_vuelos     | Top-K vuelos prioritarios            | O(n + K log n) |
| siguiente_vuelo      | Búsqueda de conexión                 | O(log n)    |
| borrar               | Eliminación por rango                | O(K log n)  |

---
## 🧾 Comandos y uso

El sistema funciona mediante comandos por entrada estándar (CLI). A continuación se describe brevemente cada uno:

- **agregar_archivo <archivo.csv>**  
  Carga un archivo CSV con vuelos al sistema. Cada ejecución puede agregar o actualizar vuelos existentes.

- **ver_tablero** <K> <asc/desc> <desde> <hasta> 
  Muestra hasta K vuelos dentro del rango de fechas indicado, ordenados por fecha de despegue en orden ascendente o descendente.

- **info_vuelo** código_vuelo  
  Muestra toda la información asociada a un vuelo específico identificado por su código.

- **prioridad_vuelos** K
  Muestra los K vuelos con mayor prioridad registrados en el sistema.

- **siguiente_vuelo** origen destino fecha  
  Devuelve el próximo vuelo directo entre dos aeropuertos a partir de una fecha dada.

- **borrar** desde hasta
  Elimina todos los vuelos cuya fecha de despegue esté dentro del rango indicado.

📌 Todas las fechas deben tener formato:  
`YYYY-MM-DDTHH:MM:SS`
---
## ⚙️ Uso

Compilar:
```bash
go build -o algueiza main.go
./algueiza
```
## Ejemplo de ejecucion en CLI

```bash
agregar_archivo data/vuelos.csv
prioridad_vuelos 5
ver_tablero 10 asc 2018-10-10T00:00:00 2018-10-10T23:59:59
```

## 🧪 Tests de los TDAS
```bash
go test ./... -v
```
