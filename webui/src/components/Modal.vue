<template>
    <div class="modal fade show" :class="{ show: isVisible }" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true" v-show="isVisible">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">Modal title</h5>
            <button type="button" class="btn-close" @click="closeModal" aria-label="Close">
            </button>
          </div>
          <div class="modal-body">
            aggiungi 2
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary btn-sm" @click="closeModal">Close</button>
            <button type="button" class="btn btn-primary btn-sm" @click="saveChanges">Save changes</button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
        isVisible: false
    },
    methods: {
      closeModal() {
        this.$emit('close');
      },
      saveChanges() {
        this.$emit('save');
      }
    }
  }
  </script>
  
  <style scoped>
  .modal.show {
    display: block;
    opacity: 1;
  }
  
  .modal-backdrop.show {
    opacity: 0.5;
  }
  </style>
  