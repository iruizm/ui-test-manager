<script>
import axios from 'axios';
import { inject } from 'vue';
import store from '../store/store'

export default {
  data() {
    return {
      responseData: null,
    };
  },
  mounted() {
    const eventBus = inject('$eventBus');
    this.fetchData()
    this.eventBus.on("get-data", async data => {
      try {
        const response = await axios.get();
        this.responseData = response.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    });
  },
  methods: {
    async fetchData() {
      try {
        const response = await axios.get('http://localhost:8080/api/sequences');
        console.log(response)
        this.responseData = response.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
  },
  mounted() {
    this.fetchData();
  },
};


</script>
