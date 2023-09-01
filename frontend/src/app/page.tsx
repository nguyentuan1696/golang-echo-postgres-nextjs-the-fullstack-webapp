import { NewProductList, TagProductList } from "@/components/Products"
import { SearchBarHomePage } from "@/components/Searchbar"
import { getLatestPosts } from "@/api/posts/getPosts"
import { getTags } from "@/api/tags/getTags"
import CategoryList from '@/components/CategoryList'
export default async function Home() {
  const latestPostsData = getLatestPosts({ pageSize: 8, pageIndex: 1 })
  const tagListRes = getTags()
  const [latestPosts, tagList] = await Promise.all([latestPostsData, tagListRes])

  return (
    <div className="space-y-14">
      <SearchBarHomePage />
      <NewProductList data={latestPosts.data} />
      <CategoryList />
      <TagProductList data={tagList.data} />
    </div>
  )
}
