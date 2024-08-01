<script>
import LoadingSpinner from '../components/LoadingSpinner.vue';


export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
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
    <div class="d-flex justify-content-center flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">Welcome to WASAPhoto</h1>
	</div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

	<LoadingSpinner></LoadingSpinner>

	<form @submit.prevent="login">
		<div class="container d-flex align-items-center justify-content-center">
			<div class="form-group mb-2 col-4">
				<label class="mb-2">Username</label>
				<input type="text"
					v-model="username"
					class="form-control"
					placeholder="Enter username" 
					maxlength="16">
				<small class="form-text text-muted">Must be 3-16 characters long.</small>
			</div>
			<div>
				<button type="submit" class="btn btn-primary"
				:disabled="invalidUsername(username)">Submit</button>
			</div>
		</div>
	</form>
    
</template>

<style>

	body {
        font-family: 'Arial', sans-serif;
        background-color: #f8f9fa;
        color: #000000;
    }

    .btn-primary {
		margin-left: 5px;
    }
	
</style>
