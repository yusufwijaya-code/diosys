import { Routes } from '@angular/router';

import { authGuard } from './core/guards/auth.guard';

export const routes: Routes = [
  {
    path: 'admin/login',
    loadComponent: () =>
      import('./pages/admin/admin-login/admin-login.component').then((m) => m.AdminLoginComponent),
  },
  {
    path: 'admin',
    canActivate: [authGuard],
    loadComponent: () =>
      import('./pages/admin/admin-layout/admin-layout.component').then((m) => m.AdminLayoutComponent),
    children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      {
        path: 'dashboard',
        loadComponent: () =>
          import('./pages/admin/admin-dashboard/admin-dashboard.component').then((m) => m.AdminDashboardComponent),
      },
      {
        path: 'developers',
        loadComponent: () =>
          import('./pages/admin/admin-developers/admin-developers.component').then((m) => m.AdminDevelopersComponent),
      },
      {
        path: 'developers/:userID/portfolio',
        loadComponent: () =>
          import('./pages/admin/admin-portfolio/admin-portfolio.component').then((m) => m.AdminPortfolioComponent),
      },
      {
        path: 'projects',
        loadComponent: () =>
          import('./pages/admin/admin-projects/admin-projects.component').then((m) => m.AdminProjectsComponent),
      },
      {
        path: 'professional-projects',
        loadComponent: () =>
          import('./pages/admin/admin-professional-projects/admin-professional-projects.component').then(
            (m) => m.AdminProfessionalProjectsComponent
          ),
      },
      {
        path: 'inbox',
        loadComponent: () =>
          import('./pages/admin/admin-inbox/admin-inbox.component').then((m) => m.AdminInboxComponent),
      },
      {
        path: 'services',
        loadComponent: () =>
          import('./pages/admin/admin-services/admin-services.component').then((m) => m.AdminServicesComponent),
      },
      {
        path: 'pricing',
        loadComponent: () =>
          import('./pages/admin/admin-pricing/admin-pricing.component').then((m) => m.AdminPricingComponent),
      },
      {
        path: 'settings',
        loadComponent: () =>
          import('./pages/admin/admin-settings/admin-settings.component').then((m) => m.AdminSettingsComponent),
      },
    ],
  },
  {
    path: '',
    loadComponent: () =>
      import('./pages/public/public-layout/public-layout.component').then((m) => m.PublicLayoutComponent),
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./pages/public/home/home.component').then((m) => m.HomeComponent),
      },
      {
        path: ':username/projects/:projectID',
        loadComponent: () =>
          import('./pages/public/project-detail/project-detail.component').then(
            (m) => m.ProjectDetailComponent
          ),
      },
      {
        path: ':username/professional-projects/:projectID',
        loadComponent: () =>
          import('./pages/public/professional-project-detail/professional-project-detail.component').then(
            (m) => m.ProfessionalProjectDetailComponent
          ),
      },
      {
        path: ':username',
        loadComponent: () =>
          import('./pages/public/developer-profile/developer-profile.component').then(
            (m) => m.DeveloperProfileComponent
          ),
      },
    ],
  },
  { path: '**', redirectTo: '' },
];
