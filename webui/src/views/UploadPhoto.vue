<script>

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
        },
        cancelImage() {
            this.image = null;
            this.imageUrl = null;
            this.$refs.fileInput.value = ''; 
        }
    }
}
</script>

<template>
    <div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">New Photo</h1>
	</div>

    <ErrorMsg v-if="errormsg"></ErrorMsg>

    <div class="upload-content">
        <input type="file" accept="image/*" @change="onChange" class="file-input" ref="fileInput"/>
        <div v-if="imageUrl" class="photo-container image-preview">
            <h5 class="d-flex justify-content-center">Image Preview</h5>
            <img :src="imageUrl" class="img-thumbnail" alt="Image preview" />
        </div>
        <div class="btn-toolbar">
            <div class="btn-group me-2">
                <button v-if="image" @click="uploadImage" class="btn btn-primary btn-lg">Upload Image</button>
            </div>
            <div class="btn-group me-2">
                <button v-if="image" @click="cancelImage" class="btn btn-secondary btn-lg">Cancel</button>
            </div>
        </div>
    </div>
</template>

<style>

    .upload-content {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .file-input {
        margin-bottom: 20px;
        padding: 10px;
        border: 1px solid #ccc;
        border-radius: 4px;
        font-size: 16px;
    }

    .image-preview {
        max-width: 300px; /* Larghezza massima della box della foto */
    }
    
</style>
  