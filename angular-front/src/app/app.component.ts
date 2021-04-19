import {Component, OnInit} from '@angular/core';
import {TranslateService} from '@ngx-translate/core';
import {NavigationStart, Router} from '@angular/router';
import {HttpService} from './http.service';
import {User} from './structures/user';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  lang: string;

  constructor(private httpService: HttpService, public translate: TranslateService, private router: Router){}

  ngOnInit(): void {
      this.translate.setDefaultLang('en');

      this.router.events.subscribe(
      event  => {
        if (event instanceof NavigationStart){
          if (!!localStorage.getItem('userId')){
            this.httpService.getUserById(localStorage.getItem('userId'))
              .subscribe((data: User) => {
                this.lang = data.Language;
                if (data.Language !== localStorage.getItem('lang')){
                  localStorage.setItem('lang', data.Language);
                  this.translate.setDefaultLang(data.Language);
                }
              });
          } else {
            localStorage.setItem('lang', 'en');
            this.lang = 'en';
          }
        }
      }
    );
  }
}
