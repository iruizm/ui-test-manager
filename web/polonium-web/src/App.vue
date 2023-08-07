<template>
  <div>
    <h1>Element Table</h1>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Value</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(element, index) in elements" :key="index">
          <td>{{ element.name }}</td>
          <td>{{ element.value }}</td>
          <td>
            <button @click="removeElement(index)">Remove</button>
          </td>
        </tr>
      </tbody>
    </table>

    <button @click="addElement">Add Element</button>

    <button @click="openFileManager">Add Files...</button>
    <input type="file" ref="fileInput" style="display: none" @change="addElements" />

  </div>
</template>

<script>
export default {
  data() {
    return {
      elements: [],
      fileManagerOpen: false
    };
  },
  methods: {
    addElements(event) {
      console.log(event.target.files)
      this.elements.concat(event.target.files)
    },
    removeElement(index) {
      this.elements.splice(index, 1);
    },
    openFileManager() {
      this.$refs.fileInput.click();
    },
  }
};
</script>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
