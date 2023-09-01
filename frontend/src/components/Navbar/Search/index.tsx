"use client"

import * as React from "react"
import { useRouter } from "next/navigation"
import { Button } from "@/ui/button"
import { CommandDialog, CommandGroup, CommandInput, CommandItem, CommandList } from "@/ui/command"
import { getSearch } from "@/api/search/getSearch"
import { cn } from "@/lib/utils"
import useDebounce from "@/hooks/useDebounce"
import type { Search } from "@/types/Search"
import { DialogProps } from "@radix-ui/react-alert-dialog"

export default function NavbarSearch({ ...props }: DialogProps) {
  const [open, setOpen] = React.useState(false)
  const [search, setSearch] = React.useState("")
  const [results, setResults] = React.useState({} as Search)
  const [isSearching, setIsSearching] = React.useState(false)
  const debouncedSearch = useDebounce(search, 1000)
  const router = useRouter()

  React.useEffect(() => {
    const searchPost = async () => {
      let results = {} as Search
      setIsSearching(true)
      if (debouncedSearch) {
        const res = await getSearch(debouncedSearch)
        results = res || {}
      }
      setIsSearching(false)
      setResults(results)
    }
    searchPost()
  }, [debouncedSearch])

  React.useEffect(() => {
    const down = (e: KeyboardEvent) => {
      if (e.key === "k" && (e.metaKey || e.ctrlKey)) {
        e.preventDefault()
        setOpen((open) => !open)
      }
    }

    document.addEventListener("keydown", down)
    return () => document.removeEventListener("keydown", down)
  }, [])

  const runCommand = React.useCallback((command: () => unknown) => {
    setOpen(false)
    command()
  }, [])

  return (
    <>
     <Button
        variant="outline"
        className={cn(
          "relative w-full justify-start text-sm text-muted-foreground sm:pr-12 md:w-40 lg:w-64"
        )}
        onClick={() => setOpen(true)}
        {...props}
      >
        <span className="hidden lg:inline-flex">Tìm kiếm tài liệu...</span>
        <span className="inline-flex lg:hidden">Tìm kiếm ...</span>
        <kbd className="pointer-events-none absolute right-1.5 top-1.5 hidden h-5 select-none items-center gap-1 rounded border bg-muted px-1.5 font-mono text-[10px] font-medium opacity-100 sm:flex">
          <span className="text-xs">⌘</span>K
        </kbd>
      </Button>
      <CommandDialog open={open} onOpenChange={setOpen}>
        <CommandInput placeholder="Tìm kiếm tài liệu..." onChangeCapture={(e) => setSearch((e.target as HTMLInputElement).value)} />
        <CommandList>
          {/* <CommandEmpty>No results found.</CommandEmpty> */}
          <CommandGroup heading="links">
            {isSearching || Object.keys(results).length === 0
              ? null
              : results.data.data.map((d, i) => (
                  <CommandItem key={i} onSelect={() => runCommand(() => router.push(d.slug))}>
                    <span className="line-clamp-1">{d.title}</span>
                  </CommandItem>
                ))}
          </CommandGroup>
        </CommandList>
      </CommandDialog>
    </>
  )
}
