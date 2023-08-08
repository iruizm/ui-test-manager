<template>
  <div>
    <div id="graph"></div>
    <button @click="addNode">Add Node</button>
  </div>
</template>

<script>
import { Network } from "vis-network/standalone/esm/vis-network.esm";
import "vis-network/styles/vis-network.css";

export default {
  data() {
    return {
      nodes: new vis.DataSet([]),
      edges: new vis.DataSet([]),
      network: null
    };
  },
  methods: {
    addNode() {
      const newNodeId = this.nodes.length + 1;
      this.nodes.add({ id: newNodeId, label: `Node ${newNodeId}` });
    },
    addEdge(event) {
      const fromNode = this.network.getSelectedNodes()[0];
      const toNode = event.nodes[0];
      if (fromNode !== toNode) {
        this.edges.add({ from: fromNode, to: toNode });
      }
    }
  },
  mounted() {
    const container = document.getElementById("graph");
    const data = {
      nodes: this.nodes,
      edges: this.edges
    };
    const options = {};
    this.network = new Network(container, data, options);
    this.network.on("click", this.addEdge);
  }
};
</script>

<style>
/* You can add any custom styles here */
#graph {
  width: 100%;
  height: 400px;
}
</style>
