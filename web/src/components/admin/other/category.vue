<template>
  <n-list class="List" hoverable clickable bordered>
    <template #header>
      <h2>
        分类管理
        <span style="font-size: 10px">({{ props.menuValue.name }})</span>
        ：
      </h2>
    </template>
    <n-spin :show="cateSpin">
      <draggable
        v-model="categoryList"
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
        <n-button size="small" secondary strong>保存</n-button>
        <n-button size="small" secondary strong @click="AddCategory">
          新建
        </n-button>
        <n-button size="small" secondary strong>重置</n-button>
      </n-space>
    </template>
  </n-list>
  <n-list class="List" hoverable clickable bordered>
    <template #header><h2>未使用类别：</h2></template>
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
import { ref } from "vue";
import { Pencil, Trash } from "@vicons/ionicons5";
const props = defineProps(["menuValue"]);
const setName = ref("");
const menuValue = computed(() => {
  get: () => props.menuValue;
});

const drag = ref(false);
const cateSpin = ref(false);
const categoryList = ref([
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
  categoryList.value[index].name = setName.value;
}
function AddCategory() {}
</script>

<style lang="scss" scoped>
@import "list.scss";
</style>
