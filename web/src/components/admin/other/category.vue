<template>
  <n-list class="List" hoverable clickable bordered>
    <template #header>
      <h2>
        分类管理
        <span style="font-size: 10px">({{ menuValue.name }})</span>
        ：
      </h2>
    </template>
    <n-spin :show="cateSpin">
      <n-list-item v-if="menuValue.name == '主页'">
        <div>下列内容不显示在主页</div>
      </n-list-item>
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
                      <n-icon><Pencil /></n-icon>
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
          <n-input v-model:value="newName" type="text" placeholder="输入name" />
        </n-popconfirm>
        <n-button size="small" secondary strong>重置</n-button>
      </n-space>
    </template>
  </n-list>
  <n-list class="List" hoverable clickable bordered>
    <template #header>
      <h2>
        {{ menuValue.name == "主页" ? "所有类别：" : "未使用类别：" }}
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
<script setup>
import draggable from "vuedraggable";
import { ref, onMounted } from "vue";
import { Pencil, Trash } from "@vicons/ionicons5";
import axios from "axios";
import { useMessage } from "naive-ui";

const message = useMessage();
const menuValue = ref({ id: 0, name: "主页", ename: "home", sort: -1 });
const setName = ref("");
const newName = ref("");
const drag = ref(false);
const cateSpin = ref(false);

const categoryList = ref([]);
const CheckList = ref([
  { name: "a1" },
  { name: "a2" },
  { name: "a3" },
  { name: "a4" },
]);
const uncheckedList = ref([
  { name: "b1" },
  { name: "b2" },
  { name: "b3" },
  { name: "b4" },
]);

function setNamepositive(index) {
  CheckList.value[index].name = setName.value;
}
function AddCategory() {
  if (menuValue.value.name == "主页") {
    uncheckedList.value.push({ name: newName.value });
  } else {
    CheckList.value.push({ name: newName.value });
  }
  categoryList.value.push({ name: newName.value, ty: "new" });
  newName.value = "";
}

function toggleMenu(data) {
  menuValue.value = data;
}
function save() {
  message.error("未完成！！！！！！！！！！");
}

defineExpose({ toggleMenu });
onMounted(() => {
  axios.get("/api/v1/category").then((res) => {
    if (res.data.status == 200) {
      categoryList.value = res.data.data;
    } else {
      message.error(res.data.message);
    }
  });
});
</script>

<style lang="scss" scoped>
@import "list.scss";
</style>
