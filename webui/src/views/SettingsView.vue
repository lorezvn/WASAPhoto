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
        }
    }
}
</script>

<template>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">Settings</h1>
	</div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <form @submit.prevent="changeUsername">
        <label for="newUsername">Change username</label>
        <input type="text" v-model="newUsername"/>
        <button type="submit" class="btn btn-primary">Confirm</button>
    </form>

</template>

<style>

</style>