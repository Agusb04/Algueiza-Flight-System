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

## ⚙️ Uso

Compilar:
```bash
go build -o algueiza main.go
./algueiza
```
## Ejemplo de ejecucion en CLI

```bash
agregar_archivo data/vuelos-algueiza-parte-01.csv
prioridad_vuelos 5
ver_tablero 10 asc 2018-10-10T00:00:00 2018-10-10T23:59:59
```

## 🧪 Tests
```bash
go test ./... -v
```

Incluye tests unitarios de las estructuras de datos (TDAs), las funciones auxiliares, el sistema de vuelos, y el dispatch de comandos.