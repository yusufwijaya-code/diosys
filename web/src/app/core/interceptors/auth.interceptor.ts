import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, throwError } from 'rxjs';

/**
 * Attaches the bearer token to outgoing requests and redirects to the login
 * page when the API rejects the token.
 */
export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const router = inject(Router);
  const token = localStorage.getItem('diosys_token');

  const authedRequest = token
    ? req.clone({ setHeaders: { Authorization: `Bearer ${token}` } })
    : req;

  return next(authedRequest).pipe(
    catchError((error) => {
      if (error.status === 401 && token) {
        localStorage.removeItem('diosys_token');
        localStorage.removeItem('diosys_user');
        router.navigate(['/admin/login']);
      }
      return throwError(() => error);
    })
  );
};
