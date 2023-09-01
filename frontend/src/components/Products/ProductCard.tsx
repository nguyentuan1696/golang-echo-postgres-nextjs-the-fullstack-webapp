import { toCapFirstWordString } from "@/lib/utils"
import type { Props } from "@/types/Products/ProductCard"

export default function ProductCard({ title, description, content, footer, className, id, ...props }: Props) {
  return (
    <div className="border-b border-dotted pb-1">
      <div className="flex items-center justify-between">
        <div className="flex items-center">
          <p className="mr-2 flex h-4 w-4 items-center justify-center rounded-full text-muted-foreground">
            <span>{id == undefined ? "" : (id += 1)}.</span>
          </p>
          <p className="line-clamp-1 text-lg font-medium"> {toCapFirstWordString(title)}</p>
        </div>
      </div>
    </div>
  )
}
