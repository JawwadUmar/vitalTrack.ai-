import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AiAnalysisPage } from './ai-analysis-page';

describe('AiAnalysisPage', () => {
  let component: AiAnalysisPage;
  let fixture: ComponentFixture<AiAnalysisPage>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AiAnalysisPage],
    }).compileComponents();

    fixture = TestBed.createComponent(AiAnalysisPage);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
