<template>
  <div>
    <div ref="container" style="width: 800px; height: 600px;"></div>
  </div>
</template>

<script setup>
import { ref, inject, onMounted, toRaw, watch } from 'vue';
import { Network } from 'vis-network';
import { DataSet } from 'vis-data';
import { store } from '../data/store.js'

const eventBus = inject('eventBus');
const container = ref(null);
let network = null;
let selectedNode = null

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
  const options = {
    nodes: {
      shape: 'square',
      size: 25,
      color: 'white',
      font: {
        size: 14,
        color: '#white',
      },
    }
  };

  // Create the network
  const data = {
    nodes: nodes,
    edges: edges,
  };
  network = new Network(container.value, data, options);

  watch(() => store.sequences, () => {
    data.nodes = new DataSet(Object.keys(store.sequences).map((key) => ({ id: store.sequences[key].id, label: store.sequences[key].name })));
    data.edges = new DataSet(Object.keys(store.sequences).flatMap((key) =>
      store.sequences[key].precedents.map((precedent) => ({ from: precedent, to: store.sequences[key].id }))
    ))
    network.setData(data) 
  });

  network.on('click', (params) => {
    if (params.nodes.length === 1) {
      let clickedNode = params.nodes[0];
      if (selectedNode) {
        var sequence = store.sequences[selectedNode]
        if (!sequence.precedents.includes(selectedNode)) {
          sequence.precedents.push(clickedNode)
          eventBus.emit("save-sequence", toRaw(sequence));
        }
        selectedNode = null
        network.unselectAll()
      } else {
        selectedNode = clickedNode;
      }
    }
  });
});
</script>

<style scoped>
.network-container {
  width: 100%;
  height: 500px;
}
</style>