-- Elimina las tablas en el orden necesario para evitar conflictos
-- Asegúrate de ejecutar esta sección solo si deseas eliminar las tablas antes de recrearlas.

-- Elimina la tabla de historial de chat
DROP TABLE IF EXISTS chat_history CASCADE;

-- Elimina la tabla de mensajes
DROP TABLE IF EXISTS messages CASCADE;

-- Elimina la tabla de chats
DROP TABLE IF EXISTS chats CASCADE;

-- Elimina la tabla de usuarios
DROP TABLE IF EXISTS users CASCADE;

-- Re-crea las tablas en el orden necesario

-- Tabla para almacenar los usuarios (opcional)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla para almacenar los chats
CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    system_fingerprint VARCHAR(255),
    model_used VARCHAR(255),
    total_tokens INT
);

-- Tabla para almacenar los mensajes (preguntas y respuestas)
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    chat_id INT REFERENCES chats(id) ON DELETE CASCADE,
    role VARCHAR(50) CHECK (role IN ('user', 'assistant')), -- 'user' para preguntas, 'assistant' para respuestas
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    finish_reason VARCHAR(50), -- Para almacenar el motivo por el cual se detuvo el chat (si aplica)
    prompt_tokens INT,
    completion_tokens INT,
    total_tokens INT
);

-- Tabla para almacenar el historial de interacciones por cada chat
CREATE TABLE IF NOT EXISTS chat_history (
    id SERIAL PRIMARY KEY,
    chat_id INT REFERENCES chats(id) ON DELETE CASCADE,
    message_id INT REFERENCES messages(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
