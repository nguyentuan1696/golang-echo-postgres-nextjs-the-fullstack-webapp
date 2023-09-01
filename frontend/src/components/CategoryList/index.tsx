import Link from "next/link"
import { Icons } from "@/components/Icon"
import { Separator } from "@/ui/separator"

export default function CategoryList() {
  return (
    <div>
      <div className="space-y-2">
        <div className="flex items-center justify-between">
          <h2 className="scroll-m-20 text-2xl font-bold tracking-tight md:text-3xl">Danh sách chuyên mục</h2>
          <Link href="/p/n">
            <div className="flex items-center">
              <span className="hidden  pr-2 md:block text-muted-foreground">Xem thêm</span>
              <Icons.moveRight className="text-muted-foreground" />
            </div>
          </Link>
        </div>
        <Separator className="my-4 md:my-6" />
        <div className="">
          <ul className="grid gap-4 md:grid-cols-2 lg:grid-cols-4"></ul>
        </div>
      </div>
    </div>
  )
}
