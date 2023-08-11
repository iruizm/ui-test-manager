<template>
  <div>
    <div ref="networkContainer" style="width: 800px; height: 600px;"></div>
  </div>
</template>

<script>
import { ref, inject } from 'vue';
import { Network } from 'vis-network';
import { DataSet } from 'vis-data';
export default {
  name: 'SequenceGraph',  
  mounted() {
    this.networkContainer = ref(null);
    if (this.networkContainer.value) {
      // Create a dataset for nodes and edges
      const nodes = new DataSet([
        { id: 1, label: 'Node 1' },
        { id: 2, label: 'Node 2' },
        // Add more nodes...
      ]);

      const edges = new DataSet([
        { from: 1, to: 2 },
        // Add more edges...  
      ]);

      const data = { nodes, edges };

      const options = {
        // Configure options...
      };

      new Network(networkContainer.value, data, options);
    }

  }
};
</script>

<style scoped>
.network-container {
  width: 100%;
  height: 500px;
}
</style>