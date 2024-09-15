// stores/useActionStore.ts
import { create } from "zustand";

interface NewChatState {
  isNewChatCreated: boolean;
  markChatAsCreated: () => void;
  markChatAsNotCreated: () => void;
}

export const useNewChatStore = create<NewChatState>((set) => ({
  isNewChatCreated: true,
  markChatAsCreated: () => set({ isNewChatCreated: true }),
  markChatAsNotCreated: () => set({ isNewChatCreated: false }),
}));
