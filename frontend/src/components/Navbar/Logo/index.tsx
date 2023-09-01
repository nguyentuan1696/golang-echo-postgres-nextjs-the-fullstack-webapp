import Link from "next/link"

export default function NavbarLogo() {
  return (
    <Link href="/" className="mr-6 flex items-center space-x-2">
      <h1 className="hidden font-bold sm:inline-block">Thích Tài Liệu</h1>
    </Link>
  )
}
