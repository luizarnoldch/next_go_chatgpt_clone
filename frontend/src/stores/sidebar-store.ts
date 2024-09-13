import { create } from "zustand";

interface SidebarState {
  isMobileSidebarOpen: boolean;
  isDesktopSidebarOpen: boolean;
  openMobileSidebar: () => void;
  closeMobileSidebar: () => void;
  openDesktopSidebar: () => void;
  closeDesktopSidebar: () => void;
}

export const useSidebarStore = create<SidebarState>()((set) => ({
  isMobileSidebarOpen: false,
  isDesktopSidebarOpen: true,
  openMobileSidebar: () => set({ isMobileSidebarOpen: true }),
  closeMobileSidebar: () => set({ isMobileSidebarOpen: false }),
  openDesktopSidebar: () => set({ isDesktopSidebarOpen: true }),
  closeDesktopSidebar: () => set({ isDesktopSidebarOpen: false }),
}));
