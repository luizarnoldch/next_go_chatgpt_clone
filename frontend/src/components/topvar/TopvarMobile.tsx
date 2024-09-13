'use client'

// Lib
import { useTheme } from "next-themes"

// Components
import { Button } from "@/components/ui/button"

// Store
import { useSidebarStore } from "@/stores/sidebar-store"

// Icon
import { MenuIcon, SquarePenIcon, Moon, Sun } from "lucide-react"

type Props = {}

const TopvarMobile = (props: Props) => {
  const { setTheme, resolvedTheme } = useTheme()
  const openMobileSidebar = useSidebarStore((state) => state.openMobileSidebar)
  return (
    <div className="w-full justify-between items-center flex">
      <Button variant="outline" size='icon' onClick={openMobileSidebar}>
        <MenuIcon />
      </Button>
      <p>ChatGPT</p>
      <div className="flex gap-2">
        <Button variant="outline" size='icon'>
          <SquarePenIcon />
        </Button>
        <Button
          variant="outline"
          size="icon"
          onClick={() => setTheme(resolvedTheme === "light" ? "dark" : "light")}
        >
          <Sun className="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
          <Moon className="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
          <span className="sr-only">Toggle theme</span>
        </Button>
      </div>


    </div>
  )
}

export default TopvarMobile