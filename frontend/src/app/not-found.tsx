import { cn } from "@/lib/utils"

export default function NotFound() {
  return (
    <div className="container">
      <div className="p-6">
        <h2 className={cn("font-heading mt-12 scroll-m-20 pb-2 text-2xl font-semibold tracking-tight first:mt-0")}>👻 Không tìm thấy trang</h2>
        <p className="flex flex-col">
          <span>Chúng tôi không thể tìm thấy những gì bạn đang tìm kiếm.</span>
          <span>Vui lòng liên hệ với chủ sở hữu của trang web đã liên kết bạn với bản gốc URL và cho họ biết liên kết của họ bị hỏng.</span>
        </p>
      </div>
    </div>
  )
}
