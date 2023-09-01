"use client"

// Error components must be Client Components
import { useEffect } from "react"
import { cn } from "@/lib/utils"

export default function Error({ error, reset }: { error: Error; reset: () => void }) {
  useEffect(() => {
    // Log the error to an error reporting service
    console.error(error)
  }, [error])

  return (
    <div className="container">
      <div className="p-6">
        <h2 className={cn("font-heading mt-12 scroll-m-20 pb-2 text-2xl font-semibold tracking-tight first:mt-0")}>Something went wrong!</h2>
        {/* <button
        onClick={
          // Attempt to recover by trying to re-render the segment
          () => reset()
        }
      >
        Try again
      </button> */}
      </div>
    </div>
  )
}
