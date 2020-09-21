import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import VueRouter from 'vue-router';
import App from './App.vue';
//import LaunchSudokuHome from './components/LaunchSudokuHome.vue';
import Settings from './components/Settings.vue';
import Home from './components/Home.vue';

Vue.use(VueRouter);

const routes = [
	{ path: '/settings', component: Settings },
	{ path: '/', component: Home },
	{ path: '/Home', component: Home },
];

const router = new VueRouter({
	routes,
	mode: 'history'
});

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
	new Vue({
		router,
		render: h => h(App)
	}).$mount('#app');
});
