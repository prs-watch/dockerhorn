<template lang="md">
  <el-container>
    <el-header>
      <div v-html="$md.render('`docker info`')" class="markdown-body subtitle"></div>
    </el-header>
    <el-main>
      <el-row :gutter="12">
        <el-col :span="6">
          <el-card class="box-card">
            <div slot="header" class="card-title">
              <span>Docker Engine Version</span>
            </div>
            <div>
              {{ info.ServerVersion }}
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="box-card">
            <div slot="header" class="card-title">
              <span>Total Memory</span>
            </div>
            <div>
              {{ info.MemTotal }}
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="box-card">
            <div slot="header" class="card-title">
              <span>Images</span>
            </div>
            <div>
              {{ info.Images }}
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="box-card">
            <div slot="header" class="card-title">
              <span>Containers</span>
            </div>
            <div>
              Running {{ info.ContainersRunning }} / Paused {{ info.ContainersPaused }} / Stopped {{ info.ContainersStopped }}
            </div>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      info: {}
    };
  },
  mounted() {
    this.$axios.get("/info").then(res => {
      this.info = res.data.Info;
    });
  }
};
</script>

<style>
.el-row {
  margin-bottom: 1rem;
}
</style>
