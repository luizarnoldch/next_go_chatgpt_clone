"use client";

import React, { useState, useRef } from 'react';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { SendIcon } from 'lucide-react';

type ChatBoxProps = {
  handleButton: (input: string) => void;
};

const ChatBox: React.FC<ChatBoxProps> = ({ handleButton }) => {
  const [input, setInput] = useState('');
  const textAreaRef = useRef<HTMLTextAreaElement>(null);

  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInput(e.target.value);
    adjustHeight();
  };

  const adjustHeight = () => {
    if (textAreaRef.current) {
      textAreaRef.current.style.height = 'auto';
      textAreaRef.current.style.height = `${Math.min(textAreaRef.current.scrollHeight, 600)}px`;
    }
  };

  return (
    <div className='border-t p-4 flex bg-background gap-2 justify-center items-center max-w-7xl mx-auto rounded-lg w-full'>
      <Textarea
        ref={textAreaRef}
        placeholder='Type a message...'
        value={input}
        onChange={handleChange}
        rows={1}
        className='w-full resize-none p-4 border rounded-md'
        style={{
          maxHeight: '600px',
          minHeight: '60px',
        }}
      />
      <Button onClick={() => { handleButton(input); setInput(''); }} disabled={input.length < 1}>
        <SendIcon />
      </Button>
    </div>
  );
};

export default ChatBox;
