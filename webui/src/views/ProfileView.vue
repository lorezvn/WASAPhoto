
<script>
import LoadingSpinner from '../components/LoadingSpinner.vue';
import Photo from '../components/Photo.vue';
import UserModal from '../components/UserModal.vue';

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: "",
			currentUserID: null,

			photos: [],
			followers: [],
			following: [],
			followersCount: 0,

			isFollowing: false,
			isBanning: false,

			showModal: false,
			modalTitle: "",
			modalUsers: [],
		}
	},
	computed: {
		owner() {
			return localStorage.getItem('token') === this.currentUserID;
		}
  	},
	methods: {
		async getUserProfile() {

			if (this.currentUserID === undefined){
                return
            }

			this.errormsg = null;
			this.loading = true;
			
			try {
				let response = await this.$axios.get(`/users/${this.currentUserID}/profile`);
				
				this.username = response.data.username;
				this.photos = response.data.photos;
				this.followers = response.data.followers;
				this.following = response.data.following;
				this.followersCount = response.data.followers.length;
				this.isFollowing = this.followers.some(follower => follower.userID == localStorage.getItem('token'));

				await this.getBannedUsers();

			} catch (e) {
				if (e.response.status === 403) {
					this.$router.replace("/banned");
				}
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async getBannedUsers() {
			try {
				let response = await this.$axios.get(`/users/${localStorage.getItem('token')}/ban/`);
				this.bannedUsers = response.data;
				this.isBanning = this.bannedUsers.some(bannedUser => bannedUser.userID == this.currentUserID);
				
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async followUser() {
			try {
				if (!this.isFollowing) {
					await this.$axios.put(`/users/${localStorage.getItem('token')}/follow/${this.currentUserID}`);
					this.followersCount += 1
					
					// Add the current user's profile to the list of followers
					this.followers.push({
						userID: localStorage.getItem('token'),
						username: localStorage.getItem('username'), 
					});

				} else {
					await this.$axios.delete(`/users/${localStorage.getItem('token')}/follow/${this.currentUserID}`);
					this.followersCount -= 1

					// Remove the current user's profile from the list of followers
					this.followers = this.followers.filter(follower => follower.userID != localStorage.getItem('token'));
				}
				this.isFollowing = !this.isFollowing
			} catch(e) {
				this.errormsg = e.toString();
			}
		},
		async banUser() {
			try {
				if (!this.isBanning) {
					await this.$axios.put(`/users/${localStorage.getItem('token')}/ban/${this.currentUserID}`);
					this.isFollowing = false;
				} else {
					await this.$axios.delete(`/users/${localStorage.getItem('token')}/ban/${this.currentUserID}`);
					this.getUserProfile();
				}
				this.isBanning = !this.isBanning
			} catch(e) {
				this.errormsg = e.toString();
			}
		},
		showFollowersModal() {
			this.modalTitle = "Followers"
			this.modalUsers = this.followers;
			this.showModal = true;
			document.body.classList.add("modal-open");
		},
		showFollowingModal() {
			this.modalTitle = "Following";
			this.modalUsers = this.following;
			this.showModal = true;
			document.body.classList.add("modal-open");
		},
		closeModal() {
            this.showModal = false;
            document.body.classList.remove("modal-open");
        },
		handlePhotoDeleted(photoID) {
			this.photos = this.photos.filter(photo => photo.photoID != photoID);
		},
		goToUpload() {
			this.$router.push("/upload");
		},
		goToSettings() {
			this.$router.push("/users/"+this.currentUserID+"/profile/settings")
		},
	},
	watch: {
        '$route.params.userID': {
            handler(newUserID) {
                if (newUserID !== this.currentUserID) {
                    this.currentUserID = newUserID;
                    this.getUserProfile();
                }
            },
        }
    },

    async mounted() {
		this.currentUserID = this.$route.params.userID;
        await this.getUserProfile();
    }
}
</script>

<template>
	<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">{{ username }}</h1>
		<div v-if="owner" class="btn-toolbar mb-2 mb-md-0">
			<div class="btn-group me-2">
				<button type="button" class="btn btn-sm btn-outline-primary" @click="goToUpload">
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#upload"/></svg>
					Upload Photo
				</button>
				<button type="button" class="btn btn-sm btn-outline-secondary" @click="goToSettings">
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings"/></svg>
					Settings
				</button>
			</div>
		</div>
	</div>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

	<LoadingSpinner :loading="loading">	
		<div class="profile-info mb-4">
			<div class="d-flex justify-content-center text-center">
				<div class="profile-counter">
					<div class="counter-number">{{ photos.length }}</div>
					<div class="counter-text">photos</div>
				</div>
				<div class="profile-counter clickable" @click="showFollowersModal">
					<div class="counter-number">{{ followersCount}}</div>
					<div class="counter-text">followers</div>
				</div>
				<div class="profile-counter clickable" @click="showFollowingModal">
					<div class="counter-number">{{ following.length }}</div>
					<div class="counter-text">following</div>
				</div>
			</div>

			<UserModal
				:isVisible="showModal" 
				:title="modalTitle" 
				:users="modalUsers" 
				@close="closeModal">
			</UserModal>

			<div v-if="!owner" class="btn-toolbar justify-content-center mt-3">
				<div class="btn-group me-3">
					<div v-if="!isBanning"
						:class="['btn btn-sm', isFollowing ? 'btn-secondary' : 'btn-primary']" 
						@click="followUser"
					>
						{{ isFollowing ? 'Unfollow' : 'Follow' }}
						<svg class="feather">
							<use :href="'/feather-sprite-v4.29.0.svg#' + (isFollowing ? 'user-minus' : 'user-plus')"/>
						</svg>
					</div>
				</div>
				<div class="btn-group me-3">
					<div 
						:class="['btn btn-sm', isBanning ? 'btn-secondary' : 'btn-primary']" 
						@click="banUser"
					>
						{{ isBanning ? 'Unban' : 'Ban' }}
						<svg class="feather">
							<use :href="'/feather-sprite-v4.29.0.svg#' + (isBanning ? 'unlock' : 'lock')"/>
						</svg>
					</div>
				</div>
			</div>
		</div>

		<div v-if="!isBanning" class="photos-section mt-4">
			<div v-if="photos.length > 0">
				<ul class="photo-list">
					<li v-for="photo in photos">
						<Photo
							:photo="photo"
							:owner="owner"
							@photoDeleted="handlePhotoDeleted">
						</Photo>
					</li>
				</ul>	
			</div>
			<div v-else class="no-photos d-flex flex-column align-items-center justify-content-center pt-3 pb-2 border-top">
            	<svg id="no-photos-icon" class="feather"><use href="/feather-sprite-v4.29.0.svg#camera-off"/></svg>
				<p class="no-photos-text">No photos</p>
			</div>
		</div>
	</LoadingSpinner>
</template>

<style>

	.profile-info {
		float: left;
	}

	.profile-counter {
		width: 70px;
		margin: 0 10px;
		text-align: center;
		padding: 10px;
	}

	.counter-number {
		font-size: 20px;
		font-weight: bold;
	}

	.counter-text {
		font-size: 14px;
		color: gray;
	}

	.photos-section {
		clear: both;
	}

	.photo-list {
		list-style-type: none;
		padding: 0;
		margin: 0;
		display: flex;
		gap: 32px;
		flex-wrap: wrap;
	}

	.photo-list li {
		width: 100%;
		max-width: 300px;
		border-radius: 12px;
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
		padding: 0;
		position: relative;
	}

	.no-photos {
		height: 40vh;
	}

	#no-photos-icon{
		color: gray;
		width: 70px;
		height: 70px;
		margin-bottom: 5px;
	}

	.no-photos-text {
		color: gray;
		font-weight: bold;
	}

</style>
