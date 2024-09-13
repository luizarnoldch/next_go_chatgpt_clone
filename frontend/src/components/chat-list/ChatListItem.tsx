// Lib
import Link from 'next/link';

// Components
import { Button } from '@/components/ui/button';
import { ReactNode } from 'react';

type ChatListItemProps = {
  href: string;
  children: ReactNode;
};

const ChatListItem: React.FC<ChatListItemProps> = ({ href, children }) => {
  return (
    <Button className="w-full relative overflow-x-hidden" variant="secondary" size="sm">
      <Link href={href} className="w-full flex justify-start items-center gap-2">
        {children}
      </Link>
    </Button>
  );
};

export default ChatListItem;
