"use client"

import * as React from "react"
import { useEffect, useState } from "react"
import { Icons } from "@/components/Icon"
import { Toggle } from "@/ui/toggle"
import { useTheme } from "next-themes"

export function NavbarColorModeToggle() {
  const [mounted, setMounted] = useState(false)
  const [isDark, setIsDark] = useState(false)
  const { theme, setTheme } = useTheme()

  useEffect(() => {
    setMounted(true)
    setIsDark(theme === "dark")
  }, [theme])

  if (!mounted) {
    return null
  }

  const toggleDarkMode = () => {
    setIsDark(!isDark)
    setTheme(theme === "light" ? "dark" : "light")
  }

  return (
    <Toggle onClick={toggleDarkMode} aria-label="toggle mode">
      <Icons.sun className="rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
      <Icons.moon className="absolute rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
    </Toggle>
  )
}
