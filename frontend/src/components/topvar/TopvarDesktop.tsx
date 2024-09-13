"use client"

// Lib
import { useTheme } from "next-themes"

// Components
import { Button } from "@/components/ui/button"

// Store
import { useSidebarStore } from "@/stores/sidebar-store"

// Icons
import { MenuIcon, Moon, SquarePenIcon, Sun } from "lucide-react"

type TopvarDesktopProps = {}

const TopvarDesktop: React.FC<TopvarDesktopProps> = ({ }) => {
  const { setTheme, resolvedTheme } = useTheme()

  const { isDesktopSidebarOpen, openDesktopSidebar } = useSidebarStore((state) => ({
    isDesktopSidebarOpen: state.isDesktopSidebarOpen,
    openDesktopSidebar: state.openDesktopSidebar,
  }))

  return (
    <div className="w-full justify-between items-center flex">
      {
        !isDesktopSidebarOpen && <div className="flex gap-2">
          <Button variant="outline" size="icon" onClick={openDesktopSidebar}>
            <MenuIcon />
          </Button>
          <Button variant="outline" size='icon'>
            <SquarePenIcon />
          </Button>
        </div>
      }
      <p>ChatGPT</p>
      <Button
        variant="outline"
        size="icon"
        onClick={() => setTheme(resolvedTheme === "light" ? "dark" : "light")}
      >
        <Sun className="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0 text-primary" />
        <Moon className="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
        <span className="sr-only">Toggle theme</span>
      </Button>
    </div>
  )
}

export default TopvarDesktop
