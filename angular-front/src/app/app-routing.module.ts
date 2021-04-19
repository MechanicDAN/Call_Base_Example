import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from './component/login/login.component';
import {AuthorizedComponent} from './component/authorized/authorized.component';
import {AuthGuard} from './guards/auth.guard';
import {AlreadyAuthGuard} from './guards/already-auth.guard';
import {UsersComponent} from './component/users/users.component';
import {CallsComponent} from './component/calls/calls.component';
import {EditUserComponent} from './component/edit-user/edit-user.component';

const routes: Routes = [
  { path: 'login',
    component: LoginComponent,
    canActivate: [AlreadyAuthGuard]
  },
  { path: 'authorized',
    component: AuthorizedComponent,
    canActivate: [AuthGuard]
  },
  { path: 'users',
    component: UsersComponent,
    canActivate: [AuthGuard]
  },
  { path: 'users/edit/:id',
    component: EditUserComponent,
    canActivate: [AuthGuard]
  },
  { path: 'calls',
    component: CallsComponent,
    canActivate: [AuthGuard]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
