<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			usernameQuery: "",
			users: [],
		}
	},
	methods: {
		async searchUsers() {
			// this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get(`/users/?username=${this.usernameQuery}`);
				this.users = response.data;

			} catch (e) {
				this.errormsg = e.toString();
			}
			// this.loading = false;
		},

		visitProfile(userID) {
			this.$router.push("/users/"+userID+"/profile")
		}
	},
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
			<div class="form-group col-auto">
				<div class="input-group mb-2">
					<div class="input-group-prepend">
						<div class="input-group-text">
							<svg style="height:24px;" class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
						</div>
					</div>
					<input type="text" v-model="usernameQuery" class="form-control" @input="searchUsers" placeholder="Search for users..."/>
				</div>	
			</div>
        </form> 

		<ul v-if=users.length class="list-group">
			<li v-for="user in users" 
				class="list-group-item d-flex justify-content-between align-items-center" 
				@click="visitProfile(user.userID)" 
				style="cursor: pointer;">
				{{ user.username }}
			</li>
    	</ul>

		<p v-else-if="usernameQuery">No users found.</p>
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

	.input-group-text {
		border-top-right-radius: 0px;
		border-bottom-right-radius: 0px;
	}

</style>
