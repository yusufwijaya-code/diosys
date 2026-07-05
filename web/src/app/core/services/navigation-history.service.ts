import { Injectable } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { filter } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class NavigationHistoryService {
  private _previousUrl = '';

  constructor(router: Router) {
    let currentUrl = '';
    router.events.pipe(filter(e => e instanceof NavigationEnd)).subscribe((e: NavigationEnd) => {
      this._previousUrl = currentUrl;
      currentUrl = e.urlAfterRedirects;
    });
  }

  get previousUrl(): string {
    return this._previousUrl;
  }
}
