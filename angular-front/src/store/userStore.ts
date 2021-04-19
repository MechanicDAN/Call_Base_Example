import { Store, StoreConfig } from '@datorama/akita';
import {User} from '../app/structures/user';
import {Injectable} from '@angular/core';

export function createInitialState(): User {
  return {
    Id: 0,
    Login: '',
    Name: '',
    Phone: '',
    Language: '',
  };
}
@Injectable({
  providedIn: 'root'
})

@StoreConfig({ name: 'user' })
export class UserStore extends Store<User> {
  constructor() {
    super(createInitialState());
  }
}

