
# NatsChatCLI  
![GitHub top language](https://img.shields.io/github/languages/top/alerone/NatsChatCLI?color=%2377CDFF) ![GitHub last commit](https://img.shields.io/github/last-commit/alerone/NatsChatCLI?color=%23bc0bbf) ![GitHub Created At](https://img.shields.io/github/created-at/alerone/NatsChatCLI?color=%230dba69) ![GitHub repo size](https://img.shields.io/github/repo-size/alerone/NatsChatCLI?color=%23390385)

## 📝 Descripción  
_Tarea para la asignatura **SAD** sobre el seminario de **NATS**._

NatsChatCLI es un servidor de chat ligero que utiliza **NATS** para la comunicación entre canales y usuarios. 🗨️ Conecta dispositivos fácilmente y permite mensajes persistentes con la ayuda de **JetStream**.

---

## Instalación 🚀

Para comenzar a usar NatsChatCLI, sigue estos pasos:

1. **Inicia el servidor NATS con Docker Compose**:
   ```bash
   docker-compose up --build
   ```

2. **Compila el proyecto con Go**:
   ```bash
   go build .
   ```

3. **Conéctate a un canal y empieza a chatear**:
   ```bash
   ./natsChat <NATS_SERVER> <CHANNEL> <YOUR_NAME>
   ```

   ### Parámetros:
   - `NATS_SERVER`: Dirección del servidor NATS (por ejemplo, `localhost:4222`).
   - `CHANNEL`: Nombre del canal (p. ej., `hello.world`).
   - `YOUR_NAME`: Tu nombre de usuario. Debe cumplir:
     - ❌ No contener: `'.', '>', '*'`.
     - ⚠️ Ser único en el canal.

---

## ¿Cómo funciona NatsChat? 🧠 

Una vez conectado:
- Puedes **enviar mensajes** al canal.
- Verás los mensajes **guardados desde hace 1 hora** gracias a JetStream.

---

## Estructura del Proyecto 📂 

```plaintext
├── config/                  # Configuración de conexiones y streams
│   ├── cancelNatsConn.go       # Cancelar conexión a NATS
│   ├── connectToJetstream.go   # Crear el contexto Jetstream
│   ├── connectToNats.go        # Conexión a NATS
│   ├── createChatStream.go     # Crear un stream para el chat
│   ├── createConsumer.go       # Crear un consumidor
│   └── getClientArgs.go        # Obtener argumentos del cliente
│
├── models/                  # Definición de modelos de datos
│   └── clientConnection.go     # Modelo para los argumentos del cliente
│
├── service/                 # Lógica de servicios de consumo y publicación
│   ├── consume.go              # El consumidor muestra los mensajes por pantalla
│   └── publish.go              # Publicación de mensajes en el ChatStream
│
├── .gitignore               # Archivos y carpetas ignorados por Git
├── docker-compose.yml       # Configuración para Docker Compose
├── go.mod                   # Módulo de dependencias de Go
├── go.sum                   # Checksum de las dependencias
├── main.go                  # Punto de entrada principal de la aplicación
└── README.md                # Documentación del proyecto
```

---

## Detalles por Paquete 📦

### Paquete `config` 🔧 
Inicializa y configura:
- Conexión a NATS.
- Contexto JetStream para la persistencia.
- Stream para los mensajes del chat.
- Consumidor para leer mensajes.
- Argumentos del cliente.

👉 Este paquete es la base de la aplicación.

### Paquete `main` 🔑 
- **Punto de entrada principal**.
- Inicializa la configuración y las conexiones.
- Lanza:
  - Consumo de mensajes.
  - Entrada de teclado para enviar mensajes.

### Paquete `models` 🏗️ 
Define la estructura para los argumentos que los usuarios pasan a la aplicación.

### Paquete `service` 📡 
Contiene la lógica para:
- **Consumir mensajes**: Muestra mensajes del canal al usuario.
- **Publicar mensajes**: Envía mensajes al canal.

---

## Explicación de JetStream 🌐 

**JetStream** se utiliza para guardar mensajes durante 1 hora. Aquí te explico cómo funciona:

1. **Contexto JetStream**  
   - [connectToJetstream.go](./config/connectToJetstream.go): Configura el contexto de JetStream.

2. **Stream de Chat**  
   - [createChatStream.go](./config/createChatStream.go): Configura un stream con:
     ```go
     jetstream.StreamConfig{
         Name:        "chats",
         Description: "Stream for chatting",
         Subjects:    []string{"chats.>"},
         MaxAge:      time.Hour,
     }
     ```
     - Guarda los mensajes enviados a `chats.<CANAL>` durante 1 hora.

3. **Consumidor**  
   - [createConsumer.go](./config/createConsumer.go): Configura el consumidor para leer mensajes:
     ```go
     jetstream.ConsumerConfig{
         Name:          consumerName,
         Durable:       consumerName,
         AckPolicy:     jetstream.AckExplicitPolicy,
         FilterSubject: fmt.Sprintf("chats.%s", ClientConn.Channel),
         DeliverPolicy: jetstream.DeliverAllPolicy,
     }
     ```
     - **Nombre del consumidor**: `<CANAL>_<USUARIO>`.
     - **Filtros**: Solo consume mensajes del canal correspondiente.

4. **Consumo**  
   - [consume.go](./service/consume.go): Muestra los mensajes en pantalla para el usuario.

---

¡Disfruta chateando con **NatsChatCLI**! 🚀💬











