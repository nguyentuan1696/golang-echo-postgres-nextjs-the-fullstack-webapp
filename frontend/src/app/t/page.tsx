import { Metadata } from "next"
import Link from "@/components/Link"
import { Separator } from "@/ui/separator"
import { getTags } from "@/api/tags/getTags"
import { toCapFirstLetter } from "@/lib/utils"

export const metadata: Metadata = {
  title: "Tổng hợp chủ đề",
}

export default async function Page() {
  const res = await getTags()
  return (
    <div>
      <div className="space-y-2">
        <h1 className="scroll-m-20 text-4xl font-bold tracking-tight">{toCapFirstLetter("chủ đề nổi bật")}</h1>
        <p className="line-clamp-2 text-lg text-muted-foreground">
          Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making
          it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin
          words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable
          source.
        </p>
      </div>
      <Separator className="my-4 md:my-6" />
      <div>
        <ul className="flex flex-wrap justify-start">
          {Array.from(res.data).map((d, i) => (
            <li key={i} className="mb-2 mr-2 rounded-lg bg-muted  p-2 ">
              <Link href={`t/${d.slug}`}>
                <span>{d.title.toLocaleLowerCase()}</span>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </div>
  )
}
