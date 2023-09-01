import MDXComponents from "@/components/MDXComponents";
import { Separator } from "@/ui/separator";
import { getPost } from "@/api/posts/getPosts";
import { toCapFirstWordString } from "@/lib/utils"


export default async function Page({ params }: { params: { subPost: string } }) {
  const res = await getPost({ slug: params.subPost })
  const markdown = res.data.post_detail.content

  return (
    <div>
      <div className="py-6">
        <div className="space-y-2">
          <h1 className="line-clamp-2 scroll-m-20 text-2xl font-bold tracking-tight">{toCapFirstWordString(res.data.post_detail.title)}</h1>
        </div>
        <Separator className="my-4 md:my-6" />
      </div>
      <MDXComponents source={markdown} />
    </div>
  )
}