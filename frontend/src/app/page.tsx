"use client";

import ChatBox from '@/components/chat-box/ChatBox';
import Topvar from '@/components/topvar/Topvar';
import React, { useState } from 'react';

type Props = {};

const HomePage = (props: Props) => {
  const [title, setTitle] = useState<string>("title default"); // Para capturar el título del chat

  // Función para crear un nuevo chat
  const sendMessage = async (input: string) => {
    // Datos que serán enviados a la API para crear un nuevo chat

    const chatData = {
      title: title,
      lastMessage: input,
      user_id: 1, // Aquí puedes definir el user_id que corresponda
      system_fingerprint: 'fingerprint1', // Puedes cambiar esto o generarlo dinámicamente
      model_used: 'modelA', // Modelo que se esté utilizando
      total_tokens: 1000, // Número de tokens o puedes calcularlo
    };

    try {
      const response = await fetch('http://localhost:4000/chats', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(chatData), // Enviar los datos en formato JSON
      });

      if (!response.ok) {
        throw new Error('Error creating chat');
      }

      const data = await response.json();
      console.log('Chat created:', data);
      // Puedes hacer algo aquí después de que el chat se haya creado, como limpiar los inputs o hacer un refetch.
    } catch (error) {
      console.error('Failed to create chat:', error);
    }
  };

  return (
    <section className="bg-secondary h-screen w-full p-4 flex-1 flex flex-col">
      <Topvar />
      <div className="flex-1 m-6" />
      <ChatBox handleButton={sendMessage} />
    </section>
  );
};

export default HomePage;
