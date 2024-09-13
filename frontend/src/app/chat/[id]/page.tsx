"use client";

import React, { useEffect, useState, useCallback } from 'react';
import Topvar from '@/components/topvar/Topvar';
import ChatBox from '@/components/chat-box/ChatBox';

type Props = { params: { id: string } };

const ChatWithID = ({ params }: Props) => {
  // const chat = useChatStore((state) => state.getChatById(params.id));
  // const addSystemMessage = useChatStore((state) => state.addSystemMessage);
  // const [hasAddedSystemMessage, setHasAddedSystemMessage] = useState(false);

  // const isLastMessageFromUser = useCallback(() => {
  //   const lastMessage = chat?.messages[chat.messages.length - 1];
  //   return lastMessage?.user === 'Person';
  // }, [chat]);

  // useEffect(() => {
  //   if (chat && isLastMessageFromUser() && !hasAddedSystemMessage) {
  //     const responseMessage = "This is an automatic system response."; // Placeholder response
  //     addSystemMessage(responseMessage, params.id);
  //     setHasAddedSystemMessage(true);
  //   }
  // }, [chat, isLastMessageFromUser, hasAddedSystemMessage, addSystemMessage, params.id]);

  // if (!chat) {
  //   return <div>No chat found with ID {params.id}</div>;
  // }

  const sendMessage = (input: string) => {
    // Placeholder for future message sending logic
  };

  return (
    <section className='bg-secondary h-screen w-full p-4 flex-1 flex flex-col'>
      <Topvar />
      <div className='flex-1 m-6'>
        <aside className='max-w-7xl mx-auto p-4'>
          <div className='space-y-4'>
            {/* {chat.messages.map((message, index) => (
              <div
                key={index}
                className={`flex ${message.user === 'Person' ? 'justify-end' : 'justify-start'}`}
              >
                <div
                  className={`${message.user === 'Person' ? 'bg-blue-500 text-white' : 'bg-gray-300 text-black'} p-3 rounded-lg max-w-md`}
                >
                  <p className='font-bold'>{message.user}</p>
                  <p>{message.message}</p>
                </div>
              </div>
            ))} */}
            {params.id}
          </div>
        </aside>
      </div>
      <ChatBox handleButton={sendMessage} />
    </section>
  );
};

export default ChatWithID;
