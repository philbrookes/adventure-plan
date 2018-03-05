import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

/*
  Generated class for the UserServiceProvider provider.

  See https://angular.io/guide/dependency-injection for more info on providers
  and Angular DI.
*/
@Injectable()
export class UserService {

  constructor(public http: HttpClient) {
    console.log('Hello UserServiceProvider Provider');
  }

  getUsers() {
    return new Promise((resolve, reject) => {
      this.http.get('http://127.0.0.1:8080/users/echo').subscribe(data => {
        resolve(data);
      }, err => {
        reject(err);
      });
    });
  }

}
