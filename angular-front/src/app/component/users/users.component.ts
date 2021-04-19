import {Component, OnInit} from '@angular/core';
import {HttpService} from '../../http.service';
import {User} from '../../structures/user';
import {Column} from '../../structures/column';
import {Router} from '@angular/router';
import {UserStore} from '../../../store/userStore';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})

export class UsersComponent implements OnInit {
  columns = [ new Column('Id', 'users.th.id', 'id'), new Column('Name', 'users.th.name', ''),
    new Column('Phone', 'users.th.phone', ''), new Column('Login', 'users.th.login', '')];
  users: User[];
  filterName = '';
  filterLogin = '';
  filterPhone = '';
  sortBy = '';
  sortType = '';

  constructor(private httpService: HttpService, private router: Router) { }

  submit(): void{
    this.httpService
      .tableGetUsers( this.filterName , this.filterLogin , this.filterPhone, this.sortBy , this.sortType)
      .subscribe((data: User[]) => {this.users = data; console.log(this.users); } );
  }

  ngOnInit(): void {
    this.httpService.getUsers().subscribe((data: User[]) => {this.users = data; console.log(this.users); } );
  }
  setSort(sort): void{
    this.sortBy = sort[0];
    this.sortType = sort[1];
  }
  edit(id: number): void{
    this.router.navigate(['users/edit/' + id]);
  }
}



