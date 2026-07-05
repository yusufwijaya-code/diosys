import { Injectable, signal } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { filter } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class WhatsappContextService {
  private _link = signal('');

  readonly link = this._link.asReadonly();

  constructor(router: Router) {
    router.events.pipe(filter(e => e instanceof NavigationEnd)).subscribe((e: NavigationEnd) => {
      const url = (e as NavigationEnd).urlAfterRedirects;
      // Clear developer override when navigating to home or admin pages
      if (url === '/' || url.startsWith('/#') || url.startsWith('/admin')) {
        this._link.set('');
      }
    });
  }

  setDeveloper(phone: string, name: string): void {
    const number = (phone || '').replace(/[^0-9]/g, '');
    if (!number) { this._link.set(''); return; }
    const text = encodeURIComponent(`Hi ${name}, I'd like to discuss a project with you.`);
    this._link.set(`https://wa.me/${number}?text=${text}`);
  }

  clear(): void {
    this._link.set('');
  }
}
