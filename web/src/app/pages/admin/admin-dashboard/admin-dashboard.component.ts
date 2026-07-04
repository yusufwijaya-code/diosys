import { Component, OnInit, signal } from '@angular/core';
import { RouterLink } from '@angular/router';

import { CmsService } from '../../../core/services/cms.service';

@Component({
  selector: 'app-admin-dashboard',
  standalone: true,
  imports: [RouterLink],
  template: `
    <div class="admin-topbar"><h1>Dashboard</h1></div>

    <div class="admin-grid">
      <a class="dashboard-tile" routerLink="/admin/developers">
        <div class="tile-value">{{ developerCount() }}</div>
        <div class="tile-title">Developers</div>
      </a>
      <a class="dashboard-tile" routerLink="/admin/inbox">
        <div class="tile-value">{{ unreadCount() }}</div>
        <div class="tile-title">Unread messages</div>
      </a>
      <a class="dashboard-tile" routerLink="/admin/services">
        <div class="tile-value">{{ serviceCount() }}</div>
        <div class="tile-title">Services</div>
      </a>
      <a class="dashboard-tile" routerLink="/admin/settings">
        <div class="tile-icon">⚙</div>
        <div class="tile-title">Settings</div>
      </a>
    </div>
  `,
})
export class AdminDashboardComponent implements OnInit {
  developerCount = signal(0);
  unreadCount = signal(0);
  serviceCount = signal(0);

  constructor(private cms: CmsService) {}

  ngOnInit(): void {
    this.cms.getDevelopers().subscribe((d) => this.developerCount.set(d.length));
    this.cms.getServices().subscribe((s) => this.serviceCount.set(s.length));
    this.cms.getMessages().subscribe((m) =>
      this.unreadCount.set(m.filter((x) => x.isRead === 0 && x.isArchived === 0).length)
    );
  }
}
