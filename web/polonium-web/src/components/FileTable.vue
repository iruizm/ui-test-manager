<template>
    <div>
        <h2>Files</h2>
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Premises</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(element, index) in store.sequences" :key="index">
                    <td>{{ element.name }}</td>
                    <td>{{ element.premises }}</td>
                    <td>
                        <button @click="removeElement(element.id)">Remove</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <button @click="openFileManager">Add Files...</button>
        <input type="file" ref="fileInput" style="display: none" @change="addElements" multiple />
    </div>
</template>

<script setup>
import { ref, inject } from 'vue';
import { store } from '../data/store.js'

const eventBus = inject('$eventBus');
const fileInput = ref(null);

async function addElements(event) {
    for (const file of event.target.files) {
        const fileContent = await readFile(file);
        const sequence = {
            name: file.name.split(".")[0],
            content: fileContent
        };
        eventBus.emit("save-sequence", sequence);
    }
}
function removeElement(id) {
    eventBus.emit("delete-sequence", id)
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
    background-color: #3a3a3a;
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
</style>