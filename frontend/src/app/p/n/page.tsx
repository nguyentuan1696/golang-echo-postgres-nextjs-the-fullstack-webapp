import Link from "next/link"
import { ProductCard } from "@/components/Products"
import { getLatestPosts } from "@/api/posts/getPosts"

export default async function Page() {
  const res = await getLatestPosts({ pageSize: 20, pageIndex: 1 })
  return (
    <div>
      <ul className="flex-col space-y-4">
        {Array.from(res.data).map((d, i) => (
          <li key={i}>
            <Link href={`p/${d.slug}`}>
              <ProductCard title={d.title} id={i} />
            </Link>
          </li>
        ))}
      </ul>
    </div>
  )
}
