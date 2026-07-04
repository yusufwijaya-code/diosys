import { Component, OnInit, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';

import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { FooterComponent } from '../../../components/footer/footer.component';
import { WhatsappFloatComponent } from '../../../components/whatsapp-float/whatsapp-float.component';
import { PublicService } from '../../../core/services/public.service';
import { SettingsMap } from '../../../core/models/diosys.model';

@Component({
  selector: 'app-public-layout',
  standalone: true,
  imports: [RouterOutlet, NavbarComponent, FooterComponent, WhatsappFloatComponent],
  template: `
    <app-navbar />
    <main class="public-main">
      <router-outlet />
    </main>
    <app-footer [settings]="settings()" />
    <app-whatsapp-float [settings]="settings()" />
  `,
  styles: [`
    .public-main { min-height: 60vh; }
  `],
})
export class PublicLayoutComponent implements OnInit {
  settings = signal<SettingsMap>({});

  constructor(private publicService: PublicService) {}

  ngOnInit(): void {
    this.publicService.getSettings().subscribe({
      next: (settings) => this.settings.set(settings),
      error: () => this.settings.set({}),
    });
  }
}
