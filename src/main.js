import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'

Vue.use(VueRouter)

import TplAbout from './components/about.vue'
import TplBooks from './components/books.vue'

Vue.config.productionTip = false

// 일렉트론인 경우 히스토리모드 해제
let histroyMode = "history"
var userAgent = navigator.userAgent.toLowerCase();
if (userAgent.indexOf(' electron/') > -1) {
  // Electron-specific code
  histroyMode = ""
}

const router = new VueRouter({
  mode: histroyMode,
  routes: [
    { path: '/', component: TplAbout },
    { path: '/books', component: TplBooks }
  ]
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
