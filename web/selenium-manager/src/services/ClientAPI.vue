<script setup>
import axios from 'axios';
import { inject} from 'vue';
import { store } from '../data/store.js'

const eventBus = inject('eventBus')
eventBus.on("get-sequences", getSequences)
eventBus.on("get-patterns", getPatterns)
eventBus.on("get-tests", getTests)
eventBus.on("save-sequence", saveSequence)
eventBus.on("save-pattern", savePattern)
eventBus.on("delete-sequence", deleteSequence)
eventBus.on("delete-precedent", deletePrecedent)
eventBus.on("delete-pattern", deletePattern)

async function getSequences() {
  try {
    const res = await axios.get('http://localhost:8080/api/sequences');
    store.sequences = res.data
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
async function getPatterns() {
  try {
    const res = await axios.get('http://localhost:8080/api/patterns');
    store.patterns = res.data
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
async function getTests() {
  try {
    const res = await axios.get('http://localhost:8080/api/tests');
    store.tests = res.data
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
async function saveSequence(data) {
  try {
    await axios.post('http://localhost:8080/api/sequences', data);
    await getSequences()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
async function savePattern(data) {
  try {
    await axios.post('http://localhost:8080/api/patterns', data);
    await getPatterns()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
async function deleteSequence(id) {
  try {
    await axios.delete(`http://localhost:8080/api/sequences/${id}`);
    await getSequences()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}

async function deletePattern(id) {
  try {
    await axios.delete(`http://localhost:8080/api/patterns/${id}`);
    await getPatterns()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}

async function deletePrecedent(e) {
  try {
    await axios.delete(`http://localhost:8080/api/sequences/${e.idSequence}/${e.idPrecedent}`);
    await getSequences()
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
</script>
