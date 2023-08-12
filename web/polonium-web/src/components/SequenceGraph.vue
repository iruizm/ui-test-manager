<template>
  <div>
    <div ref="networkContainer" style="width: 800px; height: 600px;"></div>
  </div>
</template>

<script>
import { ref, inject, onMounted } from 'vue';
import { Network } from 'vis-network';
import { DataSet } from 'vis-data';
export default {
  name: 'SequenceGraph',
  setup() {
    const networkContainer = ref(null);
    let network = null;

    onMounted(() => {
      // Define nodes and edges
      const nodes = new DataSet([
        { id: 1, label: 'Node 1' },
        { id: 2, label: 'Node 2' },
        { id: 3, label: 'Node 3' },
      ]);

      const edges = new DataSet([
        { from: 1, to: 2 },
        { from: 2, to: 3 },
        { from: 3, to: 1 },
      ]);

      // Set up options for the network
      const options = {};

      // Create the network
      const data = {
        nodes: nodes,
        edges: edges,
      };

      network = new Network(networkContainer.value, data, options);
    });

    return {
      networkContainer,
    };
  },
};
</script>

<style scoped>
.network-container {
  width: 100%;
  height: 500px;
}
</style>