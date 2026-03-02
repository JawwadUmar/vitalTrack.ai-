import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailsView } from './details-view';

describe('DetailsView', () => {
  let component: DetailsView;
  let fixture: ComponentFixture<DetailsView>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DetailsView],
    }).compileComponents();

    fixture = TestBed.createComponent(DetailsView);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
