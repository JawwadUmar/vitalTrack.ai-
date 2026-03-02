import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CalendarDashboard } from './calendar-dashboard';

describe('CalendarDashboard', () => {
  let component: CalendarDashboard;
  let fixture: ComponentFixture<CalendarDashboard>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CalendarDashboard],
    }).compileComponents();

    fixture = TestBed.createComponent(CalendarDashboard);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
