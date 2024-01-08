<template>
  <v-card style="height: 100%">
    <v-toolbar color="teal-darken-3" density="comfortable" floating>


      <v-text-field v-model="searchText" append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
      <v-spacer></v-spacer>
      <v-tooltip text="Add" location="top">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" variant="text" icon="mdi-plus-box" @click="openFileManager"></v-btn>
        </template>
      </v-tooltip>
      <v-tooltip text="Generate" location="top">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" variant="text" icon="mdi-play-circle-outline" @click="openGenerateDialog"></v-btn>
        </template>
      </v-tooltip>
    </v-toolbar>
    <v-list elevation="2" class="ma-2" color="indigo-darken-3">
      <v-list-item elevation="1" v-for="sequence in paginatedData" :key="sequence.id" :style="{ height: '64px' }"
        variant="plain">
        <template v-slot:prepend>
          {{ sequence.name }}
        </template>
        <v-container v-if="hasPrecedents(sequence)">
          <v-slide-group show-arrows>
            <v-slide-group-item v-for="item in sequence.precedents" :key="item">
              <v-chip class="ma-1" :closable="true" :ripple="true" variant="tonal" size="small"
                @click:close="removePrecedent(sequence.id, item)"> {{ store.sequences[item].name }}
              </v-chip>
            </v-slide-group-item>
          </v-slide-group>
        </v-container>
        <template v-slot:append>
          <v-btn-group>
            <v-tooltip text="View" location="top">
              <template v-slot:activator="{ props }">
                <v-btn v-bind="props" class="ma-1" icon="mdi-eye" @click="openContentDialog(sequence)"></v-btn>
              </template>
            </v-tooltip>
            <v-tooltip text="Remove" location="top">
              <template v-slot:activator="{ props }">
                <v-btn v-bind="props" @click="removeElement(sequence.id)" class="ma-1" icon="mdi-delete"></v-btn>
              </template>
            </v-tooltip>
          </v-btn-group>
        </template>
      </v-list-item>
    </v-list>
    <v-pagination small v-model="currentPage" :length="pageCount" @input="changePage"></v-pagination>
  </v-card>

  <template>
    <v-row justify="center">
      <v-dialog v-model="contentVisible" width="1024">
        <v-card class="fullscreen-card">
          <v-card-title>Script</v-card-title>
          <v-card-text>
            <v-textarea class="full-screen-textarea" v-model="sequenceRef.content" auto-grow></v-textarea>
          </v-card-text>
          <v-card-actions class="justify-end">
            <v-tooltip text="Save">
              <template v-slot:activator="{ props }">
                <v-btn v-bind="props" class="ma-1" icon="mdi-content-save" @click="saveContentDialog"></v-btn>
              </template>
            </v-tooltip>
            <v-tooltip text="Close">
              <template v-slot:activator="{ props }">
                <v-btn v-bind="props" class="ma-1" icon="mdi-close" @click="closeContentDialog"></v-btn>
              </template>
            </v-tooltip>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </template>

  <template>
    <v-dialog v-model="generateVisible" width="600px">
      <v-card>
        <v-toolbar dark color="teal-darken-3">
          <v-btn icon @click="closeGenerateDialog">
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-toolbar-title>Sequence</v-toolbar-title>
        </v-toolbar>
        <v-card-text>
          <v-list elevation="4" height="670px" class="ma-2" color="indigo-darken-3">
            <v-list-item elevation="1" v-for="test in store.tests" :key="test.id" :style="{ height: '65px' }"
              variant="plain">
              <template v-slot:prepend>
                {{ test.name }}
              </template>
              <template v-slot:append>
                <v-btn-group>
                  <v-btn class="ma-1" icon="mdi-eye" @click="openContentDialog(test)">
                  </v-btn>
                </v-btn-group>
              </template>
            </v-list-item>
          </v-list>
        </v-card-text>
        <v-card-actions class="justify-end">
          <v-tooltip text="Save">
            <template v-slot:activator="{ props }">
              <v-btn v-bind="props" class="ma-1" icon="mdi-content-save" @click="generateFiles"></v-btn>
            </template>
          </v-tooltip>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </template>
  <input type="file" ref="fileInput" style="display: none" @change="addElements" multiple />
</template>


<script setup>
import { ref, inject, computed, watch } from 'vue'
import { store } from '../data/store.js'
import JSZip from 'jszip';

const eventBus = inject('eventBus')
const fileInput = ref(null)
const searchText = ref('');

const tableData = computed(() => {
  return Object.values(store.sequences).map(sequence => ({
    id: sequence.id,
    name: sequence.name,
    precedents: sequence.precedents,
    content: sequence.content,
  }));
});

const itemsPerPage = ref(10);
const currentPage = ref(1);

const filteredData = computed(() => {
  const lowerCaseSearch = searchText.value.toLowerCase();
  return tableData.value.filter(sequence =>
    sequence.name?.toLowerCase().includes(lowerCaseSearch)
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


function changePage(page) {
  currentPage.value = page;
}

async function addElements(event) {
  for (const file of event.target.files) {
    const fileContent = await readFile(file);
    const sequence = {
      name: file.name.split(".")[0],
      content: fileContent
    };
    eventBus.emit("save-sequence", sequence);
  }
  event.target.value = ''
}
function removeElement(id) {
  eventBus.emit("delete-sequence", id)
}
function removePrecedent(idSequence, idPrecedent) {
  var e = {
    idSequence: idSequence,
    idPrecedent: idPrecedent
  }
  eventBus.emit("delete-precedent", e)
}

function openFileManager() {
  console.log(tableData)
  fileInput.value.click();
}

async function readFile(file) {
  return await new Promise((resolve, reject) => {
    const reader = new FileReader();

    reader.onload = (event) => {
      const content = event.target.result;
      resolve(content);
    };

    reader.onerror = (event) => {
      reject(event.target.error);
    };

    reader.readAsText(file);
  });
}

function hasPrecedents(element) {
  return (
    element?.precedents?.length !== undefined &&
    element.precedents.length > 0
  );
}

const sequenceRef = ref(null);
const generateRef = ref(null)
const contentVisible = ref(null);
const generateVisible = ref(null);


function openContentDialog(content) {
  sequenceRef.value = content;
  contentVisible.value = true;
}

function saveContentDialog() {
  eventBus.emit("save-sequence", sequenceRef.value)
  contentVisible.value = false;
}

function closeContentDialog() {
  contentVisible.value = false;
}

function openGenerateDialog(content) {
  generateTests();
  generateRef.value = content;
  generateVisible.value = true;
}

function closeGenerateDialog() {
  generateVisible.value = false;
}

function generateTests(){
  eventBus.emit("get-tests")
}

function generateFiles() {
  const zip = new JSZip();
  var padding = store.tests.length.toString().length + 1
  store.tests.forEach((element, index) => {
    const fileName = `test_${(index + 1).toString().padStart(padding, '0')}_${element.name}.py`;
    const fileContent = element.content;
    zip.file(fileName, fileContent);
  });

  zip.generateAsync({ type: 'blob' }).then(content => {
    const link = document.createElement('a');
    link.href = URL.createObjectURL(content);
    var date = (new Date()).toISOString().replace(/[:._]/g, '-');
    link.download = 'tests_' + date + '.zip';
    link.click();
  });
}
</script>