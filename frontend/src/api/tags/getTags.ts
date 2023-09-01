import { notFound } from "next/navigation";
import { getBaseUrl } from "@/lib/utils";
import type { PostByTag } from "@/types/PostByTag";
import type { Tag } from "@/types/Tag";
import type { TagInfo } from "@/types/TagInfo";
// `server-only` guarantees any modules that import code in file
// will never run on the client. Even though this particular api
// doesn't currently use sensitive environment variables, it's
// good practise to add `server-only` preemptively.
import "server-only";


/**
 * getTagInfo get info tag page
 */
export async function getTagInfo(subTag: string) {
  const res = await fetch(`${getBaseUrl()}/docs/api/v1/tag/info?slug=${subTag}`, {
    cache: process.env.NODE_ENV === "development" ? "no-store" : "force-cache",
  })

  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const tagInfo = (await res.json()) as TagInfo

  if (Object.keys(tagInfo.data).length === 0 || tagInfo.data === null) {
    notFound()
  }

  return tagInfo
}

/**
 * getTags
 */
export async function getTags() {
  const res = await fetch(`${getBaseUrl()}/docs/api/v1/tag`)

  if (!res.ok) {
    // Render the closest `error.js` Error Boundary
    throw new Error("Something went wrong!")
  }

  const tags = (await res.json()) as Tag

  if (tags.data.length === 0) {
    // Render the closest `not-found.js` Error Boundary
    notFound()
  }
  return tags
}

/**
 * getTag
 */
export async function getTag({ slug }: { slug: string }) {
  const res = await fetch(`${getBaseUrl()}/docs/api/v1/tag/${slug}`)

  if (!res.ok) {
    throw new Error("Something went wrong!")
  }

  const postsByTag = (await res.json()) as PostByTag

  if (postsByTag.data.length === 0) {
    notFound()
  }

  return postsByTag
}