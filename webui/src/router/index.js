import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import PostDetailPage from '../components/PostDetailPage.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/home', component: HomeView},
		{path: '/link2', component: HomeView},
		{path: '/profiles/:user_id/profile', component: ProfileView},
		{path: '/session', component:LoginView},
		{path: '/', redirect: '/session'},
		{path: '/profiles/:user_id/photos/:photo_id', component: PostDetailPage},
	]
})

export default router
