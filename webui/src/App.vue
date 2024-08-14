<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data: function(){
		return {
			showNavBar: this.$route.path != '/login',
			userID: localStorage.getItem('token')
		}
	},
	watch: {
		$route(to) {
      		this.showNavBar = to.path !== '/login';
			this.userID = localStorage.getItem('token');
		},
	},
	methods: {
		logout() {
			localStorage.clear();
			this.$router.replace("/login");
		}
	},
}
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6">WASAPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
		<div v-if="showNavBar" class="btn-toolbar mb-2 mb-md-0">
			<div class="btn-group me-3">
				<button type="button" class="btn btn-sm btn-primary" @click="logout">
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
					Logout
				</button>
			</div>
		</div>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav v-if="showNavBar" id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/home" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								Home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/search" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
								Search
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink :to="`/users/${userID}/profile`" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
								Profile
							</RouterLink>
						</li>
					</ul>
				</div>
			</nav>

			<main v-if="showNavBar" class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>

			<main v-else>
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
</style>
