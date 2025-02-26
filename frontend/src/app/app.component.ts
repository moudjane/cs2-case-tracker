import { Component } from '@angular/core';
import { PriceTrackerComponent } from './components/price-tracker/price-tracker.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [PriceTrackerComponent],
  template: `<app-price-tracker></app-price-tracker>`,
  styleUrls: ['./app.component.css']
})
export class AppComponent {}
