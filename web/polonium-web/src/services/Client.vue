<script setup>
import axios from 'axios';
import { inject, ref, onMounted } from 'vue';
import { store } from '../data/store.js'

const eventBus = inject('eventBus')
eventBus.on("get-sequences", getSequences);
eventBus.on("save-sequence", saveSequence)
eventBus.on("delete-sequence", deleteSequence)
eventBus.on("delete-precedent", deletePrecedent)

async function getSequences() {
  try {
    const res = await axios.get('http://localhost:8080/api/sequences');
    store.sequences = res.data
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
async function saveSequence(data) {
  try {
    const res = await axios.post('http://localhost:8080/api/sequences', data);
    await getSequences()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
async function deleteSequence(id) {
  try {
    const res = await axios.delete(`http://localhost:8080/api/sequences/${id}`);
    await getSequences()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}

async function deletePrecedent(e) {
  try {
    const res = await axios.delete(`http://localhost:8080/api/sequences/${e.idSequence}/${e.idPrecedent}`);
    await getSequences()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
</script>
