
<script>
import LoadingSpinner from '../components/LoadingSpinner.vue';
import Photo from '../components/Photo.vue';

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: "",
			currentUserID: null,
			photos: [],
			followers: [],
			followersCount: 0,
			following: [],
			isFollowing: false,
			isBanning: false,
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

				} else {
					await this.$axios.delete(`/users/${localStorage.getItem('token')}/follow/${this.currentUserID}`);
					this.followersCount -= 1
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
		<div style="float:left;" class="d-flex-column mb-4">
			<div class="d-flex justify-content-center">
				<div id="counter">
					<div class="follow-number">
						{{ photos.length }}
					</div>
					<div class="follow-text">
						photos
					</div>
				</div>
				<div id="counter">
					<div class="follow-number">
						{{ followersCount}}
					</div>
					<div class="follow-text">
						followers
					</div>
				</div>
				<div id="counter">
					<div class="follow-number">
						{{ following.length }}
					</div>
					<div class="follow-text">
						following
					</div>
				</div>
			</div>

			<div v-if="!owner" class="btn-toolbar justify-content-center mt-3">
				<div v-if="!isBanning" class="btn-group me-2">
					<button v-if="!isFollowing" class="btn btn-primary btn-sm" @click="followUser">
						Follow
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user-plus"/></svg>
					</button>
					<button v-if="isFollowing" class="btn btn-secondary btn-sm" @click="followUser">
						Unfollow
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user-minus"/></svg>
					</button>
				</div>
				<div class="btn-group me-2">
					<button v-if="!isBanning" class="btn btn-primary btn-sm" @click="banUser">
						Ban
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#lock"/></svg>
					</button>
					<button v-if="isBanning" class="btn btn-secondary btn-sm" @click="banUser">
						Unban
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#unlock"/></svg>
					</button>
				</div>
			</div>
		</div>

		<div v-if="!isBanning" style="clear:both;" class="mt-4">
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
			<div v-else class="d-flex flex-column align-items-center justify-content-center pt-3 pb-2 border-top" style="height:40vh;">
            	<svg id="icon-no-post" class="feather"><use href="/feather-sprite-v4.29.0.svg#camera-off"/></svg>
				<p style="color:gray; font-weight: bold;">No photos</p>
			</div>
		</div>
	</LoadingSpinner>
</template>

<style>

	#counter {
		width: 80px;
		text-align: center;
	}

	.follow-number {
		font-size: 20px;
		font-weight: bold;
		text-align: center;
	}

	.follow-text {
		font-size: 14px;
		color:gray;
		text-align: center;
	}

	.photo-list {
		list-style-type: none;
		padding: 0;
		margin: 0;
		display: flex;
		gap: 32px; /* Spazio tra le foto */
		flex-wrap: wrap; /* Permette la visualizzazione in più righe, se necessario */
	}

	.photo-list li {
		width: 100%; /* Può essere modificato se si desidera un layout a colonne */
		max-width: 300px; /* Larghezza massima della box della foto */
		border-radius: 12px; /* Angoli arrotondati per la box */
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2); /* Ombra sotto la box */
		padding: 0; /* Rimuovi il padding interno per evitare spazi aggiuntivi */
		position: relative; /* Necessario per posizionare il testo all'interno della box */
	}

	#icon-no-post {
        color:gray;
        width: 70px; /* Imposta la larghezza dell'icona */
        height: 70px; /* Imposta l'altezza dell'icona */
        margin-bottom: 5px; /* Distanza tra l'icona e la scritta */
    }

</style>
