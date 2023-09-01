import Balancer from "@/components/Balancer"
import { Separator } from "@/ui/separator"
import { getTagInfo } from "@/api/tags/getTags"
import type { TagInfoSlugProps } from "@/types/TagInfoSlugProps"
import { toCapFirstLetter } from "@/lib/utils";

export default async function Layout({ children, params }: { children: React.ReactNode; params: TagInfoSlugProps }) {
  const res = await getTagInfo(params.subTag)

  return (
    <div className="container">
      <div className="py-6">
        <div className="space-y-2">
          <h1 className="scroll-m-20 text-4xl font-bold tracking-tight">🐳 Chủ đề: {toCapFirstLetter(res.data.title)}</h1>
          <p className="text-muted-foreground text-lg">
            <Balancer title={res.data.description} />
          </p>
        </div>
        <Separator className="my-4 md:my-6" />
        <div>{children}</div>
      </div>
    </div>
  )
}
