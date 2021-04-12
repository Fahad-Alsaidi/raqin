<template>

  <div class="q-ma-xl">
    <div class="q-py-md">
        <q-btn @click="newProject = !newProject">
        بدأ مشروع رقن
      </q-btn>
      </div>

      <ProjectsTable />

      <q-dialog v-model="newProject">
        <BookProject />
      </q-dialog>
  </div>
    
    
</template>

<script>

import BookProject from "components/bookProject/newBookProject";
import ProjectsTable from "components/bookProject/projectsTable"

export default {
  name: "project",
  components: {
    BookProject,
    ProjectsTable
  },
  data(){
    return {
      newProject: false
    }
  }
};
</script>
