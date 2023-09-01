import { cache } from "react"
import { ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

/**
 * getIdSlug
 */
export function getIdSlug(slug: string): string | undefined {
  return slug.split("-").pop()
}

/**
 * cn
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

/**
 * formatDate
 */
export function formatDate(input: string | number): string {
  const date = new Date(input)
  return date.toLocaleDateString("en-US", {
    month: "long",
    day: "numeric",
    year: "numeric",
  })
}

/**
 * hasProtocol checks protocol of url
 */
export function hasProtocol(url: string): boolean {
  return /^(?:\w*:|\/\/)/.test(url)
}

/**
 * isInternalUrlUtils
 */
export function isInternalUrlUtils(url?: string): boolean {
  return typeof url !== "undefined" && !hasProtocol(url)
}

/**
 * titleFormatter
 */
export function titleFormatter(title?: string | undefined): string {
  return title?.trim().length ? `${title.trim()} - Thích Tiếng Anh` : "Thích Tiếng Anh"
}

/**
 * getBaseUrl
 */
export const getBaseUrl = cache(() => {
  return process.env.NODE_ENV === "development" ? process.env.DOCS_API_BASE_URL : process.env.DOCS_API_BASE_URL
})

/**
 *
 */
export function toCapFirstLetter(title: string): string {
  return title
    .toLowerCase()
    .split(" ")
    .map((s) => s.charAt(0).toUpperCase() + s.substring(1))
    .join(" ")
}

/**
 *
 */
export function toCapFirstWordString(title: string): string {
  return title.charAt(0).toUpperCase() + title.slice(1).toLowerCase()
}

/**
 * decodeQueryParam Decoding query parameters from a URL
 */
export function decodeQueryParam(s: string): string {
  return decodeURIComponent(s.replace(/\+/g, " "))
}

export function randomValueArray(array: string[]): string {
  return array[Math.floor(Math.random() * array.length)]
}
