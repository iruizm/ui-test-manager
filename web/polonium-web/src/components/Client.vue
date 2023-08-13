<script>
import axios from 'axios';
import { inject, ref, onMounted } from 'vue';
import { store } from '../data/store.js'

export default {
  name: 'Client',
  setup() {
    const apiClient = ref(null);
    const eventBus = inject('$eventBus')
    onMounted(() => {
      getSequences()
    });

    eventBus.on("get-sequences", async data => {
      console.log(this.getSequences())
    });

    const getSequences = async () => {
      try {
        console.log("AHAHAHA")
        const res = await axios.get('http://localhost:8080/api/sequences');
        console.log(res)
        return res.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }
    const saveSequence = async (json) => {
      try {
        const res = await axios.get('http://localhost:8080/api/sequences/add', json);
        console.log(response)
        this.responseData = response.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }
    const removeSequence = async (json) => {
      try {
        const res = await axios.post('http://localhost:8080/api/sequences/remove', json);
        console.log(response)
        this.responseData = response.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }


    return {
      getSequences, saveSequence, removeSequence
    };
  }
}
</script>
