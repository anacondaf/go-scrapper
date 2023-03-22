import {PostDTO, PostImages} from "./post.dto";
import {ForbiddenException, Injectable} from "@nestjs/common";
import {HttpService} from "@nestjs/axios";
import {catchError, map} from 'rxjs';

@Injectable()
export class PostService {
    private readonly _httpService: HttpService

    constructor(httpService: HttpService) {
        this._httpService = httpService
    }

    async getAllPosts(){
        return this._httpService
          .get('http://localhost:3000/api/v1/posts')
          .pipe(
            map((res) => {
                console.log(res);
                return res.data;
            }),
          )
          .pipe(
            catchError((err) => {
                console.log(err);
              throw new ForbiddenException('API not available');
            }),
          );
    }
}