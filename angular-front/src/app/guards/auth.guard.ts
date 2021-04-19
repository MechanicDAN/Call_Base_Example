import { Injectable } from '@angular/core';
import {CanActivate,  Router} from '@angular/router';
import {HttpService} from '../http.service';
import {Md5} from 'ts-md5';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private httpService: HttpService, private router: Router) { }

  canActivate(): boolean {
    if (!!localStorage.getItem('token') &&
      localStorage.getItem('token') === Md5.hashStr(localStorage.getItem('userId') + 23) as string){
      return true;
    } else {
      this.router.navigate(['login']).then(r => {});
      return false;
    }
  }
}
