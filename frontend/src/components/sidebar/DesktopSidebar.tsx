"use client"

// Lib
import Link from "next/link";
import { cn } from "@/lib/utils";

// Components
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger
} from "@/components/ui/tooltip";
import { Button } from "@/components/ui/button";

// Stores
import { useSidebarStore } from "@/stores/sidebar-store"
// Stores
import { useActionStore } from "@/stores/create-new-chat"

// Icons
import { MenuIcon, SquarePenIcon } from "lucide-react";

import ChatGPTIcon from "../assets/icons/ChatGPTIcon";
import ChatList from "../chat-list/ChatList";


type DesktopSidebarProps = {}

const DesktopSidebar: React.FC<DesktopSidebarProps> = () => {
  const { isDesktopSidebarOpen, closeDesktopSidebar } = useSidebarStore((state) => ({
    isDesktopSidebarOpen: state.isDesktopSidebarOpen,
    closeDesktopSidebar: state.closeDesktopSidebar,
  }))

  const { sayHello } = useActionStore((state) => ({
    sayHello: state.sayHello
  }))

  return (
    <div
      className={cn(
        "h-full transition-[width] duration-500 ease-in-out overflow-hidden",
        isDesktopSidebarOpen ? "w-[250px]" : "w-0"
      )}
    >
      <div className={cn("p-4 ", !isDesktopSidebarOpen && 'opacity-0 pointer-events-none')}>
        <div className="flex items-center justify-between w-full mb-4">
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Button variant="ghost" size="icon" onClick={closeDesktopSidebar}>
                  <MenuIcon className="h-6 w-6 cursor-pointer" />
                </Button>
              </TooltipTrigger>
              <TooltipContent>
                <p>Close the Sidebar</p>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>

          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Button
                  variant="ghost"
                  size="icon"
                  onClick={sayHello}
                >
                  <SquarePenIcon className="h-6 w-6 cursor-pointer" />
                </Button>
              </TooltipTrigger>
              <TooltipContent>
                <p>New Chat</p>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>
        </div>

        <div className="grid gap-4 py-4 group/item">
          <Button className="w-full relative">
            <Link href="/" className="w-full flex justify-start items-center gap-2">
              <ChatGPTIcon className="size-6 fill-primary-foreground" />
              <p>ChatGPT</p>
              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <SquarePenIcon className="absolute right-4 size-4 hidden group-hover/item:block" />
                  </TooltipTrigger>
                  <TooltipContent>
                    <p>New Chat</p>
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>
            </Link>
          </Button>
        </div>

        <ChatList />
      </div>
    </div>
  )
}

export default DesktopSidebar
