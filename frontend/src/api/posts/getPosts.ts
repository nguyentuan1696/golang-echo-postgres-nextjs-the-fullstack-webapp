import { notFound } from "next/navigation"
import { getBaseUrl, getIdSlug } from "@/lib/utils"
import type { LatestPosts } from "@/types/LatestPosts"
import type { Post } from "@/types/Post"
import type { PostsInSameCategory } from "@/types/PostsInSameCategory"
// `server-only` guarantees any modules that import code in file
// will never run on the client. Even though this particular api
// doesn't currently use sensitive environment variables, it's
// good practise to add `server-only` preemptively.
import "server-only"

/**
 * GetPost
 */
export async function getPost({ slug }: { slug: string }) {
  const res = await fetch(`${getBaseUrl()}/docs/api/v1/post/${getIdSlug(slug)}`, {
    cache: process.env.NODE_ENV === "development" ? "no-store" : "force-cache",
  })

  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const post = (await res.json()) as Post

  if (Object.keys(post.data).length === 0) {
    notFound()
  }

  return post
}

/**
 * getPostsInSameCategry get list of related posts by category
 */
export async function getPostsInSameCategry(slug: string) {
  const res = await fetch(`${getBaseUrl()}/docs/api/v1/post/same-category?id=${getIdSlug(slug)}&pageSize=5&pageIndex=1`)

  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const post = (await res.json()) as PostsInSameCategory

  if (Object.keys(post.data).length === 0) {
    notFound()
  }

  return post
}

/**
 * getLatestPosts get list of last 4 created posts
 */
export async function getLatestPosts({ pageSize, pageIndex }: { pageSize?: number; pageIndex?: number }) {
  const res = await fetch(`${getBaseUrl()}/docs/api/v1/post/latest?pageSize=${pageSize}&pageIndex=${pageIndex}`)

  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const post = (await res.json()) as LatestPosts

  if (Object.keys(post.data).length === 0) {
    notFound()
  }

  return post
}
