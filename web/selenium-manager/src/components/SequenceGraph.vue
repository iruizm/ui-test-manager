<template>
    <v-card>
      <v-card-text>
        <div ref="container" style="border: 2px solid #333333"></div>
      </v-card-text>
    </v-card>
</template>

<script setup>
import { ref, inject, onMounted, toRaw, watch } from 'vue';
import { Network } from 'vis-network/standalone';
import { DataSet } from 'vis-data';
import { store } from '../data/store.js'

const eventBus = inject('eventBus');
const container = ref(null);
let network = null;
let selectedNode = null

const graphResize = () => {
  if (container.value) {
    container.value.style.height = `${window.innerHeight - 138}px`;
  }
};

onMounted(() => {

  graphResize();
  window.addEventListener('resize', graphResize);
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
        color: 'white',
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
    } else {
      selectedNode = null
    }
  });

});
</script>