<template>
  <n-spin class="category" :show="cateSpin">
    <n-data-table :columns="columns" :data="categoryList" :bordered="false" :pagination="{ pageSize: 15 }" />
    <n-space>
      <n-button size="small" secondary strong @click="save">
        保存
      </n-button>
      <n-popconfirm @positive-click="AddCategory" :show-icon="false">
        <template #trigger>
          <n-button size="small" secondary strong>新建</n-button>
        </template>
        <n-input v-model:value="newData.name" type="text" placeholder="输入name" />
      </n-popconfirm>
      <n-button size="small" secondary strong @click="categoryList = [...categoryList1]">
        重置
      </n-button>
    </n-space>
  </n-spin>
</template>

<script setup lang="ts">
import { ref, onMounted, Ref, h, defineComponent, nextTick } from "vue";
import { useMessage, NSelect, NCheckbox, NInput } from "naive-ui";
import type { DataTableColumns } from 'naive-ui'
import api from "@/api";
const props = defineProps(["menulist"])
const message = useMessage();
const newData = ref({ name: "", homeshow: false, mid: 0 }); //新建的菜单名
const cateSpin = ref(true); //是否显示加载中

type Category = {
  id: number
  name: string
  mid: number | undefined,
  homeshow: boolean,
  menushow: boolean
}
const categoryList: Ref<Category[]> = ref([]); //分类列表
const categoryList1: Category[] = []; //分类列表
let categoryChange = {}; // 变动列表

const ShowOrEdit = defineComponent({
  props: {
    value: [String, Number],
    onUpdateValue: [Function, Array]
  },
  setup(props) {
    const isEdit = ref(false)
    const inputRef = ref(null)
    const inputValue = ref(props.value)
    function handleOnClick() {
      isEdit.value = true
      nextTick(() => {
        inputRef.value.focus()
      })
    }
    function handleChange() {
      props.onUpdateValue(inputValue.value)
      isEdit.value = false
    }
    return () =>
      h(
        'div',
        {
          style: 'min-height: 22px',
          onClick: handleOnClick
        },
        isEdit.value
          ? h(NInput, {
            ref: inputRef,
            value: inputValue.value,
            onUpdateValue: (v) => {
              inputValue.value = v
            },
            onChange: handleChange,
            onBlur: handleChange
          })
          : props.value
      )
  }
})

const columns: DataTableColumns<Category> = [
  {
    title: 'ID',
    key: 'id'
  },
  {
    title: 'Name',
    key: 'name',
    render(row) {
      return h(ShowOrEdit, {
        value: row.name,
        onUpdateValue(v) {
          row.name = v
          categoryChange[row.id] = row
        }
      })
    }

  },
  {
    title: '菜单项',
    key: 'mid',
    render(row) {
      return h(
        NSelect,
        {
          placeholder: "未分类",
          options: props.menulist,
          filterable: true,
          value: row.mid == 0 || row.mid == null ? null : row.mid,
          labelField: "name",
          valueField: "id",
          onUpdateValue: (v) => {
            row.mid = v
            categoryChange[row.id] = row
          }
        }
      )
    }
  },
  {
    title: '主页是否显示',
    key: 'homeshow',
    render(row) {
      return h(
        NCheckbox,
        {
          checked: row.homeshow,
          onUpdateChecked: (v) => {
            row.homeshow = v
            categoryChange[row.id] = row
          }
        }
      )
    }
  },
  {
    title: '菜单内是否显示',
    key: 'menushow',
    render(row) {
      return h(
        NCheckbox,
        {
          checked: row.menushow,
          onUpdateChecked: (v) => {
            row.menushow = v
            categoryChange[row.id] = row
          }
        }
      )
    }
  }
]

//添加分类
function AddCategory() {
  api.category.add(newData.value).then((res) => {
    message.success("新建分类成功！");
    categoryList.value.push({ ...newData.value, id: res.data });
    newData.value = { name: "", homeshow: false, mid: 0 };
  });
}

//保存
function save() {
  const List = [];
  for (let key in categoryChange) {
    List.push(categoryChange[key])
  }
  api.category.put(List).then((res) => {
    message.success("保存成功！");
  });
}

onMounted(() => {
  //请求分类
  api.category.get().then((res) => {
    categoryList.value = res.data
    categoryList1.push(...res.data)
    cateSpin.value = false
  });
});

</script>

<style lang="scss" scoped>
.category {
  height: 100%;

  :deep(.n-spin-content) {
    height: 100%;
    display: flex;
    align-items: center;
    flex-direction: column;
    justify-content: space-between;
  }
}
</style>
