<script>

export default {
    data: function() {
        return {
            errormsg: null,
            loading: false,
            newUsername: "",
            errorStatus: null
        }
    },
    watch: {
        newUsername() {
            this.clearErrors();
        }
    },
    computed: {
        validationClass() {
            if (this.newUsername === "") return "";
            if (this.errormsg || this.invalidUsername(this.newUsername)) return "is-invalid";
            return "is-valid";
        }
    },
    methods: {
        async changeUsername() {
            // this.loading = true;
			this.errormsg = null;
            this.errorStatus = null;
			try {
                let userID = localStorage.getItem('token');
				this.response = await this.$axios.put("/users/"+userID+"/username", {
                    username: this.newUsername
                });

                localStorage.setItem('username', this.newUsername);
                this.$router.push("/users/"+userID+"/profile");

			} catch (e) {
                this.errorStatus = e.response.status
                this.errormsg = e.toString();
			}
        },
        invalidUsername(username) {
			return username.trim().length < 3 || username.trim().length > 16
		},
        clearErrors() {
            this.errormsg = null;
            this.errorStatus = null;
        }
    }
}
</script>

<template>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">Settings</h1>
    </div>

    <ErrorMsg v-if="errormsg && errorStatus !== 409" :msg="errormsg"></ErrorMsg>

    <form class="change-username-section" @submit.prevent="changeUsername">
        <h4 class="d-flex justify-content-center align-items-center">Change username</h4>
        <div class="d-flex justify-content-center align-items-center">
            <div class="form-group mb-2 col-4">
                <label class="mb-2">Username</label>
                <input type="text"
                    v-model="newUsername"
                    :class="['form-control', validationClass]" 
                    placeholder="Enter new username">
                <div v-if="invalidUsername(newUsername)" 
                    :class="[newUsername === '' ? 'form-text text-muted' : 'invalid-feedback']"> 
                    Must be 3-16 characters long.
                </div>
                <div v-else-if="errorStatus === 409" class="invalid-feedback">Username chosen already exists</div>
                <div v-else class="valid-feedback">Looks good!</div>
            </div>
            <div>
                <button type="submit" class="btn btn-primary ms-1" :disabled="invalidUsername(newUsername)">Submit</button>
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