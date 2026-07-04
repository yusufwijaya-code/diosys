import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ViewportScroller } from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet],
  template: '<router-outlet />',
})
export class AppComponent {
  constructor(scroller: ViewportScroller) {
    scroller.setOffset([0, 88]);
  }
}
