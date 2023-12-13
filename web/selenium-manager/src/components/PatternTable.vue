<template>
  <v-row>
    <v-col cols="6">
      <v-card style="height: 100%" class="fullscreen-card">
        <v-toolbar color="teal-darken-3" density="comfortable" floating>
          <v-text-field v-model="searchText" append-icon="mdi-magnify" label="Search" single-line
            hide-details></v-text-field>
          <v-spacer></v-spacer>
          <v-tooltip text="Add" location="top">
            <template v-slot:activator="{ props }">
              <v-btn v-bind="props" variant="text" icon="mdi-plus-box" @click="addElement"></v-btn>
            </template>
          </v-tooltip>
        </v-toolbar>
        <v-card>
          <v-list elevation="4" class="ma-2" color="indigo-darken-3" min-height="659px" style="height: 100%">
            <v-list-item elevation="1" v-for="pattern in paginatedData" :key="pattern.id" :style="{ height: '65px' }"
              variant="plain">
              <template v-slot:prepend>
                {{ pattern.name }}
              </template>
              <template v-slot:append>
                <v-btn-group>
                  <v-tooltip text="Edit" location="top">
                    <template v-slot:activator="{ props }">
                      <v-btn v-bind="props" class="ma-1" icon="mdi-file-edit" @click="showPattern(pattern)"> </v-btn>
                    </template>
                  </v-tooltip>
                  <v-tooltip text="Remove" location="top">
                    <template v-slot:activator="{ props }">
                      <v-btn v-bind="props" @click="removeElement(pattern.id)" class="ma-1" icon="mdi-delete"></v-btn>
                    </template>
                  </v-tooltip>
                </v-btn-group>
              </template>
            </v-list-item>
          </v-list>
          <v-pagination small v-model="currentPage" :length="pageCount" @input="changePage"></v-pagination>
        </v-card>
      </v-card>
    </v-col>
    <v-col v-if="selected != null" cols="6">
      <v-card style="height: 100%" class="fullscreen-card">
        <v-toolbar color="teal-darken-3" density="comfortable" floating>
          <v-toolbar-title>Pattern</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-tooltip text="Save" location="top">
            <template v-slot:activator="{ props }">
              <v-btn v-bind="props" variant="text" icon="mdi-content-save" @click="saveElement()"></v-btn>
            </template>
          </v-tooltip>
        </v-toolbar>
        <v-row>
          <v-col cols="12">
            <v-text-field label="Name" v-model="selected.name"></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-card-text>
              <v-textarea label="Original" class="full-screen-textarea" rows="10" v-model="selected.regex"></v-textarea>
            </v-card-text>
          </v-col>
          <v-col cols="12">
            <v-card-text>
              <v-textarea label="Replacement" class="full-screen-textarea" rows="10"
                v-model="selected.replacement"></v-textarea>
            </v-card-text>
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</template>


<script setup>
import { ref, inject, computed, watch } from 'vue'
import { store } from '../data/store.js'

const eventBus = inject('eventBus')
const searchText = ref('');

const tableData = computed(() => {
  console.log(store.patterns)
  return Object.values(store.patterns).map(pattern => ({
    id: pattern.id,
    name: pattern.name,
    regex: pattern.regex,
    replacement: pattern.replacement,
  }));
});

const itemsPerPage = ref(10);
const currentPage = ref(1);

const filteredData = computed(() => {
  const lowerCaseSearch = searchText.value.toLowerCase();
  return tableData.value.filter(pattern =>
    pattern.name !== 'string' || pattern.name.toLowerCase().includes(lowerCaseSearch)
  );
});

const paginatedData = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage.value;
  const endIndex = startIndex + itemsPerPage.value;
  return filteredData.value.slice(startIndex, endIndex);
});

const pageCount = computed(() => {
  return Math.ceil(filteredData.value.length / itemsPerPage.value);
});

watch(searchText, () => {
  currentPage.value = 1;
});

const selected = ref(null)

function showPattern(pattern) {
  selected.value = pattern
  console.log(selected.value)
}

function changePage(page) {
  currentPage.value = page;
}

async function addElement(event) {
  const pattern = {
    name: "",
  };
  eventBus.emit("save-pattern", pattern);
  event.target.value = ''
}

function saveElement() {
  eventBus.emit("save-pattern", selected.value)
  selected.value = null
}

function removeElement(id) {
  eventBus.emit("delete-pattern", id)
}
</script>