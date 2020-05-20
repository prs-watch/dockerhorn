<template>
  <el-container>
    <el-header>
      <div v-html="$md.render('`docker ps -a`')" class="markdown-body subtitle"></div>
    </el-header>
    <el-main>
      <el-table :data="containers" stripe>
        <el-table-column label="Status">
          <template slot-scope="scope">
            <el-tag
              :type="scope.row.state === 'running' ? 'success' : 'danger'"
            >{{ scope.row.state }}</el-tag>
            <span>{{ scope.row.status }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="Name" />
        <el-table-column prop="image" label="Image" />
        <el-table-column>
          <template slot-scope="scope">
            <el-button
              type="warning"
              @click="execStop(scope.row.id)"
              v-if="scope.row.state === 'running'"
            >Stop</el-button>
            <el-button type="primary" @click="execStart(scope.row.id)" v-else>Start</el-button>
          </template>
        </el-table-column>
        <el-table-column>
          <template slot-scope="scope">
            <el-button type="success" @click="preCommit(scope.row.id)">Commit</el-button>
          </template>
        </el-table-column>
        <el-table-column>
          <template slot-scope="scope">
            <el-button
              type="danger"
              @click="execRemove(scope.row.id)"
              v-if="scope.row.state !== 'running'"
            >Remove</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-dialog :visible.sync="dialogVisible" width="50%" :before-close="handleClose">
        <el-form :rules="rules">
          <el-form-item label="Container ID">
            <el-input :disabled="true" v-model="containerID" />
          </el-form-item>
          <el-form-item label="Repository" prop="repo">
            <el-input v-model="repo" />
          </el-form-item>
          <el-form-item label="Tag" prop="tag">
            <el-input v-model="tag" />
          </el-form-item>
          <el-form-item>
            <el-button @click="dialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="execCommit()">Commit</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      containers: [],
      dialogVisible: false,
      containerID: "",
      repo: "",
      tag: "",
      rules: {
        repo: [{ required: true, message: "Repository must be filled in!" }],
        tag: [{ required: true, message: "Tag must be filled in!" }]
      }
    };
  },
  mounted() {
    this.$axios.get("/container/ps").then(res => {
      const data = res.data.Containers;
      let rows = [];
      data.forEach(container => {
        const row = {
          id: container.Id,
          name: container.Names[0],
          image: container.Image,
          state: container.State,
          status: container.Status
        };
        rows.push(row);
      });
      this.containers = rows;
    });
  },
  methods: {
    execStop(id) {
      this.$axios.get("/container/stop/" + id).then(res => {
        if (res.data.Error != null) {
          this.$message.error(err.data.Error);
          return;
        }
        window.location.reload(true);
      });
    },
    execStart(id) {
      this.$axios.get("/container/start/" + id).then(res => {
        if (res.data.Error != null) {
          this.$message.error(err.data.Error);
          return;
        }
        window.location.reload(true);
      });
    },
    execRemove(id) {
      this.$axios.get("/container/remove/" + id).then(res => {
        if (res.data.Error != null) {
          this.$message.error(err.data.Error);
          return;
        }
        window.location.reload(true);
      });
    },
    preCommit(id) {
      this.dialogVisible = true;
      this.containerID = id;
    },
    execCommit() {
      this.$axios
        .get(
          "/container/commit/" +
            this.containerID +
            "?repo=" +
            this.repo +
            "&tag=" +
            this.tag
        )
        .then(res => {
          if (res.data.Error != null) {
            this.$message.error(err.data.Error);
            return;
          }
          this.$router.push("/image");
        });
    }
  }
};
</script>
