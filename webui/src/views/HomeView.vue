
<script>
import Photo from '../components/Photo.vue';

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: [],
			userID: localStorage.getItem('token')
		}
	},
	methods: {
		async getStream() {
			this.errormsg = null;
			this.loading = true;
			try {
				let response = await this.$axios.get(`/users/${this.userID}/stream`);
				this.stream = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		visitProfile(userID) {
			this.$router.push("/users/"+userID+"/profile")
		},
	},
	async mounted() {
		await this.getStream()
	}
}
</script>

<template>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<LoadingSpinner :loading="loading">
			<div v-if="stream.length > 0">
				<h5>Photos posted by the users you follow</h5>
				<ul style="margin-top: 20px;" class="photo-list">
					<li v-for="photo in stream" :key="photo.photoID">
						<button @click="visitProfile(photo.userID)" class="btn btn-sm btn-secondary position-absolute photo-owner">
							<strong>@{{ photo.username }}</strong>
						</button>
						<Photo
							:photo="photo"
							:owner="false">
						</Photo>
					</li>
				</ul>
			</div>
			<div v-else>
				<h5>Your feed is empty! Follow someone!</h5>
			</div>
		</LoadingSpinner>
</template>

<style>
	.photo-owner {
		top: 8px;
		left: 8px;
		background-color: rgb(255, 255, 255);
		color:black
	}
</style>
