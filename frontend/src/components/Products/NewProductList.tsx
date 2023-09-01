import Link from "next/link"
import { Icons } from "@/components/Icon"
import { Separator } from "@/ui/separator"
import { toCapFirstWordString } from "@/lib/utils"
import type { LatestPosts } from "@/types/LatestPosts"

export default function NewProductList(data: LatestPosts) {
  return (
    <div>
      <div className="space-y-2">
        <div className="flex items-center justify-between">
          <h2 className="scroll-m-20 text-2xl font-bold tracking-tight md:text-3xl">Tài liệu mới cập nhật</h2>
          <Link href="/p/n">
            <div className="flex items-center">
              <span className="hidden pr-2 md:block text-muted-foreground">Xem thêm</span>
              <Icons.moveRight className="text-muted-foreground" />
            </div>
          </Link>
        </div>
        <Separator className="my-4 md:my-6" />
        <div className="">
          <ul className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
            {Array.from(data.data).map((d, i) => (
              <li key={i} className="rounded-lg border p-4">
                <Link href={`p/${d.slug}`}>
                  <p className="line-clamp-2 text-lg font-bold		">{toCapFirstWordString(d.title)}</p>
                </Link>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  )
}
