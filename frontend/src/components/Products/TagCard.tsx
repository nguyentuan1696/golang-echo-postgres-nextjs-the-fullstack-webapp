import { cn } from "@/lib/utils"

export default function TagCard({ content, className, ...props }: { content: string; className?: string }) {
  return (
    <p className={cn("border p-3", className)} {...props}>
      {content}
    </p>
  )
}
