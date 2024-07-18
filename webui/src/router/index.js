import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/home', component: HomeView},
		{path: '/link1', component: HomeView},
		{path: '/link2', component: ProfileView},
		{path: '/some/:id/link', component: HomeView},
		{path: '/session', component:LoginView},
		{path: '/', redirect: '/session'},
	]
})

export default router
