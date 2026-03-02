import { Routes } from '@angular/router';
import { Login } from './features/auth/login/login';
import { Signup } from './features/auth/signup/signup';
import { CalendarDashboard } from './features/dashboard/calendar-dashboard/calendar-dashboard';
import { DetailsView } from './features/details/details-view/details-view';
import { AiAnalysisPage } from './features/ai-analysis/ai-analysis-page/ai-analysis-page';
import { authGuard } from './core/guards/auth.guard';

export const routes: Routes = [
  { path: '', redirectTo: '/login', pathMatch: 'full' },
  { path: 'login', component: Login },
  { path: 'signup', component: Signup },
  { path: 'dashboard', component: CalendarDashboard, canActivate: [authGuard] },
  { path: 'details', component: DetailsView, canActivate: [authGuard] },
  { path: 'analysis', component: AiAnalysisPage, canActivate: [authGuard] },
  { path: '**', redirectTo: '/login' }
];
