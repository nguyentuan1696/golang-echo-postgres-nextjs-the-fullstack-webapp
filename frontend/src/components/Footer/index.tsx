"use client"

import { usePathname } from "next/navigation"
import Link from "@/components/Link"

const data = [
  {
    title: "Chuyên mục nổi bật",
    child: [
      {
        label: "Tài liệu TOEIC",
        link: "/c/tieng-anh/toeic",
        isHidden: false,
      },
      {
        label: "Tài liệu TOEIC",
        link: "/c/tieng-anh/toeic",
        isHidden: false,
      },
      {
        label: "Tài liệu TOEIC",
        link: "/c/tieng-anh/toeic",
        isHidden: false,
      },
    ],
  },
  {
    title: "Liên hệ",
    child: [
      {
        label: "Liên hệ với chúng tôi",
        link: "/c/tieng-anh/toeic",
        isHidden: false,
      },
      {
        label: "Liên hệ quảng cáo",
        link: "/c/tieng-anh/toeic",
        isHidden: false,
      },
      {
        label: "Gửi yêu cầu hỏi đáp",
        link: "/c/tieng-anh/toeic",
        isHidden: false,
      },
    ],
  },
  {
    title: "Kết nối",
    child: [
      {
        label: "Facebook",
        link: "https://www.facebook.com/ThichTiengAnhFP",
        isHidden: false,
      },
      {
        label: "Discord",
        link: "/c/tieng-anh/toeic",
        isHidden: false,
      },
      {
        label: "Tiktok",
        link: "/c/tieng-anh/toeic",
        isHidden: false,
      },
    ],
  },
  {
    title: "Liên Kết ",
    child: [
      {
        label: "Thích Tiếng Anh",
        link: "https://thichtienganh.com",
        isHidden: true,
      },
      {
        label: "Thích Văn Học",
        link: "https://thichvanhoc.com/",
        isHidden: true,
      },
    ],
  },
]

export default function Footer(): JSX.Element {
  const pathname = usePathname()
  const isHomePagePath = pathname === "/"
  return (
    <footer className="border-t py-6">
      <div className="container flex flex-col justify-between md:flex-row [&>*:not(:first-child)]:mt-10 md:[&>*:not(:first-child)]:mt-0">
        {data.map((d, i) => (
          <div key={i}>
            <h3 className="font-medium">{d.title}</h3>
            <ul>
              {d.child.map((c, i) => (
                <li key={i}>
                  <p className="mt-3 md:mt-4">
                    {!isHomePagePath && c.isHidden ? (
                      <span className="text-muted-foreground">{c.label}</span>
                    ) : (
                      <Link href={c.link} className="text-muted-foreground">
                        {c.label}
                      </Link>
                    )}
                  </p>
                </li>
              ))}
            </ul>
          </div>
        ))}
      </div>
      <div className="mt-14 text-center text-muted-foreground">
        <p>Copyright &#169; 2021 &#9473; {new Date().getFullYear()} Thích Tài Liệu, vận hành và phát triển bởi Tuan Nguyen </p>
      </div>
    </footer>
  )
}
