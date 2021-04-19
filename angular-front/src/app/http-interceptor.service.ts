import { Injectable } from '@angular/core';
import {
  HttpInterceptor,
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpErrorResponse
} from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import {catchError} from 'rxjs/operators';
import {MatDialog} from '@angular/material/dialog';
import {ErrorDialogComponent} from './component/error-dialog/error-dialog.component';

@Injectable()
export class Interceptor implements HttpInterceptor {
  urlsToNotUse: Array<string>;
  constructor(public dialog: MatDialog)
  {
    this.urlsToNotUse = [
      '/login/auth',
    ];
  }
  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    if (request.headers.get('skip')){
      return next.handle(request);
    }else{
      return next.handle(request).pipe(
        catchError((error: HttpErrorResponse) => {
          this.dialog.closeAll();
          this.dialog.open(ErrorDialogComponent, {
              width: '250px'
            });
          return throwError(error);
        }));
    }
  }
}
