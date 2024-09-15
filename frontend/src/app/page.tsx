"use client";

import { useRouter } from 'next/navigation'

import ChatBox from '@/components/chat-box/ChatBox';
import Topvar from '@/components/topvar/Topvar';
import React, { useState } from 'react';

export type Chat = {
  id: number;
  user_id: number;
  title: string;
  created_at: string;
  system_fingerprint: string;
  model_used: string;
  total_tokens: number;
};

// Stores
import { useNewChatStore } from "@/stores/create-new-chat"

type Props = {};

const HomePage = (props: Props) => {
  const router = useRouter()

  const [title, setTitle] = useState<string>("title default"); // Para capturar el título del chat

  const { markChatAsCreated, } = useNewChatStore((state) => ({
    markChatAsCreated: state.markChatAsCreated,
  }))

  // Función para crear un nuevo chat
  const sendMessage = async (input: string) => {
    const chatData = {
      title: title,
      lastMessage: input, // The message text
      user_id: 1, // Adjust this value based on your actual user_id
      system_fingerprint: 'fingerprint1',
      model_used: 'modelA',
      total_tokens: 1000,
    };
  
    try {
      // Step 1: Create the chat
      const chatResponse = await fetch('http://localhost:4000/chats', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(chatData),
      });
  
      if (!chatResponse.ok) {
        throw new Error('Error creating chat');
      }
  
      const chatResult = await chatResponse.json();
      console.log('Full API response for chat:', chatResult);
      console.log('Chat created:', chatResult.data);
  
      // Ensure chatResult.data and chatResult.data.id are available
      if (chatResult.data && chatResult.data.id) {
        const newChatId = chatResult.data.id;
  
        // Create the first message using the newChatId
        const messageData = {
          chat_id: newChatId,
          role: 'user',
          content: input,
          prompt_tokens: 100,
          completion_tokens: 900,
          total_tokens: 1000,
        };
  
        const messageResponse = await fetch('http://localhost:4000/messages', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(messageData),
        });
  
        if (!messageResponse.ok) {
          throw new Error('Error creating message');
        }
  
        const messageResult = await messageResponse.json();
        console.log('Message created:', messageResult.data);
  
        // Mark the chat as created and navigate to the chat page
        markChatAsCreated();
        router.push(`/chat/${newChatId}`);
      } else {
        console.error('Chat ID is missing from the response data');
      }
    } catch (error) {
      console.error('Failed to create chat or message:', error);
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
