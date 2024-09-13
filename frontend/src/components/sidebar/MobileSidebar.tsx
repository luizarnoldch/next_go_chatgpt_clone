"use client"

// Lib
import Link from "next/link"

// Components
import { Button } from "@/components/ui/button"
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetFooter,
} from "@/components/ui/sheet"

import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"

// Stores
import { useSidebarStore } from "@/stores/sidebar-store"
import { useActionStore } from "@/stores/create-new-chat"

// Icons
import ChatGPTIcon from "../assets/icons/ChatGPTIcon"
import { SquarePenIcon } from "lucide-react"


type MobileSidebarProps = {}

const MobileSidebar: React.FC<MobileSidebarProps> = ({ }) => {
  const { isMobileSidebarOpen, closeMobileSidebar } = useSidebarStore((state) => ({
    isMobileSidebarOpen: state.isMobileSidebarOpen,
    closeMobileSidebar: state.closeMobileSidebar,
  }))

  const { sayHello } = useActionStore((state) => ({
    sayHello: state.sayHello
  }))

  return (
    <Sheet open={isMobileSidebarOpen} onOpenChange={closeMobileSidebar}>
      <SheetContent side='left'>
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
        <SheetFooter>
          <SheetClose asChild>
            <Button type="submit" >Save changes</Button>
          </SheetClose>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  )
}

export default MobileSidebar