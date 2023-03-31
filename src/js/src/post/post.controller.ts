import { Controller, Get, Inject } from '@nestjs/common';
import { PostService } from './post.service';
import {
  ClientProxy,
  Ctx,
  EventPattern,
  Payload,
  RmqContext,
} from '@nestjs/microservices';

@Controller({
  version: '1',
})
export class PostController {
  private readonly _postService: PostService;

  constructor(
    postService: PostService,
    @Inject('RMQ_SERVICE') private client: ClientProxy,
  ) {
    this._postService = postService;
  }

  @Get()
  getAllPosts() {
    return this._postService.getAllPosts();
  }

  @EventPattern('post')
  getAllPostsNotification(@Payload() data: any, @Ctx() context: RmqContext) {
    console.log(data);
  }

  @EventPattern('hello_message')
  getEventMessage(@Payload() data: any, @Ctx() context: RmqContext) {
    console.log('Some event has been sent!');
  }
}
