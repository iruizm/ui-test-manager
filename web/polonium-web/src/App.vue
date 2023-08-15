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
  <header>
  </header>
  <main>
    <div class="wrapper">
      <FileTable />
    </div>
    <div class="wrapper">
      <SequenceGraph />
    </div>
  </main>
</template>

<style scoped></style>
