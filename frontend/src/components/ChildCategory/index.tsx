import Link from "next/link"
import { toCapFirstWordString } from "@/lib/utils"
import type { ChildCategoryListView } from "@/types/ChildCategoryListView"

export default function ChildCategoryList(data: ChildCategoryListView, parentCate: any) {
  return (
    <div>
      <h2 className="font-heading mt-12 scroll-m-20 border-b pb-2 text-xl font-semibold tracking-tight first:mt-0">
        {toCapFirstWordString("chuyên mục cùng chủ đề")}
      </h2>
      <ul className="pt-3">
        {Array.from(data.data)
          .filter((p) => p.slug != parentCate)
          .map((d, i) => (
            <Link href={`c/${parentCate}/${d.slug}`}>
              <li key={i} className="mt-2">
                <p>{toCapFirstWordString(d.title)}</p>
              </li>
            </Link>
          ))}
      </ul>
    </div>
  )
}
