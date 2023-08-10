import axios from 'axios';

export default {
  data() {
    return {
      responseData: null,
    };
  },
  methods: {
    async fetchData() {
      try {
        const response = await axios.get('http://your-go-api-url/api/endpoint');
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
