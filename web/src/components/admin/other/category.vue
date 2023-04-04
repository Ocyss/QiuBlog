<template>
  <n-list class="List" hoverable clickable bordered>
    <template #header>
      <h2 v-if="menuValue.id != 0">
        分类管理
        <span style="font-size: 10px">({{ menuValue.name }})</span>
        ：
      </h2>
      <h2 v-else>主页不显示：</h2>
    </template>
    <n-spin :show="cateSpin">
      <draggable
        v-model="CheckList"
        group="cat"
        @start="drag = true"
        @end="drag = false"
        item-key="id"
      >
        <template #item="{ element, index }">
          <n-list-item>
            <div class="listcontent">
              <h3>{{ element.name }}</h3>
              <n-popconfirm
                @positive-click="setNamepositive(index)"
                :show-icon="false"
              >
                <template #trigger>
                  <n-button quaternary circle @click="setName = element.name">
                    <template #icon>
                      <n-icon>
                        <Pencil />
                      </n-icon>
                    </template>
                  </n-button>
                </template>
                <n-input
                  v-model:value="setName"
                  type="text"
                  placeholder="输入name"
                />
              </n-popconfirm>
            </div>
          </n-list-item>
        </template>
      </draggable>
    </n-spin>
    <template #footer>
      <n-space>
        <n-button size="small" secondary strong @click="save">保存</n-button>
        <n-popconfirm @positive-click="AddCategory" :show-icon="false">
          <template #trigger>
            <n-button size="small" secondary strong>新建</n-button>
          </template>
          <n-input
            v-model:value="newData.name"
            type="text"
            placeholder="输入name"
          />
        </n-popconfirm>
        <n-button size="small" secondary strong>重置</n-button>
      </n-space>
    </template>
  </n-list>
  <n-list class="List" hoverable clickable bordered>
    <template #header>
      <h2>
        {{ menuValue.id == 0 ? "主页显示：" : "未使用类别：" }}
      </h2>
    </template>
    <n-spin :show="cateSpin">
      <draggable
        v-model="uncheckedList"
        group="cat"
        @start="drag = true"
        @end="drag = false"
        item-key="id"
      >
        <template #item="{ element, index }">
          <n-list-item>
            <div class="listcontent">
              <h3>{{ element.name }}</h3>
            </div>
          </n-list-item>
        </template>
      </draggable>
    </n-spin>
  </n-list>
</template>
<script setup lang="ts">
import draggable from "vuedraggable";
import { ref, onMounted, computed } from "vue";
import { Pencil, Trash } from "@vicons/ionicons5";

import { useMessage } from "naive-ui";
import api from "@/api";
const message = useMessage();
//菜单项
const menuValue = ref({ id: 0, name: "主页", ename: "home", sort: -1 });
const setName = ref(""); //修改的名字
const newData = ref({ name: "", homeshow: false, mid: 0 }); //新建的菜单名
const drag = ref(false); //是否在拖动
const cateSpin = ref(false); //是否显示加载中

const categoryList = ref([]); //分类列表
const categoryChange = []; // 变动列表

const CheckList = computed({
  get: () => {
    if (menuValue.value.id == 0) {
      return categoryList.value.filter((item) => {
        return !item.homeshow;
      });
    } else {
      return categoryList.value.filter((item) => {
        return item.mid == menuValue.value.id;
      });
    }
  },
  set: (val) => {
    val.map((item) => {
      if (item.mid == 0) {
        categoryList.value.map((item2, index) => {
          if (item2.id == item.id) {
            categoryList.value[index].mid = menuValue.value.id;
            if (
              categoryChange.every((v) => v.id !== categoryList.value[index].id)
            ) {
              categoryChange.push(categoryList.value[index]);
            }
          }
        });
      }
    });
  },
});

const uncheckedList = computed({
  get: () => {
    if (menuValue.value.id == 0) {
      return categoryList.value.filter((item) => {
        return item.homeshow;
      });
    } else {
      return categoryList.value.filter((item) => {
        return item.mid == null || item.mid == 0;
      });
    }
  },
  set: (val) => {
    val.map((item) => {
      if (item.mid != 0) {
        categoryList.value.map((item2, index) => {
          if (item2.id == item.id) {
            categoryList.value[index].mid = null;
            if (
              categoryChange.every((v) => v.id !== categoryList.value[index].id)
            ) {
              categoryChange.push(categoryList.value[index]);
            }
          }
        });
      }
    });
  },
});

//设置菜单名
function setNamepositive(index) {
  CheckList.value[index].name = setName.value;
}
//添加分类
function AddCategory() {
  api.category.add(newData.value).then((res) => {
    message.success("新建分类成功！");
    categoryList.value.push({ ...newData.value, id: res.id });
    newData.value = { name: "", homeshow: false, mid: 0 };
  });
}

//修改菜单名
function toggleMenu(data) {
  menuValue.value = data;
}

//保存
function save() {
  api.category.put(categoryChange).then((res) => {
    message.success("保存成功！");
  });
}

onMounted(() => {
  //请求分类
  api.category.get().then((res) => {
    categoryList.value = res.data;
  });
});

//函数暴露
defineExpose({ toggleMenu });
</script>

<style lang="scss" scoped>
@import "./list.scss";
</style>
