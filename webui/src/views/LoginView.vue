<script>
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
			// this.loading = true;
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
			// this.loading = false;
		},
	},
}
</script>

<template>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">Login</h1>
	</div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <form @submit.prevent="login">
        <label for="username">Insert your username</label>
        <input type="text" v-model="username"/>
        <button type="submit" class="btn btn-primary">Login</button>
    </form> 
    
</template>

<style>

	body {
        font-family: 'Arial', sans-serif;
        background-color: #f8f9fa;
        color: #000000;
    }

	label {
		margin-right: 5px;
	}

    .btn-primary {
		margin-left: 5px;
    }
	
</style>
