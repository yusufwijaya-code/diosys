import { Component, OnInit, signal } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { AuthService } from '../../../core/services/auth.service';
import { SpinnerComponent } from '../../../components/spinner/spinner.component';

@Component({
  selector: 'app-admin-login',
  standalone: true,
  imports: [SpinnerComponent],
  template: `
    <div class="login-wrapper">
      <div class="login-card">
        <img src="assets/diosys-logo-mark.svg" alt="Diosys" class="login-logo" />
        <h1>CMS Sign in</h1>
        <p class="subtitle">Authorized administrators only.</p>

        @if (error()) { <div class="alert alert-error">{{ error() }}</div> }

        @if (loading()) {
          <app-spinner size="md" minHeight="56px" />
        } @else {
          <button class="btn btn-primary btn-block gsignin" (click)="signIn()">
            <svg width="18" height="18" viewBox="0 0 48 48" aria-hidden="true">
              <path fill="#FFC107" d="M43.6 20.5H42V20H24v8h11.3C33.7 32.4 29.3 35.5 24 35.5c-6.3 0-11.5-5.2-11.5-11.5S17.7 12.5 24 12.5c2.9 0 5.5 1.1 7.5 2.9l5.7-5.7C33.6 6.5 29.1 4.5 24 4.5 13.2 4.5 4.5 13.2 4.5 24S13.2 43.5 24 43.5c10 0 19.5-7.2 19.5-19.5 0-1.3-.1-2.3-.4-3.5z"/>
              <path fill="#FF3D00" d="M6.3 14.7l6.6 4.8C14.7 15.1 19 12.5 24 12.5c2.9 0 5.5 1.1 7.5 2.9l5.7-5.7C33.6 6.5 29.1 4.5 24 4.5 16.3 4.5 9.7 8.9 6.3 14.7z"/>
              <path fill="#4CAF50" d="M24 43.5c5.2 0 9.6-2 12.9-5.2l-6-5.1c-1.9 1.4-4.3 2.3-6.9 2.3-5.3 0-9.7-3.1-11.3-7.9l-6.6 5.1C9.6 39 16.2 43.5 24 43.5z"/>
              <path fill="#1976D2" d="M43.6 20.5H42V20H24v8h11.3c-.8 2.2-2.3 4.1-4.4 5.2l6 5.1c-.4.4 6.6-4.8 6.6-14.3 0-1.3-.1-2.3-.4-3.5z"/>
            </svg>
            Sign in with Google
          </button>
        }
      </div>
    </div>
  `,
  styles: [`
    .login-logo { height: 120px; width: auto; margin: 0 auto 0.5rem; display: block; }
    .gsignin { gap: 0.6rem; }
  `],
})
export class AdminLoginComponent implements OnInit {
  loading = signal(false);
  error = signal('');

  constructor(private auth: AuthService, private route: ActivatedRoute, private router: Router) {}

  ngOnInit(): void {
    const params = this.route.snapshot.queryParamMap;
    const token = params.get('token');
    const errorMessage = params.get('error');

    if (errorMessage) {
      this.error.set(errorMessage);
      this.router.navigate([], { queryParams: {}, replaceUrl: true });
      return;
    }

    if (token) {
      this.loading.set(true);
      this.auth.completeLogin(token).subscribe({
        next: () => this.router.navigate(['/admin']),
        error: () => {
          this.loading.set(false);
          this.error.set('Could not complete sign-in. Please try again.');
          this.router.navigate([], { queryParams: {}, replaceUrl: true });
        },
      });
      return;
    }

    if (this.auth.isAuthenticated()) {
      this.router.navigate(['/admin']);
    }
  }

  signIn(): void {
    this.auth.redirectToGoogle();
  }
}
