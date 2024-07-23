
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
			<ul class="photo-list">
				<li v-for="photo in stream">
					Posted by: {{ photo.userID }}
					<div class="photo-container">
						<img :src="'data:image/jpeg;base64,'+photo.image">
						{{ formatDate(photo.date) }}
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

</style>
