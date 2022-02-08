import Vue from 'vue'
import App from './App.vue'
import {BootstrapVue, IconsPlugin} from 'bootstrap-vue'
import './assets/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'dropzone/dist/dropzone.css'

Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

Vue.config.productionTip = false

new Vue({
    render: h => h(App),
}).$mount('#app')
