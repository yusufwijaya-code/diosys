import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';

import { AuthService } from '../../../core/services/auth.service';

@Component({
  selector: 'app-admin-layout',
  standalone: true,
  imports: [RouterLink, RouterLinkActive, RouterOutlet],
  template: `
    <div class="admin-shell">
      <aside class="admin-sidebar">
        <div class="brand"><span>Diosys CMS</span></div>
        <nav class="admin-nav">
          <a routerLink="/admin/dashboard" routerLinkActive="active"><span class="label">Dashboard</span></a>
          <a routerLink="/admin/developers" routerLinkActive="active"><span class="label">Developers</span></a>
          <a routerLink="/admin/projects" routerLinkActive="active"><span class="label">Projects</span></a>
          <a routerLink="/admin/inbox" routerLinkActive="active"><span class="label">Inbox</span></a>
          <a routerLink="/admin/services" routerLinkActive="active"><span class="label">Services</span></a>
          <a routerLink="/admin/pricing" routerLinkActive="active"><span class="label">Pricing</span></a>
          <a routerLink="/admin/settings" routerLinkActive="active"><span class="label">Settings</span></a>
        </nav>
        <div class="sidebar-footer">
          <div class="full muted">{{ user?.fullName }}</div>
          <a routerLink="/" class="btn btn-secondary btn-sm btn-block" style="margin-bottom:0.4rem">← View site</a>
          <button class="btn btn-secondary btn-sm btn-block" (click)="logout()">Sign out</button>
        </div>
      </aside>

      <main class="admin-main">
        <router-outlet />
      </main>
    </div>
  `,
})
export class AdminLayoutComponent {
  constructor(private auth: AuthService) {}

  get user() {
    return this.auth.currentUser;
  }

  logout(): void {
    this.auth.logout();
  }
}
