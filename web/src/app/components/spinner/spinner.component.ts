import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-spinner',
  standalone: true,
  template: `
    <div class="spinner-wrap" [style.min-height]="minHeight">
      <div class="spinner" [class]="'spinner--' + size">
        <div class="spinner-ring"></div>
        <div class="spinner-ring spinner-ring--inner"></div>
      </div>
    </div>
  `,
  styles: [`
    .spinner-wrap {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 100%;
    }

    .spinner {
      position: relative;
      flex-shrink: 0;
    }

    .spinner--sm  { width: 28px; height: 28px; }
    .spinner--md  { width: 44px; height: 44px; }
    .spinner--lg  { width: 64px; height: 64px; }

    .spinner-ring {
      position: absolute;
      inset: 0;
      border-radius: 50%;
      border: 2px solid transparent;
      border-top-color: #a5b4fc;
      border-right-color: rgba(165, 180, 252, 0.25);
      animation: spinner-rotate 0.85s cubic-bezier(0.4, 0, 0.2, 1) infinite;
    }

    .spinner-ring--inner {
      inset: 6px;
      border-top-color: transparent;
      border-right-color: transparent;
      border-bottom-color: rgba(124, 58, 237, 0.6);
      border-left-color: rgba(124, 58, 237, 0.2);
      animation-duration: 1.3s;
      animation-direction: reverse;
    }

    .spinner--sm .spinner-ring--inner { inset: 5px; }
    .spinner--lg .spinner-ring--inner { inset: 10px; }

    @keyframes spinner-rotate {
      to { transform: rotate(360deg); }
    }
  `],
})
export class SpinnerComponent {
  @Input() size: 'sm' | 'md' | 'lg' = 'md';
  @Input() minHeight = '120px';
}
