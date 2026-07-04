import { Component, Input } from '@angular/core';

/** Minimal inline line-icon set (stroke = currentColor). */
@Component({
  selector: 'app-icon',
  standalone: true,
  template: `
    <svg [attr.width]="size" [attr.height]="size" viewBox="0 0 24 24" fill="none"
         stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">
      @switch (name) {
        @case ('globe') {
          <circle cx="12" cy="12" r="9" /><path d="M3 12h18M12 3a15 15 0 010 18M12 3a15 15 0 000 18" />
        }
        @case ('smartphone') {
          <rect x="7" y="3" width="10" height="18" rx="2" /><path d="M11 18h2" />
        }
        @case ('cpu') {
          <rect x="6" y="6" width="12" height="12" rx="2" /><rect x="9" y="9" width="6" height="6" />
          <path d="M9 2v2M15 2v2M9 20v2M15 20v2M2 9h2M2 15h2M20 9h2M20 15h2" />
        }
        @case ('layers') {
          <path d="M12 3l9 5-9 5-9-5 9-5z" /><path d="M3 13l9 5 9-5" />
        }
        @case ('code') {
          <path d="M8 8l-4 4 4 4M16 8l4 4-4 4" />
        }
        @case ('mail') {
          <rect x="3" y="5" width="18" height="14" rx="2" /><path d="M3 7l9 6 9-6" />
        }
        @case ('phone') {
          <path d="M5 4h4l2 5-3 2a12 12 0 005 5l2-3 5 2v4a2 2 0 01-2 2A16 16 0 013 6a2 2 0 012-2z" />
        }
        @case ('pin') {
          <path d="M12 21s7-6 7-11a7 7 0 10-14 0c0 5 7 11 7 11z" /><circle cx="12" cy="10" r="2.5" />
        }
        @case ('external') {
          <path d="M14 4h6v6M20 4l-9 9M19 13v6H5V5h6" />
        }
        @case ('github') {
          <path d="M9 19c-4 1.5-4-2.5-6-3m12 5v-3.5a3 3 0 00-.8-2.3c2.6-.3 5.3-1.3 5.3-5.8a4.5 4.5 0 00-1.2-3.1 4.2 4.2 0 00-.1-3.1s-1-.3-3.3 1.2a11.4 11.4 0 00-6 0C6.6 2.6 5.6 2.9 5.6 2.9a4.2 4.2 0 00-.1 3.1A4.5 4.5 0 004.3 9c0 4.5 2.7 5.5 5.3 5.8a3 3 0 00-.8 2.3V21" />
        }
        @case ('arrow') {
          <path d="M5 12h14M13 6l6 6-6 6" />
        }
        @case ('check') {
          <path d="M20 6L9 17l-5-5" />
        }
        @default {
          <path d="M12 3l2.5 6.5L21 12l-6.5 2.5L12 21l-2.5-6.5L3 12l6.5-2.5z" />
        }
      }
    </svg>
  `,
  styles: [':host { display: inline-flex; align-items: center; line-height: 0; flex-shrink: 0; }'],
})
export class IconComponent {
  @Input() name = 'default';
  @Input() size = 22;
}
