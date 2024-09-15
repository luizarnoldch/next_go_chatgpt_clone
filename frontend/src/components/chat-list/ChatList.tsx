"use client";

// Lib
import React, { useEffect, useState, useCallback } from 'react';
import ChatListItem from './ChatListItem';

// Stores
import { useNewChatStore } from "@/stores/create-new-chat"

export type Chat = {
  id: number;
  user_id: number;
  title: string;
  created_at: string;
  system_fingerprint: string;
  model_used: string;
  total_tokens: number;
};

export type ChatResponse = {
  data: Chat[];
  message: string;
  status: string;
};


const ChatList: React.FC = () => {
  const [chats, setChats] = useState<Chat[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const { isNewChatCreated, markChatAsNotCreated } = useNewChatStore((state) => ({
    isNewChatCreated: state.isNewChatCreated,
    markChatAsNotCreated: state.markChatAsNotCreated,
  }))

  // Fetch chats from the backend
  const fetchChats = useCallback(async () => {
    setLoading(true);
    setError(null);

    try {
      const response = await fetch('http://localhost:4000/chats');
      if (!response.ok) {
        throw new Error('Failed to fetch chats');
      }
      const responseData: ChatResponse = await response.json();

      console.log(responseData);
      setChats(responseData.data); // Acceder correctamente a 'data'
    } catch (error) {
      setError(error instanceof Error ? error.message : 'Unknown error');
    } finally {
      setLoading(false);
    }
  }, []);

  // Fetch chats when the component mounts
  useEffect(() => {
    fetchChats();
  }, []);

  // Refetch cuando isRefetchNeeded cambia a true
  useEffect(() => {
    if (isNewChatCreated) {
      fetchChats();
      markChatAsNotCreated(); // Limpia el estado después del refetch
    }
  }, [isNewChatCreated, fetchChats, markChatAsNotCreated]);

  return (
    <div className="flex flex-col h-full w-full">
      <h3 className="text-lg py-2">Chats</h3>
      {loading && <p>Loading chats...</p>}
      {error && <p className="text-red-500">{error}</p>}
      <div className="flex flex-col gap-2">
        {chats.length > 0 ? (
          chats.map((chat) => (
            <ChatListItem key={chat.id} href={`/chat/${chat.id}`}>
              {chat.title}
            </ChatListItem>
          ))
        ) : (
          !loading && <p>No chats available.</p>
        )}
      </div>
    </div>
  );
};

export default ChatList;
