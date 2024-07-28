
<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: [],
		}
	},
	methods: {
		async getStream() {
			// this.loading = true;
			this.errormsg = null;
			try {
				let userID = localStorage.getItem('token');
				let response = await this.$axios.get(`/users/${userID}/stream`);
				this.stream = response.data;

			} catch (e) {
				this.errormsg = e.toString();
			}
			// this.loading = false;
		},

		formatDate(inputDate) {
			const options = { year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric'};
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
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div v-if="stream.length > 0">
			<h5>Photos posted by the users you follow </h5>
			<ul style="margin-top: 20px;" class="photo-list">
				<li v-for="photo in stream">
					<p class="photo-owner"> Posted by: {{ photo.username }} </p>
					<div class="photo-container">
						<img :src="'data:image/jpeg;base64,'+photo.image">
						<p class="photo-date"> {{ formatDate(photo.date) }} </p>
					</div>
				</li>
			</ul>
		</div>
		<div v-else>
			<h5>Your feed is empty! Follow someone!</h5>
		</div>
	</div>
</template>

<style>
	.photo-owner {
		position: absolute;
		top: 8px; /* Distanza dal bordo inferiore */
		left: 8px; /* Distanza dal bordo destro */
		background-color: rgba(255, 255, 255, 0.7); /* Colore di sfondo semitrasparente per migliorare la leggibilità */
		padding: 4px 8px; /* Spazio interno per migliorare l'aspetto */
		border-radius: 4px; /* Angoli arrotondati per il box della data */
		font-size: 14px; /* Dimensione del testo */
		color: #333; /* Colore del testo */
		font-family: Arial, Helvetica, sans-serif;
	}
</style>
