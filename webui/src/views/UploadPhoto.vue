<script>
import axios from 'axios';

export default {
    data: function() {
        return {
            image: null,
            imageUrl: null
        }
    },
    methods: {
        onChange(e) {
            const file = e.target.files[0];
            this.image = file;
            this.imageUrl = URL.createObjectURL(file);
        },
        async uploadImage() {

            try {
                let userID = localStorage.getItem('token');
                let response = await this.$axios.post(`/users/${userID}/photos/`, this.image);
                this.$router.push(`/users/${userID}/profile`);
            } catch (e) {
                this.errormsg = e.toString();
            }
        }
    }
}
</script>

<template>
    <div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">Upload Photo</h1>
	</div>

    <ErrorMsg v-if="errormsg"></ErrorMsg>
    <div>
      <input type="file" accept="image/*" @change="onChange" />
      <button @click="uploadImage">Upload Image</button>
    </div>
</template>
  