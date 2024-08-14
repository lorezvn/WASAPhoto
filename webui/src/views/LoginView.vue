<script>
import LoadingSpinner from '../components/LoadingSpinner.vue';


export default {
	data: function() {
		return {
			errormsg: null,
            username: "",
			userID: null,
		}
	},
	methods: {
		async login() {
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/session", {
                    username: this.username,
                });
				this.userID = response.data;
                localStorage.setItem('token', this.userID);
				localStorage.setItem('username', this.username);
                this.$router.replace("/home");

			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		invalidUsername(username) {
			return username.length < 3 || username.length > 16 || username.trim().length < 3
		}
	},
}
</script>

<template>
	<div class="login-container">
		<div class="login-box">
			<h3 class="text-center mb-4">Login</h3>
			
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

			<form @submit.prevent="login">
				<div class="form-group mb-2 col-auto">
					<label class="mb-2">Username</label>
					<input type="text"
						v-model="username"
						class="form-control"
						placeholder="Enter username" 
						maxlength="16">
					<small class="form-text text-muted">Must be 3-16 characters long.</small>
				</div>
				<div class="d-grid mt-3">
					<button type="submit" class="btn btn-primary"
						:disabled="invalidUsername(username)">
						Continue
					</button>
				</div>
			</form>
		</div>
	</div>
</template>

<style>

	.login-container {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100%;
		width: 100%;
		height: 80vh; 
	}

	.login-box {
		background-color: white;
		padding: 40px;
		box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.2);
		min-width: 450px;
	}
	
</style>
