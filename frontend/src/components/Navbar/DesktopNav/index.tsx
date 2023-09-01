"use client";

import * as React from "react";
import { ReactNode } from "react";
import Link from "next/link";
import { Icons } from "@/components/Icon";
import { NavigationMenu, NavigationMenuContent, NavigationMenuItem, NavigationMenuLink, NavigationMenuList, NavigationMenuTrigger } from "@/ui/navigation-menu";
import { cn } from "@/lib/utils";
import { navBarLeftItems } from "@/config/navbar"

interface ListItemProps {
  className?: string
  title: string
  children: ReactNode
  href: string
}

export default function DesktopNav() {
  return (
    <div>
      <NavigationMenu >
        <Link href="/" className="mr-6 flex items-center space-x-2">
           <Icons.logo className="h-6 w-6" />
          <span className="hidden font-bold sm:inline-block">Thích Tài Liệu</span>
        </Link>
        <NavigationMenuList className="hidden md:flex">
          <NavigationMenuItem>
            <NavigationMenuTrigger>Tiếng Anh</NavigationMenuTrigger>
            <NavigationMenuContent>
              <ul className="grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-2 lg:w-[600px] ">
                {navBarLeftItems.map((component) => (
                  <ListItem key={component.title} title={component.title} href={component.href}>
                    {component.description}
                  </ListItem>
                ))}
              </ul>
            </NavigationMenuContent>
          </NavigationMenuItem>
        </NavigationMenuList>
      </NavigationMenu>
    </div>
  )
}

const ListItem = React.forwardRef<React.ElementRef<"a">, ListItemProps>(({ className, title, children, href, ...props }, ref) => {
  return (
    <li>
      <Link href={href} legacyBehavior passHref ref={ref} {...props}>
        <NavigationMenuLink
          className={cn(
            "block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
            className
          )}
        >
          <div className="text-sm font-medium leading-none">{title}</div>
          <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">{children}</p>
        </NavigationMenuLink>
      </Link>
    </li>
  )
})
ListItem.displayName = "ListItem"