import KeywordList from "@/components/Searchbar/KeywordList"
import SearchBarMain from "@/components/Searchbar/SearchBarMain"

export default async function SearchBarHomePage() {
  return (
    <div className="flex items-center justify-center">
      <div className=" space-y-2">
        <SearchBarMain />
        <KeywordList />
      </div>
    </div>
  )
}
