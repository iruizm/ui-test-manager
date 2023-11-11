<template>
    <ClientAPI />
  <v-app>
    <v-app-bar color="teal-darken-4" title="Test Manager" /> 

    <!-- <v-navigation-drawer>
      <v-list>
        <v-list-item title="Navigation drawer"></v-list-item>
      </v-list>
    </v-navigation-drawer> -->
    
    <v-main>
      <v-container>
        <v-row>
          <v-col cols="6">
            <FileTable />
          </v-col>
          <v-col cols="6">
            <SequenceGraph />
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
// import HelloWorld from './components/HelloWorld.vue';
import FileTable from '@/components/FileTable.vue'
import SequenceGraph from '@/components/SequenceGraph.vue'
import ClientAPI from '@/services/ClientAPI.vue'
import { provide, onMounted } from 'vue';
import mitt from 'mitt';

const eventBus = mitt();
provide('eventBus', eventBus);

onMounted(() => eventBus.emit("get-sequences"))
</script>
