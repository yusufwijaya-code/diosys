import { Component, OnInit, signal } from '@angular/core';
import {
  Router,
  RouterOutlet,
  NavigationStart,
  NavigationEnd,
  NavigationCancel,
  NavigationError,
} from '@angular/router';
import { ViewportScroller } from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet],
  template: `
    <div class="route-loader" [class.visible]="navigating()"></div>
    <router-outlet />
  `,
  styles: [`
    .route-loader {
      position: fixed;
      top: 0;
      left: 0;
      right: 0;
      height: 2px;
      z-index: 10000;
      pointer-events: none;
      opacity: 0;
      transition: opacity 0.25s;
    }
    .route-loader.visible {
      opacity: 1;
      background: linear-gradient(
        90deg,
        transparent 0%,
        #a5b4fc 35%,
        #818cf8 65%,
        transparent 100%
      );
      background-size: 250% 100%;
      animation: route-shimmer 1.4s linear infinite;
    }
    @keyframes route-shimmer {
      0%   { background-position: 120% 0; }
      100% { background-position: -120% 0; }
    }
  `],
})
export class AppComponent implements OnInit {
  navigating = signal(false);
  private _hideTimer: ReturnType<typeof setTimeout> | null = null;

  constructor(
    scroller: ViewportScroller,
    private router: Router,
  ) {
    scroller.setOffset([0, 88]);
  }

  ngOnInit(): void {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationStart) {
        if (this._hideTimer) clearTimeout(this._hideTimer);
        this.navigating.set(true);
      } else if (
        event instanceof NavigationEnd ||
        event instanceof NavigationCancel ||
        event instanceof NavigationError
      ) {
        this._hideTimer = setTimeout(() => this.navigating.set(false), 350);
      }
    });
  }
}
