import type { ReactNode } from "react"
import NextLink from "next/link"
import { isInternalUrlUtils } from "@/lib/utils"

export default function Link({ href, children, className = "" }: { href: string; children: ReactNode; className?: string }) {
  const isExternal = !isInternalUrlUtils(href)

  if (isExternal) {
    return (
      <a className={className} target="_blank" rel="noopener noreferrer" href={href}>
        {children}
      </a>
    )
  }

  return (
    <NextLink href={href} className={className}>
      {children}
    </NextLink>
  )
}
