import { getPostsInSameCategry } from '@/api/posts/getPosts';
import { toCapFirstWordString } from '@/lib/utils'
import Link from 'next/link'


export default async function Layout({ children, params }: { children: React.ReactNode, params: {subPost: string} }) {
    const res = await getPostsInSameCategry(params.subPost)

  return (
    <div className="grid grid-cols-1 md:grid-cols-[25%_minmax(50%,_1fr)_25%]">
      <div className="hidden md:inline">sidebar</div>
      <div className="">{children}</div>
       <div className="order-last md:pl-6 mt-10 md:mt-0">
            <h2 className="font-heading mt-12 scroll-m-20 border-b pb-2 text-xl font-semibold tracking-tight first:mt-0">{toCapFirstWordString("Bài viết cùng chuyên mục")}</h2>
            <ul className="pt-3">
              {Array.from(res.data)
                .filter((p) => p.slug != params.subPost)
                .map((d, i) => (
                  <Link href={`p/${d.slug}`}>
                    <li key={i} className="mt-2 border-b border-dotted pb-1">
                      <p className='line-clamp-2'>{toCapFirstWordString(d.title)}</p>
                    </li>
                  </Link>
                ))}
            </ul>
          </div>
    </div>
  )
}