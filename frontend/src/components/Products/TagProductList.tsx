import Link from "next/link"
import { Icons } from "@/components/Icon"
import { Separator } from "@/ui/separator"
import type { Tag } from "@/types/Tag"

export default function TagProductList(data: Tag) {
  return (
    <div className="space-y-2">
      <div className="flex items-center justify-between">
        <h2 className="scroll-m-20 text-2xl font-bold tracking-tight md:text-3xl">Chủ đề nổi bật</h2>
        <Link href="/t">
          <div className="flex items-center">
            <span className="hidden  pr-2 md:block text-muted-foreground">Xem thêm</span>
            <Icons.moveRight className="text-muted-foreground" />
          </div>
        </Link>
      </div>
      <Separator className="my-4 md:my-6" />
      <div>
        <ul className="flex flex-wrap justify-start">
          {Array.from(
            data.data.map((d, i) => (
              <li  key={i} className="mb-2 mr-2 rounded-lg bg-muted  p-2 ">
                <Link href={`t/${d.slug}`}>
                  <span>{d.title.toLowerCase()}</span>
                </Link>
              </li>
            ))
          )}
        </ul>
      </div>
    </div>
  )
}
