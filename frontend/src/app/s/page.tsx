"use client"

import { useEffect, useState } from "react"
import Link from "next/link"
import { useRouter, useSearchParams } from "next/navigation"
import { ProductCard } from "@/components/Products"
import { Button } from "@/ui/button"
import { Input } from "@/ui/input"
import { Separator } from "@/ui/separator"
import { getSearch } from "@/api/search/getSearch"
import { decodeQueryParam } from "@/lib/utils"
import { Search } from "lucide-react"

export default function Page() {
  const searchParams = useSearchParams()
  const search = searchParams.get("search")
  const [value, setValue] = useState({} as any)
  const [valueChange, setValueChange] = useState("")
  const router = useRouter()

  useEffect(() => {
    async function fetchData() {
      const res = await getSearch(decodeQueryParam(search || ""))
      setValue(res.data)
    }
    fetchData()
  }, [search])

  const handleChange = (e: any) => {
    setValueChange(e.target.value)
  }

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

  return (
    <>
      <div className="flex items-center justify-center">
        <div className="flex basis-full	space-x-2 md:basis-2/4">
          <Input type="text" placeholder="Nhập tên tài liệu..." onChangeCapture={handleChange} onKeyDownCapture={down} />
          <Button aria-label="search home page" onClickCapture={handleValueChange}>
            <Search className="h-5 w-5" />
          </Button>
        </div>
      </div>

      <Separator className="my-4 md:my-6" />
      <div className="grid grid-cols-1 md:grid-cols-[25%_minmax(50%,_1fr)_25%]">
        <div className="hidden md:inline"></div>
        <div>
          {value.total != null && value.total > 0 ? (
            <div className="space-y-6">
              <p className="text-muted-foreground">{`Khoảng ${value.total} kết quả tìm kiếm từ khóa: ${search}`}</p>
              <ul className="flex-col space-y-4">
                {Array.from(value.data).map((d: any, i) => (
                  <li key={i}>
                    <Link href={`p/${d.slug}`}>
                      <ProductCard title={d.title} id={i} />
                    </Link>
                  </li>
                ))}
              </ul>
            </div>
          ) : (
            <p className="text-muted-foreground">{`Không tìm thấy kết quả theo từ khóa: ${search}`}</p>
          )}
        </div>
        <div className="order-last mt-10 md:mt-0 md:pl-6"></div>
      </div>
    </>
  )
}
