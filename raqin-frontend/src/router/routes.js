
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: 'home', alias: '/', component: () => import('pages/home/home.vue') },
      { path: 'books', alias: '/', component: () => import('pages/book/book.vue') },
      { path: 'pages', component: () => import('pages/my-pages/page.vue') },
      { path: 'profile', component: () => import('pages/profile/profile.vue') },
      { path: 'contribute', component: () => import('pages/contribute/contribute.vue') },
    ]
  },
  { path: '/register', component: () => import('pages/auth/register/register.vue') },
  { path: '/login', component: () => import('pages/auth/login/login.vue') },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
