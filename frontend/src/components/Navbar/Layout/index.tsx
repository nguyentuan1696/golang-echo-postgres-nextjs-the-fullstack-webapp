import { ReactNode } from "react"
import { cn } from "@/lib/utils"

export default function NavbarLayout({ children }: { children: ReactNode }) {
  return (
    <header className={cn("supports-backdrop-blur:bg-background/60 sticky top-0 z-40 w-full border-b bg-background/95 backdrop-blur")}>
      {children}
    </header>
  )
}
