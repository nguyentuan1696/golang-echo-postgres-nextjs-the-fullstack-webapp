// RootLayout
declare module "@/types/RootLayout" {
  import type { ReactElement, ReactNode } from "react"

  export interface Props {
    children: ReactNode
  }

  export default function RootLayout(props: Props): JSX.Element
}
3

declare module "@/types/Products/ProductCard" {
  export interface Props {
    id?: number
    title: string
    description?: string
    content?: string
    footer?: string
    className?: string
  }

  export default function ProductCard(props: Props): JSX.Element
}
