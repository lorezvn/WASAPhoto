<script>

export default {
    data: function() {
        return {
            errormsg: null,
            loading: false,
            newUsername: ""
        }
    },
    methods: {
        async changeUsername() {
            // this.loading = true;
			this.errormsg = null;
			try {
                let userID = localStorage.getItem('token');
				const response = await this.$axios.put("/users/"+userID+"/username", {
                    username: this.newUsername
                });
                this.$router.push("/users/"+userID+"/profile");

			} catch (e) {
				this.errormsg = e.response.data.message;
			}
			//this.loading = false;
        },
        invalidUsername(username) {
			return username.length < 3 || username.length > 16 || username.trim().length < 3
		}
    }
}
</script>

<template>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">Settings</h1>
	</div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <form class="change-username-section" @submit.prevent="changeUsername">
        <h4 class="d-flex justify-content-center align-items-center">Change username</h4>
		<div class="d-flex justify-content-center align-items-center">
            <div class="form-group mb-2 col-4">
                <label class="mb-2">Username</label>
                <input type="text"
                    v-model="newUsername"
                    :class="['form-control', invalidUsername(newUsername) ? '' : 'is-valid']" 
                    placeholder="Enter new username" 
                    maxlength="16">
                <div v-if="!invalidUsername(newUsername)" class="valid-feedback">
                    Looks good!
                </div>
                <small v-else class="form-text text-muted">Must be 3-16 characters long.</small>
            </div>
            <div>
                <button type="submit" class="btn btn-primary"
                :disabled="invalidUsername(newUsername)">Submit</button>
            </div>
		</div>
	</form>
</template>

<style>

.change-username-section {
    margin-top: 20px;
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 4px;
}

</style>