# Mortal Kombat Characters - gRPC Project

Este proyecto implementa un sistema basado en gRPC para gestionar información sobre los personajes de **Mortal Kombat**. Utiliza columnas como `Name` (nombre del personaje), `Type` (lugar de origen), y `Skills` (habilidades del personaje). El servicio se implementó en **Go**, empaquetado con **Docker** y probado localmente utilizando **Postman** en `localhost:50051`.

---

## Funcionalidades

Este proyecto gRPC incluye los siguientes servicios:

1. **Unary Service**: Buscar un personaje específico por su nombre.
2. **Server Streaming Service**: Obtener la lista completa de personajes.
3. **Client Streaming Service**: Agregar múltiples personajes al sistema en una sola solicitud.
4. **Bidirectional Streaming Service**: Intercambiar información de personajes entre cliente y servidor.

---

## Requisitos

- **Go** (>= 1.23.3)
- **Protocol Buffers** (protoc v3)
- **Docker**
- **Postman**
- **Azure services**

---

## Configuración Local
para trabajar con el proyecto de manera local 
- Instalar las dependencias necesarias
-Clonar el repositorio del proyecto
-Ejecutar Postman
-Usar en terminal go run main.go
El servidor estará escuchando en `localhost:50051`.

### Pruebas con Postman

1. Configura Postman para enviar solicitudes gRPC:
   - Abre Postman y selecciona "New Request" > "gRPC".
   - Usa el endpoint `localhost:50051`.
2. Prueba los diferentes métodos:
   - **Unary**: Selecciona el método `GetCharacter` y envía un payload con el nombre del personaje.
   - **Server Streaming**: Selecciona `ListCharacters` para obtener todos los personajes.
   - **Client Streaming**: Usa `AddCharacters` para enviar múltiples personajes en una solicitud.
   - **Bidirectional Streaming**: Prueba `ExchangeCharacters` para interactuar en tiempo real con el servidor.

---

## Empaquetado con Docker

El proyecto fue empaquetado en esta herramienta donde se crearan imagenes y contenedores Docker para su despliegue en Azure (para eso claro es el archivo Dockerfile).

## Despliegue en Azure
Antes de empezar se debe tener que contar con que ya hayas instalado en tu computadora las siguientes herramientas:
- Azure CLI MSI.
El proyecto puede ser desplegado en Azure App Service utilizando contenedores Docker. Asegúrate de tener configurada tu cuenta de Azure y sigue estos pasos:

1. Inicia sesión en Azure

2. Crea un grupo de recursos

3. Crea un registro de contenedor

4. Publica la imagen en el registro

5. Crea el servicio App Service

---

## Estructura de Datos

| **Name**           | **Type**       | **Skills**                    |
|---------------------|----------------|--------------------------------|
| Scorpion            | Mundo_Exterior   | fire        |
| Sub-Zero           | Mundo_Exterior    | Ice               |
| Raiden              | Heavens       | Thunder    |
| Liu Kang            | Earth  | fire        |
| Shang Tsung         | Mundo_Exterior      | Spirit          |

---

## Archivo `.proto`

El archivo `.proto` está disponible en el repositorio para definir la interfaz del servicio gRPC. Contiene las definiciones para los servicios Unary, Server Streaming, Client Streaming y Bidirectional Streaming.

