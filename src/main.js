import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'

Vue.use(VueRouter)

import TplAbout from './components/about.vue'
import TplBooks from './components/books.vue'

const router = new VueRouter({
  mode: 'history',
  routes: [
      { path: '/', component: TplAbout },
      { path: '/books', component: TplBooks }
  ]
})

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
