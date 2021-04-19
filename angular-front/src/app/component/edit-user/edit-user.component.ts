import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {UserQuery} from '../../../store/userQuery';
import {HttpService} from '../../http.service';
import {ActivatedRoute, Router} from '@angular/router';
import {User} from '../../structures/user';
import {UserStore} from '../../../store/userStore';
import {Subscription} from 'rxjs';
import {equals} from '@ngx-translate/core/lib/util';

@Component({
  selector: 'app-edit-user',
  templateUrl: './edit-user.component.html',
  styleUrls: ['./edit-user.component.css']
})
export class EditUserComponent implements OnInit, OnDestroy{
  private subscr: Subscription;
  Form: FormGroup;
  isVisible = false;
  id: number;

  constructor(private httpService: HttpService, private router: Router, private fb: FormBuilder,
              public userQuery: UserQuery, private activatedRoute: ActivatedRoute, public userStore: UserStore) {}

  ngOnInit(): void {
    this.initForm();
    this.id = this.activatedRoute.snapshot.params.id;
    this.subscr = this.httpService.getUserById(this.id.toString()).subscribe((data: User) => {
        this.userStore.update(data);
      });
    this.userQuery.allState.subscribe(res => {
      this.Form.patchValue(res, {emitEvent: false});
    });
  }
  ngOnDestroy(): void {
    this.subscr.unsubscribe();
  }

  initForm(): void {
      this.Form = this.fb.group({
        Name: ['', [ Validators.required]],
        Login: ['', [ Validators.required]],
        Phone: [''],
        Language: ['', [ Validators.required]]
      });
      this.onChanges();
    }
  onChanges(): void {
    this.Form.valueChanges.subscribe(val => {
      this.userStore.update(this.Form.value);
      this.isVisible = false;
    });
  }

  save(): void {
    if (this.Form.valid) {
      this.subscr = this.userQuery.allState.subscribe(res => {
        console.log(this.id);
        if (res.Id === 0) {
           this.httpService.createUser(res).subscribe(() => {});
         } else {
           this.httpService.updateUser(this.id, res).subscribe(() => {});
         }
         });
      this.router.navigate(['users']);
    } else {
      this.isVisible = true;
    }
  }
}
