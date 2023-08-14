<script setup>
import FileTable from '@/components/FileTable.vue'
import SequenceGraph from '@/components/SequenceGraph.vue'
import Client from '@/components/Client.vue'
import { provide } from 'vue';
import mitt from 'mitt';

const components = {
  SequenceGraph, FileTable, Client
}

const eventBus = mitt();
provide('$eventBus', eventBus);

</script> 

<template>
  <header>
  </header>
  <main>
    <Client />
    <div class="wrapper">
      <FileTable />
    </div>
    <div class="wrapper">
      <SequenceGraph />
    </div>
  </main>
</template>

<style scoped></style>
