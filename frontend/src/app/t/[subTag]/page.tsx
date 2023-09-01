import Link from "@/components/Link"
import { ProductCard } from "@/components/Products"
import { getTag } from "@/api/tags/getTags"

export default async function Page({ params }: { params: { subTag: string } }) {
  const res = await getTag({ slug: params.subTag })

  return (
    <div className="grid grid-cols-2 gap-4 md:grid-cols-4">
      {Array.from(res.data).map((d, i) => (
        <div key={i}>
          <Link href={`p/${d.slug}`}>
            <ProductCard title={d.title} description="" />
          </Link>
        </div>
      ))}
    </div>
  )
}
