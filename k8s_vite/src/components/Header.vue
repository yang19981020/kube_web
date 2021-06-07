<template>
  <div class="header">
    <div class="title" @click="router.push('/')">
      k8s_vite
    </div>
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
      <el-menu-item index="1">集群状态</el-menu-item>
      <el-menu-item index="2">
        <el-submenu index="2-1">
        <template #title>资源管理</template>
        <el-menu-item index="2-2">pod</el-menu-item>
        <el-menu-item index="2-3">有状态应用</el-menu-item>
        <el-menu-item index="2-4">无状态应用</el-menu-item>
        </el-submenu>
      </el-menu-item>
      <el-menu-item index="3">
        <el-submenu index="3-1">
          <template #title>授权管理</template>
          <el-menu-item index="3-2">role_list_binding</el-menu-item>
          <el-menu-item index="3-3">clusterRole_list</el-menu-item>
          <el-menu-item index="3-4">clusterRole_list_binding</el-menu-item>
        </el-submenu>
      </el-menu-item>
      <el-menu-item index="4">
        <el-submenu index="4-1">
          <template #title>存储管理</template>
          <el-menu-item index="4-2">pvc_list_binding</el-menu-item>
        </el-submenu>
      </el-menu-item>
      <el-menu-item index="5">镜像管理</el-menu-item>
    </el-menu>
    <div class="go-github" @click="goGitHub">
      <i class="icon el-icon-s-comment"></i> GitHub
    </div>
  </div>
</template>

<script lang="ts">
import { useRouter } from 'vue-router'
import { defineComponent, ref, reactive, toRefs, onMounted, watch } from 'vue'
export interface NavItem {
  path: string
  name: string
  isActive: boolean
}
export default defineComponent({
  name: 'Header',
  methods: {
      goGitHub(){
        window.open('https://github.com/haozheyu')
      },
  },
  setup(){
    const activeIndex = ref('1')
    const router = useRouter()
    const handleSelect = (key:string) => {
        // console.log(key, keyPath)
       switch(key) {
         case "1" : {
          router.push("/")
          break;
         }
         case "2" : {
          router.push("/Resource")
          break;
         }
         case "2-2" : {
          router.push("/Pod")
          break;
         }
         case "2-3" : {
          router.push("/deployments")
          break;
         }
         case "2-4" : {
          router.push("/statefulSets")
          break;
         }
         case "3" : {
          router.push("/Templeton")
          break;
         }
         case "4" : {
          router.push("/Stores")
          break;
         }
         case "5" : {
          router.push("/Images")
          break;
         }
       }
    }
    watch(
      () => router.currentRoute.value,
      (_n) => {
         switch(_n.path) {
         case "/" : {
          activeIndex.value = "1"
          break;
         }
         case "/Resource" : {
          activeIndex.value = "2"
          break;
         }
         case "/Resource/pod" : {
          activeIndex.value = "2"
          break;
         }
         case "/Resource/deployments" : {
          activeIndex.value = "2"
          break;
         }
         case "/Resource/statefulSets" : {
          activeIndex.value = "2"
          break;
         }
         case "/Templeton" : {
          activeIndex.value = "3"
          break;
         }
         case "/Stores" : {
          activeIndex.value = "4"
          break;
         }
         case "/Images" : {
          activeIndex.value = "5"
          break;
         }
       }
      }
    )
    return {
      activeIndex,
      handleSelect,
      router
    }
  }
})


</script>

<style scoped lang="stylus">

.header {
  width 100%
  height 100%
  background #fff
  display flex
  justify-content space-between
  align-items center
  padding 0 40px
  box-sizing border-box
  font-weight bold

  .title {
    font-size 20px
    cursor pointer
  }

  .go-github {
    cursor pointer
    font-size 16px

    .icon {
      font-size 20px
    }
  }
}
</style>
