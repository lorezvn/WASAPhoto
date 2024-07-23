import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import HomeView from '../views/HomeView.vue'
import ProfileView from '../views/ProfileView.vue'
import SearchView from '../views/SearchView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/login'},
		{path: '/login', component: LoginView},
		{path: '/home', component: HomeView},
		{path: '/users/:userID/profile', component: ProfileView},
		{path: '/search', component: SearchView},
	]
})

export default router
