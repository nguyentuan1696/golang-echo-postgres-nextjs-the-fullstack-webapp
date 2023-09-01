import Link from "next/link"
import { ArrowRight } from "lucide-react"

export default function ProductList() {
  return (
    <div>
      <div className="space-y-2">
        <div className="flex justify-between	">
          <h2 className="scroll-m-20 text-2xl md:text-3xl font-bold tracking-tight">Tai lieu toeic</h2>
          <Link href="c/tieng-anh/toeic" className="flex items-center rounded-[0.5rem] text-sm font-medium ">
            <span>Xem thêm</span>
            <ArrowRight className="ml-1 h-4 w-4" />
          </Link>
        </div>
      </div>
      {/* `  <Separator className="my-4 md:my-6" />` */}
    </div>
  )
}
