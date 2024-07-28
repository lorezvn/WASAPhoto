<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			usernameQuery: "",
			users: [],
			visitedUsers: []
		}
	},
	methods: {
		async searchUsers() {
			// this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get(`/users?username=${this.usernameQuery}`);
				this.users = response.data;

			} catch (e) {
				this.errormsg = e.toString();
			}
			// this.loading = false;
		},
		saveVisitedUser(user) {

			let searches = JSON.parse(localStorage.getItem(localStorage.getItem('token')+'visitedUsers')) || [];

			if (!searches.some(search => search.userID === user.userID)) {

				searches.unshift(user);
				
				if (searches.length > 5) {
					searches.pop();
				}

				localStorage.setItem(localStorage.getItem('token')+'visitedUsers', JSON.stringify(searches));
				this.visitedUsers = searches;
			}
		},
		async loadVisitedUsers() {
			this.visitedUsers = JSON.parse(localStorage.getItem(localStorage.getItem('token')+'visitedUsers')) || [];
			const userPromises = this.visitedUsers.map(async user => {
				// Keep track of the usernames of recent searched users (they might change)
				try {
					const response = await this.$axios.get(`/users/${user.userID}/profile`);
					user.username = response.data.username;
					return user;
				} catch (e) {
					return null;
				}
    		});
    		const resolvedUsers = await Promise.all(userPromises);
			this.visitedUsers = resolvedUsers.filter(user => user !== null);

			localStorage.setItem(localStorage.getItem('token')+'visitedUsers', JSON.stringify(this.visitedUsers));
		},

		visitProfile(user) {
			this.saveVisitedUser(user);
			this.$router.replace("/users/"+user.userID+"/profile")
		}
	},
	mounted() {
		this.loadVisitedUsers();
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Search Users</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <form @submit.prevent="searchUsers">
            <input type="text" v-model="usernameQuery" @input="searchUsers" placeholder="Search for users..."/>
            <button type="submit" class="btn btn btn-primary">
                Search
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
            </button>
        </form> 

		<ul v-if=users.length class="list-group">
			<li v-for="user in users" 
				:key="user.userID" class="list-group-item d-flex justify-content-between align-items-center" 
				@click="visitProfile(user)" 
				style="cursor: pointer;">
				{{ user.username }}
			</li>
    	</ul>

		<p v-else-if="usernameQuery">No users found.</p>

		<div v-else>
			<ul class="list-group">
				<li v-for="user in visitedUsers" 
					:key="user.userID" 
					class="list-group-item" 
					@click="visitProfile(user)">
					{{ user.username }}
				</li>
			</ul>
		</div>
    </div>

</template>

<style>
	.list-group-item {
		cursor: pointer;
		transition: background-color 0.3s;
	}

	.list-group-item:hover {
		background-color: #e0e0e0;
	}

</style>
