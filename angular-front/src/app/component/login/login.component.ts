import {Component, OnInit} from '@angular/core';
import {HttpService} from '../../http.service';
import {Router} from '@angular/router';
import {User} from '../../structures/user';
import {TranslateService} from '@ngx-translate/core';
import {Md5} from 'ts-md5/dist/md5';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit{
  isVisible = false;
  login = '';
  password = '';

  constructor(private httpService: HttpService, private router: Router, public translate: TranslateService) {}

  submit(): void {
    this.httpService.postLogin(this.login, this.password).subscribe(
      (data: User) => {
        this.translate.setDefaultLang(data.Language);
        localStorage.setItem('userId', data.Id.toString());
        localStorage.setItem('lang', data.Language);
        localStorage.setItem('token', Md5.hashStr(data.Id.toString() + 23) as string);
        this.router.navigate(['authorized']);
        this.isVisible = false;
      }, error => {
      console.log(error);
      this.isVisible = true;
    }
    );
  }

  ngOnInit(): void { }

  closeWarning(): void {
      this.isVisible = false;
  }
}
