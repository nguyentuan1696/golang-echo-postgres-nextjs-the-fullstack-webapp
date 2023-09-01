import { notFound } from "next/navigation"
//import { getBaseUrl } from "@/lib/utils";
import type { Search } from "@/types/Search"
import type { SearchKeywords } from "@/types/SearchKeywords"

export async function getSearch(slug: string) {
  const res = await fetch(`http://34.124.177.152:9091/docs/api/v1/search?query=${slug}`, {
    cache: process.env.NODE_ENV === "development" ? "no-store" : "force-cache",
  })
  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const posts = (await res.json()) as Search

  // if (Array.from(posts.data.data).length === 0) {
  //   notFound()
  // }

  return posts
}

export async function getSearchKeywords() {
  const res = await fetch(`http://34.124.177.152:9091/docs/api/v1/search/most-keyword?pageSize=5&pageIndex=1`, {
    cache: process.env.NODE_ENV === "development" ? "no-store" : "force-cache",
  })
  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const keywords = (await res.json()) as SearchKeywords

  if (Array.from(keywords.data).length === 0) {
    notFound()
  }

  return keywords
}
