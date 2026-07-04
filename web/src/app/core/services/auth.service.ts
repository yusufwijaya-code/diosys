import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';

import { ApiService } from './api.service';
import { AuthUser } from '../models/api-response.model';
import { environment } from '../../../enviroments/environment';

const TOKEN_KEY = 'diosys_token';
const USER_KEY = 'diosys_user';

/** Handles Google sign-in (server-side redirect flow) and token persistence. */
@Injectable({ providedIn: 'root' })
export class AuthService {
  constructor(private api: ApiService, private router: Router) {}

  /** Sends the browser to the backend, which redirects to Google's consent screen. */
  redirectToGoogle(): void {
    window.location.href = `${environment.apiUrl}/auth/google/redirect`;
  }

  /**
   * Completes login after the backend redirect: stores the token returned in the
   * URL, then loads the authenticated user from /auth/me.
   */
  completeLogin(token: string): Observable<AuthUser> {
    localStorage.setItem(TOKEN_KEY, token);
    return this.api.get<AuthUser>('/auth/me').pipe(
      tap((user) => localStorage.setItem(USER_KEY, JSON.stringify(user)))
    );
  }

  logout(): void {
    localStorage.removeItem(TOKEN_KEY);
    localStorage.removeItem(USER_KEY);
    this.router.navigate(['/admin/login']);
  }

  get token(): string | null {
    return localStorage.getItem(TOKEN_KEY);
  }

  isAuthenticated(): boolean {
    return !!this.token;
  }

  get currentUser(): AuthUser | null {
    const raw = localStorage.getItem(USER_KEY);
    return raw ? (JSON.parse(raw) as AuthUser) : null;
  }
}
