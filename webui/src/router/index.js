import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import HomeView from '../views/HomeView.vue'
import ProfileView from '../views/ProfileView.vue'
import SearchView from '../views/SearchView.vue'
import SettingsView from '../views/SettingsView.vue'
import UploadPhoto from '../views/UploadPhoto.vue'
import BannedView from '../views/BannedView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/login'},
		{path: '/login', component: LoginView},
		{path: '/home', component: HomeView},
		{path: '/users/:userID/profile', component: ProfileView},
		{path: '/search', component: SearchView},
		{path: '/users/:userID/profile/settings', component: SettingsView},
		{path: '/upload', component: UploadPhoto},
		{path: '/banned', component: BannedView}
	]
})

export default router
