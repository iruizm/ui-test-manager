<template>
    <div>
        <h2>Files</h2>
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Content</th>
                    <th>Premises</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(element, index) in store.sequences" :key="index">
                    <td>{{ element.name }}</td>
                    <td>
                        <div v-popover="`popover-${element.content}`" @mouseover="showPopover(element.id)"
                            @mouseleave="hidePopover(element.id)"><button></button>
                        </div>
                    </td>
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
    <div v-for="item in store.sequences" :key="item.id" :id="`popover-${item.id}`" class="popover"
        v-show="activePopover === item.id">
        {{ item.popoverContent }}
    </div>
</template>

<script setup>
import { ref, inject } from 'vue'
import { store } from '../data/store.js'
import { createPopper } from '@popperjs/core';

const eventBus = inject('eventBus')
const fileInput = ref(null)

async function addElements(event) {
    for (const file of event.target.files) {
        const fileContent = await readFile(file);
        const sequence = {
            name: file.name.split(".")[0],
            content: fileContent
        };
        console.log(sequence)
        eventBus.emit("save-sequence", sequence);
    }
    eventBus.emit("get-sequences")
}
function removeElement(id) {
    eventBus.emit("delete-sequence", id)
    eventBus.emit("get-sequences")
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
let activePopover = null;

const showPopover = (id) => {
    activePopover = id;
};

const hidePopover = (id) => {
    activePopover = null;
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