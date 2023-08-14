<script setup>
import axios from 'axios';
import { inject, ref, onMounted } from 'vue';
import { store } from '../data/store.js'

const apiClient = ref(null);
const eventBus = inject('$eventBus')

onMounted(() => {
  getSequences()
});

eventBus.on("get-sequences", getSequences);
eventBus.on("save-sequence", saveSequence)
eventBus.on("delete-sequence", deleteSequence)

async function getSequences() {
  try {
    const res = await axios.get('http://localhost:8080/api/sequences');
    store.sequences = res.data
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
async function saveSequence(json) {
  try {
    const res = await axios.post('http://localhost:8080/api/sequences/add', json);
    await getSequences()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
async function deleteSequence(id) {
  try {
    const res = await axios.delete('http://localhost:8080/api/sequences/remove', json);
    await getSequences()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
</script>
