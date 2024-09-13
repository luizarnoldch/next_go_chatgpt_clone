// stores/useActionStore.ts
import { create } from "zustand"

interface ActionState {
  sayHello: () => void
}

export const useActionStore = create<ActionState>((set) => ({
  sayHello: () => {
    console.log("Hello World")
  },
}))
