import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {User} from './structures/user';
import {Call} from './structures/call';

@Injectable({
  providedIn: 'root'
})
export class HttpService {
  constructor(private http: HttpClient) {}
  flag: boolean;

  postLogin(log: string, pass: string): Observable<any> {
    const body = new FormData();
    body.append('login', log);
    body.append('password', pass);
    return this.http.post('/login/auth', body, {headers: {skip: 'true'}});
  }

  getUsers(): Observable<User[]> {
    return this.http.get<User[]>('/api/users');
  }
  getUserById(id: string): Observable<User> {
    return this.http.get<User>('/api/users/' + id);
  }
  tableGetUsers(filterName, filterLogin, filterPhone, sortBy, sortType): Observable<User[]> {
    return this.http.get<User[]>('/api/users' + '?filterName=' + filterName + '&filterLogin=' + filterLogin +
      '&filterPhone=' + filterPhone + '&sortBy=' + sortBy + '&sortType=' + sortType);
  }
  updateUser(id, user: User): Observable<any> {
    const body = new FormData();
    body.append('Login', user.Login);
    body.append('Name', user.Name);
    body.append('Phone', user.Phone);
    body.append('Language', user.Language);
    return this.http.post<any>('/api/users/' + id + '/update', body);
  }
  createUser(user: User): Observable<any> {
    const body = new FormData();
    body.append('Login', user.Login);
    body.append('Name', user.Name);
    body.append('Phone', user.Phone);
    body.append('Language', user.Language);
    return this.http.post<any>('/api/users/create', body);
  }

  getCalls(): Observable<Call[]> {
    return this.http.get<Call[]>('/api/calls');
  }
  tableGetCalls(filterUser, sortBy, sortType): Observable<Call[]> {
    return this.http.get<Call[]>('/api/calls' + '?filterUser=' + filterUser + '&sortBy=' + sortBy + '&sortType=' + sortType);
  }
}
