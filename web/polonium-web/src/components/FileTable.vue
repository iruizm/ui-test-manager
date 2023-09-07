<template>
    <div>
        <h2>Files</h2>
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Precedents</th>
                    <th>Content</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(sequence, index) in store.sequences" :key="index">
                    <td>{{ sequence.name }}</td>
                    <td>
                        <span v-if="hasPrecedents(sequence)">
                            <span class="sequence-item" v-for="item in sequence.precedents" :key="item">
                                <span class="sequence-item span">{{ store.sequences[item].name }}
                                    <button @click="removePrecedent(sequence.id, item)"
                                        class="delete-precedent;sequence-item button">&times;</button>
                                </span>
                            </span>
                        </span>
                    </td>
                    <td>
                        <button @click="openModal(index)">Open Modal</button>
                    </td>
                    <td>
                        <button @click="removeElement(sequence.id)">Remove</button>
                    </td>
                </tr>
                <ContentModal v-if="selectedSequence !== null" :sequence="store.sequences[selectedSequence].content"
                    @string-updated="updateString" @modal-closed="closeModal" />
            </tbody>
        </table>
        <button @click="openFileManager">Add Files...</button>
        <input type="file" ref="fileInput" style="display: none" @change="addElements" multiple />
    </div>
</template>

<script setup>
import { ref, inject } from 'vue'
import { store } from '../data/store.js'
import ContentModal from "@/components/ContentModal.vue"

const eventBus = inject('eventBus')
const fileInput = ref(null)

let selectedSequence = ref(null);

function openModal(index) {
    selectedSequence.value = index;
}

function closeModal() {
    selectedSequence.value = null;
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

function updateString({ item, newString }) {
    item.content = newString;
};

</script>

<style scoped>
table {
    border-collapse: collapse;
    width: 100%;
    margin-bottom: 20px;
}

thead th {
    background-color: #2d2d2d;
    color: #ffffff;
    border: 1px solid #333333;
    padding: 10px;
    text-align: left;
    font-weight: bold;
}

tbody tr:nth-child(even) {
    background-color: #252525;
}

td {
    border: 1px solid #333333;
    padding: 10px;
    color: #ffffff;
}

button {
    background-color: #2d2d2d;
    color: #ffffff;
    border: none;
    padding: 6px 12px;
    cursor: pointer;
    transition: background-color 0.3s ease-in-out;
}

button:hover {
    background-color: #4d4d4d;
}

.sequence-item {
    display: flex;
    align-items: center;
    margin-bottom: 2px;
}

.sequence-item span {
    display: flex;
    align-items: center;
    margin-bottom: 3px;
    background-color: #2d2d2d;
    color: #fff;
    border-radius: 1px;
    padding: 2px;
}

.sequence-item button {
    background-color: #616161;
    color: #fff;
    border: 1px;
    border-radius: 50%;
    margin-left: 3px;
    padding: 2px 3px;
    cursor: pointer;
    transition: background-color 0.3s ease-in-out;
}

.sequence-item button:hover {
    background-color: #c0392b;
}
</style>