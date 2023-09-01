import { ReactNode } from "react";
import Balancer from "@/components/Balancer";
import Link from "@/components/Link";
import { Separator } from "@/ui/separator";
import { getCategoryInfo, getChildCategoryList } from "@/api/categories/getCatgories";
import { toCapFirstLetter, toCapFirstWordString } from "@/lib/utils"
import type { CategoryInfoSlugProps } from "@/types/CategoryInfoSlugProps";


export async function generateMetadata({ params }: { params: CategoryInfoSlugProps }) {
  const res = await getCategoryInfo({
    categorySlug: params.slug[0],
    subCategorySlug: params.slug[1] === undefined ? "" : params.slug[1],
  })

  return {
    title: res.data.title,
  }
}

async function getChildCategorys(parentCate: string) {
  return await getChildCategoryList(parentCate)
  
}

async function getInfoCategory({ categorySlug, subCategorySlug }: { categorySlug: string; subCategorySlug: string }) {
  return await getCategoryInfo({ categorySlug, subCategorySlug })
  
}

export default async function Layout({ children, params }: { children: ReactNode; params: CategoryInfoSlugProps }) {
  const resChildCategoryList = getChildCategorys(params.slug[0])
  const resInfoCategory = getInfoCategory({
    categorySlug: params.slug[0],
    subCategorySlug: params.slug[1] === undefined ? "" : params.slug[1],
  })
  const [resChild, resInfo] = await Promise.all([resChildCategoryList, resInfoCategory])

  return (
    <div className="">
      <div className="">
        <div className="space-y-2">
          <h1 className="scroll-m-20 text-4xl font-bold tracking-tight">Chuyên mục: {toCapFirstLetter(resInfo.data.title)}</h1>
          <p className="text-lg text-muted-foreground">
            <Balancer title={resInfo.data.description} />
          </p>
        </div>
        <Separator className="my-4 md:my-6" />
        <div className="grid grid-cols-1 md:grid-cols-[25%_minmax(50%,_1fr)_25%]">
          <div className="hidden md:inline"></div>
          <div className="">{children}</div>
          <div className="order-last md:pl-6 mt-10 md:mt-0">
            <h2 className="font-heading mt-12 scroll-m-20 border-b pb-2 text-xl font-semibold tracking-tight first:mt-0">{toCapFirstWordString("chuyên mục cùng chủ đề")}</h2>
            <ul className="pt-3">
              {Array.from(resChild.data)
                .filter((p) => p.slug != params.slug[1])
                .map((d, i) => (
                  <Link href={`c/${params.slug[0]}/${d.slug}`}>
                    <li key={i} className="mt-2 border-b border-dotted pb-1">
                      <p>{toCapFirstWordString(d.title)}</p>
                    </li>
                  </Link>
                ))}
            </ul>
          </div>
        </div>
      </div>
    </div>
  )
}