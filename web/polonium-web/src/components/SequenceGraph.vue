<template>
  <div>
    <div ref="container" style="height: 600px;border: 2px solid #333333"></div>
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
  const nodes = new DataSet([]);

  const edges = new DataSet([]);

  const options = {
    nodes: {
      shape: 'dot',
      size: 25,
      color: {
        background: '#2d2d2d',
        border: 'grey',
        highlight: {
          background: '#2d2d2d',
          border: 'white',
        }
      },
      font: {
        size: 14,
        color: '#white',
      },
    },
    edges: {
      color: 'white',
      arrows: {
        to: {
          enabled: true,
          scaleFactor: 0.5
        }
      }
    }
  };

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
        var sequence = store.sequences[clickedNode]
        if (!sequence.precedents.includes(selectedNode)) {
          sequence.precedents.push(selectedNode)
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