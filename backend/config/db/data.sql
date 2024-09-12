-- Insertar datos de prueba en la tabla de usuarios
INSERT INTO users (username, email, created_at) VALUES
('alice', 'alice@example.com', CURRENT_TIMESTAMP),
('bob', 'bob@example.com', CURRENT_TIMESTAMP),
('carol', 'carol@example.com', CURRENT_TIMESTAMP);

-- Insertar datos de prueba en la tabla de chats
INSERT INTO chats (user_id, created_at, system_fingerprint, model_used, total_tokens) VALUES
(1, CURRENT_TIMESTAMP, 'fingerprint1', 'modelA', 1000),
(2, CURRENT_TIMESTAMP, 'fingerprint2', 'modelB', 1500),
(1, CURRENT_TIMESTAMP, 'fingerprint3', 'modelC', 2000);

-- Insertar datos de prueba en la tabla de mensajes
INSERT INTO messages (chat_id, role, content, created_at, finish_reason, prompt_tokens, completion_tokens, total_tokens) VALUES
(1, 'user', '¿Qué es el cambio climático?', CURRENT_TIMESTAMP, 'completed', 10, 50, 60),
(1, 'assistant', 'El cambio climático se refiere a los cambios a largo plazo en las temperaturas y patrones climáticos globales.', CURRENT_TIMESTAMP, 'completed', 5, 40, 45),
(2, 'user', '¿Cuáles son los beneficios del ejercicio?', CURRENT_TIMESTAMP, 'completed', 8, 55, 63),
(2, 'assistant', 'El ejercicio regular ofrece numerosos beneficios para la salud, como mejorar la salud cardiovascular, aumentar la fuerza muscular y reducir el estrés.', CURRENT_TIMESTAMP, 'completed', 7, 50, 57),
(3, 'user', '¿Qué es la inteligencia artificial?', CURRENT_TIMESTAMP, 'completed', 12, 60, 72),
(3, 'assistant', 'La inteligencia artificial (IA) es una rama de la informática que se centra en la creación de sistemas capaces de realizar tareas que normalmente requieren inteligencia humana.', CURRENT_TIMESTAMP, 'completed', 10, 55, 65);

-- Insertar datos de prueba en la tabla de historial de chat
INSERT INTO chat_history (chat_id, message_id, created_at) VALUES
(1, 1, CURRENT_TIMESTAMP),
(1, 2, CURRENT_TIMESTAMP),
(2, 3, CURRENT_TIMESTAMP),
(2, 4, CURRENT_TIMESTAMP),
(3, 5, CURRENT_TIMESTAMP),
(3, 6, CURRENT_TIMESTAMP);
