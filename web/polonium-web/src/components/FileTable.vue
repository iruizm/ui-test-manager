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
                <tr v-for="(element, index) in store" :key="index">
                    <td>{{ element.name }}</td>
                    <td>{{ element.premises }}</td>
                    <td>
                        <button @click="removeElement(index)">Remove</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <button @click="openFileManager">Add Files...</button>
        <input type="file" ref="fileInput" style="display: none" @change="addElements" multiple />
    </div>
</template>

<script>
import { inject, onMounted } from 'vue';
import { store } from '../data/store.js'

export default {
    name: 'FileTable',
    setup() {
        const eventBus = inject('$eventBus');
        onMounted(() => {
            eventBus.on("dataLoaded", isOpen => {
                this.isOpen = isOpen;
            })
        })
        return {
            store
        }
    },
    methods: {
        async loadAll() {
            //this.eventBus.emit("toggle-sidebar", this.sidebarOpen);
            this.eventBus.emit("get-sequences")
        },
        async addElements(event) {
            for (var i = 0; i < event.target.files.length; i++) {
                const file = event.target.files[i]
                const fileContent = await this.readFile(file);
                store.push({
                    fileName: file.name,
                    name: file.name.split(".")[0],
                    id: null,
                    value: fileContent,
                    premises: []
                });
            }
        },
        removeElement(index) {
            this.store.splice(index, 1);
        },
        openFileManager() {
            this.$refs.fileInput.click();
        },
        async readFile(file) {
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
    }
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

zº button {
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