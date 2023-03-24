import {Module} from '@nestjs/common';
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
  controllers: [],
  providers: [],
})

export class AppModule {}
