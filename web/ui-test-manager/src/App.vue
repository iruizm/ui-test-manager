<template>
  <ClientAPI />
  <v-app>
    <v-app-bar color="teal-darken-3" elevation="2">
      <v-tabs v-model="tab" grow>
        <v-tab value="1"> Scripts </v-tab>
        <v-tab value="2"> Rules </v-tab>
      </v-tabs>
    </v-app-bar>

    <v-main>
      <v-container>
        <v-window v-model="tab">
          <v-window-item value="1">
            <v-row>
              <v-col cols="6">
                <FileTable />
              </v-col>
              <v-col cols="6">
                <SequenceGraph />
              </v-col>
            </v-row>
          </v-window-item>
          <v-window-item value="2">
            <PatternTable/>
          </v-window-item>
        </v-window>

        <v-row value="2">

        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import FileTable from '@/components/FileTable.vue'
import SequenceGraph from '@/components/SequenceGraph.vue'
import PatternTable from '@/components/PatternTable.vue'
import ClientAPI from '@/services/ClientAPI.vue'
import { provide, onMounted, ref } from 'vue';
import mitt from 'mitt';

const eventBus = mitt();
provide('eventBus', eventBus);

const tab = ref(0)

onMounted(() => {
  eventBus.emit("get-sequences")
  eventBus.emit("get-patterns")
});
</script>