import {Controller, Get} from "@nestjs/common";
import {PostService} from "./post.service";
import {PostDTO} from "./post.dto";

@Controller()
export class PostController {
    private readonly _postService: PostService

    constructor(postService: PostService) {
        this._postService = postService
    }

    @Get()
    getAllPosts() {
        return this._postService.getAllPosts();
    }
}