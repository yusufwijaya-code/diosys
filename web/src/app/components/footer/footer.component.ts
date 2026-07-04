import { Component, Input } from '@angular/core';
import { RouterLink } from '@angular/router';

import { SettingsMap } from '../../core/models/diosys.model';

@Component({
  selector: 'app-footer',
  standalone: true,
  imports: [RouterLink],
  template: `
    <footer class="footer section-secondary">
      <div class="container footer-inner">
        <div class="footer-brand">
          <img src="assets/diosys-logo-mark.svg" alt="Diosys" class="footer-logo" />
          <p class="muted">{{ settings['companyTagline'] || 'Premium technology solutions.' }}</p>
        </div>

        <div class="footer-links">
          <a [routerLink]="['/']" fragment="services">Services</a>
          <a [routerLink]="['/']" fragment="work">Work</a>
          <a [routerLink]="['/']" fragment="developers">Developers</a>
          <a [routerLink]="['/']" fragment="contact">Contact</a>
          @if (settings['contactEmail']) {
            <a [href]="'mailto:' + settings['contactEmail']">{{ settings['contactEmail'] }}</a>
          }
        </div>
      </div>
      <div class="container copyright">
        <span class="muted">© {{ year }} {{ settings['companyName'] || 'Diosys' }}. All rights reserved.</span>
      </div>
    </footer>
  `,
  styles: [`
    .footer { padding: 4rem 0 2rem; border-top: 1px solid var(--border); }
    .footer-inner { display: flex; justify-content: space-between; gap: 2rem; flex-wrap: wrap; }
    .footer-logo { height: 96px; width: auto; margin-bottom: 0.5rem; margin-left: -0.5rem; }
    .footer-brand { max-width: 320px; }
    .footer-links { display: flex; flex-direction: column; gap: 0.6rem; }
    .footer-links a { color: var(--text-tertiary); font-size: 0.92rem; }
    .footer-links a:hover { color: var(--accent); }
    .copyright {
      display: flex; justify-content: space-between; align-items: center;
      margin-top: 3rem; padding-top: 1.5rem; border-top: 1px solid var(--border);
    }
  `],
})
export class FooterComponent {
  @Input() settings: SettingsMap = {};
  year = new Date().getFullYear();
}
