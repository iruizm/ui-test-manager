<script setup>
import FileTable from '@/components/FileTable.vue'
import SequenceGraph from '@/components/SequenceGraph.vue'

const components = {
  SequenceGraph
}
</script>

<template>
  <header>

  </header>

  <main>
    <div class="wrapper">
      <FileTable />
    </div>
    <div class="wrapper">
      <SequenceGraph />
    </div>
  </main>
</template>

<style scoped></style>
