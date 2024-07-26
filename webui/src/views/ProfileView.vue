
<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: "",
			photos: [],
			numberOfPhotos: 0,
			followers: 0,
			following: 0,
			userID: null
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

			// this.loading = true;
			this.errormsg = null;
			
			try {
				let response = await this.$axios.get(`/users/${this.userID}/profile`);

				this.username = response.data.username;
				this.photos = response.data.photos;
				this.numberOfPhotos = response.data.photos.length;
				this.followers = response.data.followers.length;
				this.following = response.data.following.length;

			} catch (e) {
				this.errormsg = e.toString();
			}
			//this.loading = false;
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
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">{{ username }}</h1>
			<div v-if="owner" class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="goToSettings">
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings"/></svg>
						Settings
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		
		<div class="d-flex flex-md-nowrap align-items-center">
			<div id="counter">
				<div class="follow-number">
					{{ numberOfPhotos }}
				</div>
				<div class="follow-text">
					Photos
				</div>
			</div>
			<div id="counter">
				<div class="follow-number">
					{{ followers }}
				</div>
				<div class="follow-text">
					Followers
				</div>
			</div>
			<div id="counter">
				<div class="follow-number">
					{{ following }}
				</div>
				<div class="follow-text">
					Following
				</div>
			</div>
		</div>

		<div style="margin-top: 20px;">
			<div v-if="photos.length > 0">
				<h5>Photos</h5>
				<ul class="photo-list">
					<li v-for="photo in photos">
						<div class="photo-container">
							<img :src="'data:image/jpeg;base64,'+photo.image">
							{{ formatDate(photo.date) }}
						</div>
					</li>
				</ul>	
			</div>
			<div v-else>
				<h5>No photos posted</h5>
			</div>
		</div>
	</div>
</template>

<style>

	#counter {
		width: 100px;
	}

	.follow-number {
		font-size: 23px;
		font-weight: bold;
		text-align: center;
	}

	.follow-text {
		font-size: 15px;
		color:gray;
		text-align: center;
	}

	.photo-container {
		max-height: 300px;
        max-width: 300px;
		align-items: center;
		text-align: end;
    }

    .photo-container img {
		max-height: 100%;
        max-width: 100%;  
    }

	.photo-list {
  		list-style-type: none;
	}

	.photo-list li {
		margin-top: 20px;
	}

</style>
