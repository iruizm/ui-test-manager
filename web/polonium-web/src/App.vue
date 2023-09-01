<script setup>
import FileTable from '@/components/FileTable.vue'
import SequenceGraph from '@/components/SequenceGraph.vue'
import Client from '@/services/Client.vue'
import { provide, onMounted } from 'vue';
import mitt from 'mitt';

const eventBus = mitt();
provide('eventBus', eventBus);

onMounted(() => eventBus.emit("get-sequences"))

const components = {
  SequenceGraph, FileTable, Client
}

</script> 

<template>
  <Client />

  <div class="dashboard">
    <div class="table">
      <FileTable />
    </div>
    <div class="graph">
      <SequenceGraph />
    </div>
  </div>
</template>

<style>
.dashboard {
  display: flex;
  height: 100vh;
  width: 80vw;
}

.table,
.graph {
  flex: 1;
  padding: 1rem;
  box-sizing: border-box;
  overflow: auto;
  /* Add scrolling if content overflows */
}
</style>
