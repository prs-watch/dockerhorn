<template>
  <div>
    <el-container>
      <el-header>
        <div v-html="$md.render('`docker pull`')" class="markdown-body subtitle"></div>
      </el-header>
      <el-main>
        <el-row :gutter="12">
          <el-col :span="18">
            <el-input v-model="keyword" placeholder="Keyword.." />
          </el-col>
          <el-col :span="6">
            <el-button type="primary" @click="execSearch(keyword)">Search Docker image</el-button>
          </el-col>
        </el-row>
        <el-table :data="searchResult" stripe max-height="300" v-loading="pulling">
          <el-table-column label="Official">
            <template slot-scope="scope">
              <i class="el-icon-check" v-if="scope.row.official === true" />
            </template>
          </el-table-column>
          <el-table-column prop="name" label="Image Name" />
          <el-table-column prop="stars" label="Starts" />
          <el-table-column prop="description" label="Description" />
          <el-table-column>
            <template slot-scope="scope">
              <el-button type="primary" @click="execPull(scope.row.name)">Pull</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>
    <el-container>
      <el-header>
        <div v-html="$md.render('`docker images`')" class="markdown-body subtitle"></div>
      </el-header>
      <el-main>
        <el-table :data="images" stripe max-height="400">
          <el-table-column prop="repo" label="Repository" />
          <el-table-column prop="version" label="Version" />
          <el-table-column prop="size" label="Size" />
          <el-table-column>
            <template slot-scope="scope">
              <el-button type="danger" @click="execRemove(scope.row.id)">Remove</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>
  </div>
</template>

<script>
// compare function
const compareStars = (a, b) => {
  var r = 0;
  if (a.star_count < b.star_count) {
    r = 1;
  } else if (a.star_count > b.star_count) {
    r = -1;
  }
  return r;
};

export default {
  data() {
    return {
      keyword: "",
      images: [],
      searchResult: [],
      pulling: false
    };
  },
  mounted() {
    this.$axios.get("/image/images").then(res => {
      const images = res.data.Images;
      let rows = [];
      images.forEach(image => {
        const repoTags = image.RepoTags;
        repoTags.forEach(repoTag => {
          const row = {
            id: image.Id,
            repo: repoTag.split(":")[0],
            version: repoTag.split(":")[1],
            size: image.Size
          };
          rows.push(row);
        });
      });
      this.images = rows;
    });
  },
  methods: {
    execRemove(id) {
      this.$axios.get("/image/remove/" + id).then(res => {
        if (res.data.Error != null) {
          this.$message.error(res.data.Error);
          return
        }
        window.location.reload(true);
      });
    },
    execSearch(keyword) {
      this.$axios.get("/image/search/" + keyword).then(res => {
        if (res.data.Error != null) {
          this.$message.error(err.data.Error);
          return
        }
        const images = res.data.Images;
        images.sort(compareStars);
        let rows = [];
        images.forEach(image => {
          const row = {
            stars: image.star_count,
            official: image.is_official,
            name: image.name,
            description: image.description
          };
          rows.push(row);
        });
        this.searchResult = rows;
      });
    },
    execPull(name) {
      this.pulling = true;
      this.$axios.get("/image/pull/" + name).then(res => {
        if (res.data.Error != null) {
          this.$message.error(err.data.Error);
          return
        }
        window.location.reload(true);
      });
    }
  }
};
</script>
