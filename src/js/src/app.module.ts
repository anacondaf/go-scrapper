import {Module, Post} from '@nestjs/common';
import { AppController } from './app.controller';
import {ConfigModule} from "@nestjs/config";
import {PostModule} from "./post/post.module";
import {RouterModule} from "@nestjs/core";

@Module({
  imports: [
      PostModule,
      RouterModule.register([
      {
        path: 'posts',
        module: PostModule,
      }]),
      ConfigModule.forRoot()
  ],
  controllers: [AppController],
  providers: [],
})

export class AppModule {}
