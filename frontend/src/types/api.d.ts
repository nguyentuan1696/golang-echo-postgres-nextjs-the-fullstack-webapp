declare module "@/types/TagInfoSlugProps" {
  export type TagInfoSlugProps = {
    subTag: string
  }
}

declare module "@/types/Search" {
  export type Search = {
    message: string
    data: {
      total: number
      data: {
        title: string
        slug: string
      }[]
    }
  }
}


declare module "@/types/SearchKeywords" {
  export type SearchKeywords = {
    message: string
    data: string[]
  }
}


declare module "@/types/CategoryInfoSlugProps" {
  export type CategoryInfoSlugProps = {
    slug: string[]
  }
}

declare module "@/types/CategorySlug" {
  export type CategorySlug = {
    slug: string[]
    params: {
      slug: string[]
    }
  }
}

declare module "@/types/CategoryInfo" {
  export type CategoryInfo = {
    message: string
    data: {
      title: string
      description: string
    }
  }
}

declare module "@/types/LatestPosts" {
  export type LatestPosts = {
    message?: string
    data: {
      id: string
      title: string
      slug: string
    }[]
  }
}

declare module "@/types/PostsInSameCategory" {
  export type PostsInSameCategory = {
    message: string
    data: {
      id: string
      title: string
      slug: string
    }[]
  }
}

declare module "@/types/ChildCategoryList" {
  export type ChildCategoryList = {
    message: string
    data: {
      id: string
      title: string
      slug: string
    }[]
  }
}

declare module "@/types/ChildCategoryListView" {
  export type ChildCategoryListView = {
    data: {
      id: string
      title: string
      slug: string
    }[]
  }
}

declare module "@/types/Category" {
  export type Category = {
    message: string
    data: {
      total_post: number
      category_info: {
        description: string
        title: string
      }
      list_post: {
        id_post: string
        title_post: string
        slug_post: string
        category_post: string
      }[]
    }
  }
}

declare module "@/types/Post" {
  export type Post = {
    message: string
    data: {
      post_detail: {
        id: string
        content: string
        description: string
        slug: string
        title: string
      }
      source_detail: {
        source_id: number
        source_title: string
        source_url: string
      }[]
      tag_list_by_post: {
        post_id: string
        tag_id: number
        slug_tag: string
        title_tag: string
      }[]
    }
  }
}

declare module "@/types/Tag" {
  export type Tag = {
    message?: string
    data: {
      id: number
      title: string
      slug: string
      description: string
    }[]
  }
}

declare module "@/types/PostByTag" {
  export type PostByTag = {
    message: string
    data: {
      id: string
      title: string
      slug: string
    }[]
  }
}

declare module "@/types/TagInfo" {
  export type TagInfo = {
    message: string
    data: {
      id: number
      title: string
      description: string
    }
  }
}
