"use client"

import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import { Button } from "@/ui/button"
import { Input } from "@/ui/input"
import { getSearchKeywords } from "@/api/search/getSearch"
import { randomValueArray } from "@/lib/utils"
import { Search } from "lucide-react"

export default function SearchBarMain() {
  const router = useRouter()
  const [value, setValue] = useState([])
  const [valueChange, setValueChange] = useState("")

  useEffect(() => {
    async function fetchData() {
      const res: any = await getSearchKeywords()
      setValue(res.data)
    }
    fetchData()
  }, [])

  const down = (e: any) => {
    if (e.target.value === "") {
      return
    }

    if (e.key === "Enter") {
      e.preventDefault()
      router.push(`/s?search=${valueChange}`)
    }
  }

  const handleValueChange = () => {
    if (valueChange.length === 0) {
      return
    }
    router.push(`/s?search=${valueChange}`)
  }

  const handleChange = (e: any) => {
    setValueChange(e.target.value)
  }

  return (
    <div className="flex items-center justify-between	space-x-2 ">
      <Input type="text" placeholder={randomValueArray(value)} onChangeCapture={handleChange} onKeyDownCapture={down} />
      <Button aria-label="search home page" onClickCapture={handleValueChange}>
        <Search className="h-5 w-5" />
      </Button>
    </div>
  )
}
