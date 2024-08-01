
<script>
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
		async likePost(photo) {
			// this.loading = true;
			this.errormsg = null;
			try {
				let liked = this.isLiked(photo);
				if (!liked) {
					await this.$axios.put(`/users/${photo.userID}/photos/${photo.photoID}/likes/${this.userID}`);
					photo.likes.push({ userID: this.userID });
				} else {
					await this.$axios.delete(`/users/${photo.userID}/photos/${photo.photoID}/likes/${this.userID}`);
					photo.likes = photo.likes.filter(user => user.userID != this.userID);
				}

			} catch (e) {
				this.errormsg = e.toString();
			}
			// this.loading = false;
		},
		isLiked(photo) {
			return photo.likes.some(user => user.userID == this.userID);
		},
		visitProfile(userID) {
			this.$router.push("/users/"+userID+"/profile")
		},
		formatDate(inputDate) {
			const options = { year: 'numeric', month: 'short', day: 'numeric', hour: 'numeric', minute: 'numeric'};
  			const formattedDate = new Date(inputDate).toLocaleDateString('en-US', options);
  			return formattedDate;
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
						<div class="photo-container">
							<img :src="'data:image/jpeg;base64,' + photo.image" alt="Image">
							<div id="like-counter" :class="['btn btn-sm', isLiked(photo) ? 'btn-outline-danger' : 'btn-outline-secondary']" @click="likePost(photo)">
								{{ photo.likes.length }}
								<svg class="feather" :class="{'liked': isLiked(photo)}">
									<use href="/feather-sprite-v4.29.0.svg#heart"/>
								</svg>
							</div>
							<button id="comment-counter" class="btn btn-outline-secondary btn-sm">
								{{ photo.comments.length }}
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
							</button>
							<p id="photo-date" class="photo-text"> {{ formatDate(photo.date) }} </p>
						</div>
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

	.feather.liked use {
		fill: red;
	}
</style>
