"use client"

import React, { useState } from 'react'

// Definir el tipo para los mensajes del chat
interface ChatMessage {
  role: 'user' | 'assistant';
  content: string;
}

const HomePage: React.FC = () => {
  const [message, setMessage] = useState<string>('');         // Tipo explícito para el mensaje
  const [chatHistory, setChatHistory] = useState<ChatMessage[]>([]);  // Tipo explícito para el historial del chat

  // Función para manejar el envío del mensaje
  const handleSendMessage = async () => {
    if (message.trim() === '') return;

    try {
      // // Enviar el mensaje al endpoint utilizando fetch
      // const response = await fetch('http://localhost:4000/chat/completion', {
      //   method: 'POST',
      //   headers: {
      //     'Content-Type': 'application/json',
      //   },
      //   body: JSON.stringify({
      //     role: 'user',
      //     content: message,
      //   }),
      // });

      // // Convertir la respuesta a JSON
      // const data = await response.json();

      // // Obtener el mensaje del asistente de la respuesta
      // const assistantMessage = data.choices[0].message.content;

      const assistantMessage = "¡Hola! ¿Cómo puedo ayudarte hoy?"

      // Actualizar el historial del chat con el mensaje del usuario y la respuesta del asistente
      setChatHistory((prevChatHistory) => [
        ...prevChatHistory,
        { role: 'user', content: message },
        { role: 'assistant', content: assistantMessage },
      ]);

      // Limpiar el campo del mensaje
      setMessage('');

    } catch (error) {
      console.error('Error sending message:', error);
    }
  };

  // Función para manejar la tecla 'Enter' en el campo de entrada usando 'onKeyDown'
  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      handleSendMessage();
    }
  };

  return (
    <div>
      <h1>Chat</h1>
      <div className="chat-box">
        {/* Mostrar el historial del chat */}
        {chatHistory.map((msg, index) => (
          <div key={index} className={msg.role === 'user' ? 'user-message' : 'assistant-message'}>
            <strong>{msg.role === 'user' ? 'You' : 'Assistant'}:</strong> {msg.content}
          </div>
        ))}
      </div>

      {/* Campo de entrada de texto */}
      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        onKeyDown={handleKeyDown}
        placeholder="Type your message..."
      />
      <button onClick={handleSendMessage}>Send</button>
    </div>
  );
}

export default HomePage;
