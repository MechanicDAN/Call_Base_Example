import { Query } from '@datorama/akita';
import {UserStore} from './userStore';
import {User} from '../app/structures/user';
import {Observable} from 'rxjs';
import {Injectable} from '@angular/core';

@Injectable({
  providedIn: 'root'
})

export class UserQuery extends Query<User> {
  allState = this.select();
  constructor(protected store: UserStore) {
    super(store);
  }
  getId(): Observable<number> {
    return this.select(state => state.Id);
  }
  getName(): Observable<string> {
    return this.select(state => state.Name);
  }
  getLogin(): Observable<string> {
    return this.select(state => state.Login);
  }
  getPhone(): Observable<string> {
    return this.select(state => state.Phone);
  }
  getLanguage(): Observable<string> {
    return this.select(state => state.Language);
  }
  getAll(): Observable<User> {
    return this.select();
  }
}
