import { notFound } from "next/navigation"
import { getBaseUrl } from "@/lib/utils"
import type { Category } from "@/types/Category"
import type { CategoryInfo } from "@/types/CategoryInfo"
import type { ChildCategoryList } from "@/types/ChildCategoryList"
// `server-only` guarantees any modules that import code in file
// will never run on the client. Even though this particular api
// doesn't currently use sensitive environment variables, it's
// good practise to add `server-only` preemptively.
import "server-only"

/**
 * getCategoryInfo get category info
 */
export async function getCategoryInfo({ categorySlug, subCategorySlug }: { categorySlug: string; subCategorySlug: string }) {
  const res = await fetch(`${getBaseUrl()}/docs/api/v1/category/info?parentCate=${categorySlug}&childCate=${subCategorySlug}`, {
    cache: process.env.NODE_ENV === "development" ? "no-store" : "force-cache",
  })

  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const categoryInfo = (await res.json()) as CategoryInfo

  if (categoryInfo.data === null) {
    notFound()
  }

  return categoryInfo
}

/**
 * getPostByCategory Get post list by category
 */
export async function getPostByCategory({ categorySlug, subCategorySlug }: { categorySlug: string; subCategorySlug: string }) {
  const res = await fetch(`${getBaseUrl()}/docs/api/v1/category/post?parentCate=${categorySlug}&childCate=${subCategorySlug}`, {
    cache: process.env.NODE_ENV === "development" ? "no-store" : "force-cache",
  })

  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const categories = (await res.json()) as Category

  if (categories.data.list_post.length === 0 || categories.data.total_post === 0) {
    notFound()
  }

  return categories
}

/**
 *
 */
export async function getChildCategoryList(parentCate: string) {
  const res = await fetch(`${getBaseUrl()}/docs/api/v1/category/child-list?parentCate=${parentCate}&pageSize=5&pageIndex=1`, {
    cache: process.env.NODE_ENV === "development" ? "no-store" : "force-cache",
  })

  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const categories = (await res.json()) as ChildCategoryList

  if (categories.data.length === 0) {
    notFound()
  }

  return categories
}
