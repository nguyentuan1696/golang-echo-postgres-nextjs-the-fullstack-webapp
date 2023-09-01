import { Icons } from "@/components/Icon"
import Link from "@/components/Link"
import { NavbarColorModeToggle } from "@/components/Navbar/ColorModeToggle"
import DesktopNav from "@/components/Navbar/DesktopNav"
import NavbarLayout from "@/components/Navbar/Layout"
import MobileNav from "@/components/Navbar/MobileNav"
import NavbarSearch from "@/components/Navbar/Search"
import { buttonVariants } from "@/ui/button"
import { cn } from "@/lib/utils"

function SocialIcon() {
  return (
    <>
      <Link href="https://www.facebook.com/ThichTiengAnhFP">
        <div
          className={cn(
            buttonVariants({
              size: "sm",
              variant: "ghost",
            }),
            "h-10 px-3"
          )}
        >
          <Icons.facebook className="h-5 w-5" />
          <span className="sr-only">GitHub</span>
        </div>
      </Link>
    </>
  )
}

export default function Navbar() {
  return (
    <NavbarLayout>
      <div className={cn("container flex h-14 items-center")}>
        <MobileNav />
        <DesktopNav />
        <div className={cn("flex flex-1 items-center justify-between space-x-2 md:justify-end")}>
          <div className="w-full flex-1 md:w-auto md:flex-none">
            <NavbarSearch />
          </div>
          <div className="flex items-center">
            <SocialIcon />
            <NavbarColorModeToggle />
          </div>
        </div>
      </div>
    </NavbarLayout>
  )
}
