<template>
  <el-container>
    <el-header height="100%">
      <template>
        <i class="el-icon-sunrise-1 title"></i>
        <span class="title" @click="$router.push('/')">dockerhorn</span>
      </template>
    </el-header>
    <el-container>
      <el-aside width="20%">
        <el-menu>
          <el-menu-item index="dashboard" @click="$router.push('/')">
            <template slot="title">
              <i class="el-icon-odometer"></i>
              <span>Dashboard</span>
            </template>
          </el-menu-item>
          <el-menu-item index="image" @click="$router.push('/image')">
            <template slot="title">
              <i class="el-icon-picture-outline"></i>
              <span>Image</span>
            </template>
          </el-menu-item>
          <el-menu-item index="container" @click="$router.push('/container')">
            <template slot="title">
              <i class="el-icon-takeaway-box"></i>
              <span>Container</span>
            </template>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main>
        <nuxt />
      </el-main>
    </el-container>
  </el-container>
</template>

<style>
@import url(https://fonts.googleapis.com/earlyaccess/roundedmplus1c.css);

body {
  font-family: "Rounded Mplus 1c";
  font-size: 1rem;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

/** header */
.title {
  font-size: 3rem;
  font-weight: bold;
}

.subtitle {
  font-size: 1.5rem;
}

/** sidemenu */
.el-submenu__title,
.el-menu-item {
  font-size: 1rem;
}

/** card */
.card-title {
  font-weight: bold;
}

.card-content {
  text-align: left;
}
</style>
