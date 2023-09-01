import { getSearchKeywords } from "@/api/search/getSearch"

export default async function KeywordList() {
     const res = await getSearchKeywords()
    return (
          <>
          <ul className="w-full line-clamp-2">
            {Array.from(res.data).map((_, i) => (
              <li className="float-left mr-4 text-muted-foreground" key={i}>
                {res.data[i]}
              </li>
            ))}
          </ul>
        </>
    )
}