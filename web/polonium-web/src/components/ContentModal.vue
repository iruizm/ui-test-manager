<template>
  <v-dialog v-model="dialog" max-width="800px">
    <template v-slot:activator="{ on }">
      <button @click="openDialog" class="edit-button">Edit String</button>
    </template>

    <!-- Rest of the modal content -->
  </v-dialog>
</template>
  
<script setup>
import { ref, defineProps, inject } from 'vue';
const eventBus = inject('eventBus')

const props = defineProps({
  item: Object, // Pass the item object from the parent component
});

const dialog = ref(false);
const editedString = ref('');

function openDialog() {
  dialog.value = true;
  editedString.value = props.item.string; // Modify to use the correct property from the item
};

function closeDialog() {
  dialog.value = false;
};

function saveAndClose() {
  dialog.value = false;
  // Emit the updated string and the corresponding item
  eventBus.emit('string-updated', { item: props.item, newString: editedString.value });
};
</script>
  
<style scoped>
/* Styles */
</style>
  