import {IDTO} from "../IDto"

export class PostImages {
    private readonly Id: string
    private readonly Url: string
    private readonly PostId: string

    constructor(id: string, url: string, postId: string) {
        this.Id = id;
        this.Url = url;
        this.PostId = postId;
    }
}

export class PostDTO implements IDTO {
    private readonly Id: string
    private readonly Title: string
    private readonly PostImages: PostImages[]

    constructor(id: string, title: string, postImages: PostImages[]) {
        this.Id = id;
        this.Title = title;
        this.PostImages = postImages;
    }
}