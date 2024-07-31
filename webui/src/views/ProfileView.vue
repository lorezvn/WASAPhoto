
<script>
import LoadingSpinner from '../components/LoadingSpinner.vue';

export default {
	data: function() {
		return {
			errormsg: null,
			loading: true,
			username: "",
			photos: [],
			numberOfPhotos: 0,
			followers: 0,
			following: 0,
			userID: null,
			isFollowing: false,
			isBanning: false,
		}
	},
	computed: {
		owner() {
			return localStorage.getItem('token') === this.userID;
		}
  	},
	methods: {
		async getUserProfile() {

			if (this.userID === undefined){
                return
            }

			this.errormsg = null;
			
			try {
				let response = await this.$axios.get(`/users/${this.userID}/profile`);
				
				this.username = response.data.username;
				this.photos = response.data.photos;
				this.numberOfPhotos = response.data.photos.length;
				this.followers = response.data.followers;
				this.following = response.data.following;
				this.isFollowing = this.followers.some(follower => follower.userID == localStorage.getItem('token'));

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async deletePhoto(photoID) {
			try {
				await this.$axios.delete(`/users/${localStorage.getItem('token')}/photos/${photoID}`);
				this.getUserProfile();
			} catch(e) {
				this.errormsg = e.toString();
			}
		},
		async followUser() {
			try {
				if (!this.isFollowing) {
					await this.$axios.put(`/users/${localStorage.getItem('token')}/follow/${this.userID}`);
				} else {
					await this.$axios.delete(`/users/${localStorage.getItem('token')}/follow/${this.userID}`);
				}
				this.isFollowing = !this.isFollowing
				this.getUserProfile();
			} catch(e) {
				this.errormsg = e.toString();
			}
		},
		async banUser() {
			try {
				if (!this.isBanning) {
					await this.$axios.put(`/users/${localStorage.getItem('token')}/ban/${this.userID}`);
				} else {
					await this.$axios.delete(`/users/${localStorage.getItem('token')}/ban/${this.userID}`);
				}
				this.isBanning = !this.isBanning
				this.getUserProfile();
			} catch(e) {
				this.errormsg = e.toString();
			}
		},
		async likePost(photo) {
			this.errormsg = null;
			try {
				let liked = this.isLiked(photo);
				if (!liked) {
					await this.$axios.put(`/users/${photo.userID}/photos/${photo.photoID}/likes/${localStorage.getItem('token')}`);
				} else {
					await this.$axios.delete(`/users/${photo.userID}/photos/${photo.photoID}/likes/${localStorage.getItem('token')}`);
				}
				this.getUserProfile();

			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		isLiked(photo) {
			return photo.likes.some(user => user.userID == localStorage.getItem('token'));
		},
		goToUpload() {
			this.$router.push("/upload");
		},
		goToSettings() {
			this.$router.push("/users/"+this.userID+"/profile/settings")
		},
		formatDate(inputDate) {
			const options = { year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric'};
  			const formattedDate = new Date(inputDate).toLocaleDateString('en-US', options);
  			return formattedDate;
		}
	},
	watch: {
        '$route.params.userID': {
            handler(newUserID) {
                if (newUserID !== this.userID) {
                    this.userID = newUserID;
                    this.getUserProfile();
                }
            },
        }
    },

    async mounted() {
		this.userID = this.$route.params.userID;
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
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#camera"/></svg>
					Upload
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
		<div>
			<div class="d-flex flex-md-nowrap align-items-left">
				<div id="counter">
					<div class="follow-number">
						{{ numberOfPhotos }}
					</div>
					<div class="follow-text">
						photos
					</div>
				</div>
				<div id="counter">
					<div class="follow-number">
						{{ followers.length }}
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

			<div v-if="!owner" style="margin-left: -910px;" class="btn-toolbar justify-content-center mt-3">
				<div class="btn-group me-2">
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
					<button v-if="!isBanning" class="btn btn-primary btn-sm" @click="follofwUser">
						Ban
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#lock"/></svg>
					</button>
					<button v-if="isBanning" class="btn btn-secondary btn-sm" @click="followUser">
						Unban
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#unlock"/></svg>
					</button>
				</div>
			</div>
		</div>

		<div style="margin-top: 20px;">
			<div v-if="photos.length > 0">
				<h3>Photos</h3>
				<ul class="photo-list">
					<li v-for="photo in photos">
						<div class="photo-container">
							<img :src="'data:image/jpeg;base64,'+photo.image" alt="Image">
							<div v-if="owner" class="delete-button-container">
								<button type="button" class="btn btn-sm btn-danger" @click="deletePhoto(photo.photoID)">
								Delete
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
							</button>
							</div>
							<div class="d-flex align-items-center justify-content-between">
								<!-- Display Like and Comment buttons if not the owner -->
								<div v-if="!owner">
									<div id="like-counter" class="btn btn-outline-secondary btn-sm me" @click="likePost(photo)">
										{{ photo.likes.length }}
										<svg class="feather" :class="{'liked': isLiked(photo)}">
											<use href="/feather-sprite-v4.29.0.svg#heart"/>
										</svg>
									</div>
									
									<button id="comment-counter" class="btn btn-outline-secondary btn-sm">
										{{ photo.comments.length }}
										<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
									</button>
								</div>

								<!-- Display only Comment button if the owner -->
								<div v-else>
									<button style="left:8px;" id="comment-counter" class="btn btn-outline-secondary btn-sm">
										{{ photo.comments.length }}
										<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
									</button>
								</div>
							</div>
							<p id="photo-date" class="photo-text"> {{ formatDate(photo.date) }} </p>
						</div>
					</li>
				</ul>	
			</div>
			<div v-else>
				<h5>No photos posted</h5>
			</div>
		</div>
	</LoadingSpinner>
</template>

<style>

	#counter {
		width: 80px;
	}

	.follow-number {
		font-size: 25px;
		font-weight: bold;
		text-align: center;
	}

	.follow-text {
		font-size: 15px;
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

	.photo-container {
		width: 100%;
		height: 300px; /* Imposta un'altezza fissa per mantenere le dimensioni uniformi */
		border-radius: 8px; /* Angoli arrotondati per l'immagine */
		overflow: hidden; /* Nasconde le parti dell'immagine che escono dalla box */
		margin-bottom: 45px;
	}

	.photo-container img {
		width: 100%;
		height: 100%;
		object-fit: cover; /* Mantiene le proporzioni senza distorcere l'immagine */
		border-radius: 8px; /* Angoli arrotondati per l'immagine */
	}

	.photo-text{
		position: absolute;
		padding: 4px 8px; /* Spazio interno per migliorare l'aspetto */
		border-radius: 4px; /* Angoli arrotondati per il box della data */
		font-size: 14px; /* Dimensione del testo */
		color: #333; /* Colore del testo */
		font-family: Arial, Helvetica, sans-serif;
	}

	#photo-date {
		bottom: -8px; /* Distanza dal bordo inferiore */
		right: 8px;
	}

	
	#like-counter {
		position: absolute;
		bottom: 8px; /* Distanza dal bordo inferiore */
		left: 8px; /* Distanza dal bordo destro */
	}

	#comment-counter {
		position: absolute;
		bottom: 8px; /* Distanza dal bordo inferiore */
		left: 60px;
	}

	.delete-button-container {
		position: absolute;
		top: 8px;
		right: 8px;
		border-radius: 4px; /* Angoli arrotondati per il contenitore del pulsante */
	}

	

</style>
