import { toCapFirstLetter, toCapFirstWordString } from "@/lib/utils";
import { Separator } from "@/ui/separator";
import Balancer from "@/components/Balancer";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div>
        <div className="space-y-2">
          <h1 className="scroll-m-20 text-4xl font-bold tracking-tight">{toCapFirstLetter("tài liệu mới cập nhật")}</h1>
          <p className="text-lg text-muted-foreground">
            <Balancer title="bai viet moi" />
          </p>
        </div>
        <Separator className="my-4 md:my-6" />
         <div className="grid grid-cols-1 md:grid-cols-[25%_minmax(50%,_1fr)_25%]">
      <div className="hidden md:inline">sidebar</div>
      <div className="">{children}</div>
      <div className="order-last mt-10 md:mt-0 md:pl-6">
        <h2 className="font-heading mt-12 scroll-m-20 border-b pb-2 text-xl font-semibold tracking-tight first:mt-0">
          {toCapFirstWordString("Bài viết cùng chuyên mục")}
        </h2>
        <ul className="pt-3">
         <p>tuan</p>
        </ul>
      </div>
    </div>
    </div>
   
  )
}
