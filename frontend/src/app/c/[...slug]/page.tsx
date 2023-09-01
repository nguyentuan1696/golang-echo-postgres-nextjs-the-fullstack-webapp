import Link from "next/link"
import { ProductCard } from "@/components/Products"
import { getPostByCategory } from "@/api/categories/getCatgories"

export default async function Page({ params }: any) {
  const res = await getPostByCategory({
    categorySlug: params.slug[0],
    subCategorySlug: (params.slug[1] && params.slug[1]) || "",
  })

  return (
    <div>
      <ul className="flex-col space-y-6">
        {Array.from(res.data.list_post).map((d, i) => (
          <li key={i}>
            <Link href={`p/${d.slug_post}`}>
              <ProductCard title={d.title_post} id={i} />
            </Link>
          </li>
        ))}
      </ul>
    </div>
  )
}
