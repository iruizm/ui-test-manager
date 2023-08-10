<script>
import axios from 'axios';

export default {
  data() {
    return {
      responseData: null,
    };
  },
  mounted() { 
    this.emitter.on("get-data", async data => {
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
        const response = await axios.get('http://localhost/api/endpoint');
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
