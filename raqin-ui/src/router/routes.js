
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: 'home', component: () => import('pages/Index.vue') },
      { path: 'project', component: () => import('pages/project.vue') },
      { path: 'contribute', component: () => import('pages/contribute.vue') },
      { path: 'profile', component: () => import('pages/profile.vue') },
      { path: 'preferences', component: () => import('pages/preferences.vue') }
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
