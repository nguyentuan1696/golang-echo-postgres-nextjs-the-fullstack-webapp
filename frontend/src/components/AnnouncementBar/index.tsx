"use client"

import { useState } from "react"
import { X } from "lucide-react"

export default function AnnouncementBar() {
  const [show, setShow] = useState(true)

  function setHideBanner() {
    setShow(false)
  }

  return (
    <div>
      {show ? (
        <div className="border-b p-2">
          <div className="flex items-center">
            <div className="shrink grow basis-auto">
              <a href="https://thichtienganh.com" target="_blank" rel="noopener noreferrer">
                <p className="text-center text-sm font-bold">
                  🚀 Trải nghiệm học Tiếng Anh hoàn toàn <span className="underline">miễn phí</span>
                </p>
              </a>
            </div>
            <div className="cursor-pointer" onClickCapture={() => setHideBanner()}>
              <X className="h-4 w-4" />
            </div>
          </div>
        </div>
      ) : null}
    </div>
  )
}
