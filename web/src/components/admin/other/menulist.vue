<template>
  <div class="main">
    <n-list class="List" hoverable clickable bordered>
      <template #header><h2>菜单子项：</h2></template>
      <n-spin :show="setMenuShow">
        <n-list-item
          :class="menuValue.ename == 'home' ? 'checked' : ''"
          @click="menuValue = { id: 0, name: '主页', ename: 'home', sort: -1 }"
        >
          <div class="home">
            <n-icon size="26" v-html="home.logo" />
            <div class="listname">
              <h4>{{ home.name }}</h4>
              <h6>{{ home.ename }}</h6>
            </div>
          </div>
        </n-list-item>
        <draggable
          v-model="menulist"
          group="people"
          @start="drag = true"
          @end="drag = false"
          item-key="id"
        >
          >
          <template #item="{ element, index }">
            <n-list-item
              :class="menuValue.id == element.ID ? 'checked' : ''"
              @click="
                menuValue = {
                  id: element.ID,
                  name: element.name,
                  ename: element.ename,
                  sort: index,
                }
              "
            >
              <div class="listcontent">
                <n-icon size="26" v-html="element.logo" />
                <div class="listname">
                  <h4>{{ element.name }}</h4>
                  <span>{{ element.ename }}</span>
                </div>
                <div>
                  <n-button quaternary circle @click="edit(index)">
                    <template #icon>
                      <n-icon><Pencil /></n-icon>
                    </template>
                  </n-button>
                  <n-button quaternary circle @click="del(index)">
                    <template #icon>
                      <n-icon><Trash /></n-icon>
                    </template>
                  </n-button>
                </div>
              </div>
            </n-list-item>
          </template>
        </draggable>
      </n-spin>
      <template #footer>
        <n-space>
          <n-button size="small" secondary strong @click="saveMenu">
            保存
          </n-button>
          <n-button size="small" secondary strong @click="showModalnew = true">
            新建
          </n-button>
          <n-button
            size="small"
            secondary
            strong
            @click="menulist = [...menulist2]"
          >
            重置
          </n-button>
        </n-space>
      </template>
    </n-list>
    <div class="cate"><categoryVue :menuValue="menuValue" /></div>
  </div>
  <n-modal
    v-model:show="showModalnew"
    preset="dialog"
    title="新建菜单子项"
    positive-text="确认"
    negative-text="算了"
    @positive-click="newMenu"
  >
    <menueditVue :value="newvalue" />
  </n-modal>
  <n-modal
    v-model:show="showModaledit"
    preset="dialog"
    title="编辑菜单子项"
    positive-text="确认"
    negative-text="算了"
    @positive-click="editMenu"
  >
    <menueditVue :value="editvalue" />
  </n-modal>
</template>

<script setup>
import { onMounted, ref } from "vue";
import draggable from "vuedraggable";
import axios from "axios";
import { Pencil, Trash } from "@vicons/ionicons5";
import menueditVue from "./menuedit.vue";
import { useMessage } from "naive-ui";
import categoryVue from "./category.vue";
const home = ref({
  id: -1,
  name: "主页",
  ename: "Home",
  logo: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 512 512"><path d="M261.56 101.28a8 8 0 0 0-11.06 0L66.4 277.15a8 8 0 0 0-2.47 5.79L63.9 448a32 32 0 0 0 32 32H192a16 16 0 0 0 16-16V328a8 8 0 0 1 8-8h80a8 8 0 0 1 8 8v136a16 16 0 0 0 16 16h96.06a32 32 0 0 0 32-32V282.94a8 8 0 0 0-2.47-5.79z" fill="currentColor"></path><path d="M490.91 244.15l-74.8-71.56V64a16 16 0 0 0-16-16h-48a16 16 0 0 0-16 16v32l-57.92-55.38C272.77 35.14 264.71 32 256 32c-8.68 0-16.72 3.14-22.14 8.63l-212.7 203.5c-6.22 6-7 15.87-1.34 22.37A16 16 0 0 0 43 267.56L250.5 69.28a8 8 0 0 1 11.06 0l207.52 198.28a16 16 0 0 0 22.59-.44c6.14-6.36 5.63-16.86-.76-22.97z" fill="currentColor"></path></svg>`,
});
const showModalnew = ref(false);
const showModaledit = ref(false);
const menulist = ref([]);
const menulist2 = ref([]);
const drag = ref(false);
const newvalue = ref({
  ty: "new",
  name: "",
  ename: "",
  logo: "",
  link: "",
});
const editvalue = ref({});
let savedata = [];
const setMenuShow = ref(false);
const message = useMessage();
const menuValue = ref({ id: 0, name: "主页", ename: "home", sort: -1 });

function newMenu() {
  menulist.value.push(newvalue.value);
  newvalue.value = {
    ty: "new",
    name: "",
    ename: "",
    logo: "",
    link: "",
  };
}
function editMenu() {
  menulist.value[editvalue.value.index].name = editvalue.value.name;
  menulist.value[editvalue.value.index].ename = editvalue.value.ename;
  menulist.value[editvalue.value.index].logo = editvalue.value.logo;
  menulist.value[editvalue.value.index].link = editvalue.value.link;
  menulist.value[editvalue.value.index].ty = "up";
}
function edit(index) {
  editvalue.value = {
    index: index,
    name: menulist.value[index].name,
    ename: menulist.value[index].ename,
    logo: menulist.value[index].logo,
    link: menulist.value[index].link,
  };
  showModaledit.value = true;
}

function del(index) {
  savedata.push({ type: "remove", id: menulist.value[index].ID });
  menulist.value.splice(index, 1);
}

function saveMenu() {
  setMenuShow.value = true;
  menulist.value.map((item, index) => {
    let da;
    if (item.ty == "new") {
      da = {
        type: "new",
        sort: index,
        name: item.name,
        ename: item.ename,
        logo: item.logo,
        link: item.link,
      };
    } else if (item.ty == "up") {
      da = {
        type: "updata",
        id: item.ID,
        sort: index,
        name: item.name,
        ename: item.ename,
        logo: item.logo,
        link: item.link,
      };
    } else if (item.Sort != index) {
      da = { type: "sort", id: item.ID, sort: index };
    }
    if (!savedata.includes(da) && da != null) {
      savedata.push(da);
    }
  });
  axios
    .put("/api/v1/menuchild/set", savedata)
    .then((res) => {
      if (res.data.status == 200) {
        message.success("菜单子项保存成功！！");
        menulist.value = res.data.data;
        menulist2.value = [...menulist.value];
        setMenuShow.value = false;
        savedata = [];
      } else {
        message.error(`保存失败  ${res.data.message}`);
        setMenuShow.value = false;
      }
    })
    .error((e) => {
      message.error(`保存失败  ${e}`);
      setMenuShow.value = false;
    });
}

onMounted(() => {
  axios.get("/api/v1/menuchild").then((res) => {
    if (res.data.status == 200) {
      menulist.value = res.data.data;
      menulist2.value = [...menulist.value];
    }
  });
});
</script>

<style lang="scss" scoped>
@import "list.scss";
.main {
  display: flex;
  height: 100%;
}

.home {
  display: flex;
  align-items: center;
  justify-content: space-around;
}

.cate {
  display: flex;
  margin-left: 20px;
}
</style>
