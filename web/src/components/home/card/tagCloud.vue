<template>
  <n-card
    class="comment"
    :title="designStore.getLocale ? '标签云' : 'tag cloud'"
    size="small"
    ref="card"
  >
    <svg
      :width="TagData.lengh"
      :height="TagData.lengh"
      @mouseover.self="svgMouseover"
      @mouseout.self="svgMouseout"
    >
      <!-- <circle
        :cx="TagData.lengh / 2"
        :cy="TagData.lengh / 2"
        :r="TagData.RADIUS"
        stroke="black"
        stroke-width="0"
        fill="rgba(0,255,0,0.3)"
      /> -->
      <a
        :href="`#/m/tag?id=${tag.id}`"
        v-for="(tag, index) in tags"
        :key="tag"
        @mouseover="mouseover(index)"
        @mouseout="mouseout(index)"
      >
        <text
          :x="tag.x"
          :y="tag.y"
          :font-size="14 * (600 / (600 - tag.z))"
          :fill-opacity="(400 + tag.z) / 550"
          :style="tag.style"
        >
          {{ tag.name }}
        </text>
      </a>
    </svg>
  </n-card>
</template>

<script setup>
import { ref, onMounted } from "vue";
import api from "@/api";
import { randomRgb } from "@/utils";
import { useDesignSettingStore } from "@/store/modules/designSetting.js";
const designStore = useDesignSettingStore();
let timer = null;
const card = ref(null);
const tags = ref([]);
const cos = Math.cos(Math.PI / 360);
const sin = Math.sin(Math.PI / 360);
const TagData = ref({
  lengh: 50,
  RADIUS: 20,
  speedX: Math.PI / 360,
  speedY: Math.PI / 360,
  tags: [],
});
//圆心X
const CX = computed(() => {
  return TagData.value.lengh / 2.2;
});
//圆心Y
const CY = computed(() => {
  return TagData.value.lengh / 2.2;
});
//计算X轴移动
const rotateX = () => {
  for (let tag of tags.value) {
    var y1 = (tag.y - CY.value) * cos - tag.z * sin + CY.value;
    var z1 = tag.z * cos + (tag.y - CY.value) * sin;
    tag.y = y1;
    tag.z = z1;
  }
};
//计算Y轴移动
const rotateY = (speedY) => {
  for (let tag of tags.value) {
    var x1 = (tag.x - CY.value) * cos - tag.z * sin + CY.value;
    var z1 = tag.z * cos + (tag.x - CY.value) * sin;
    tag.x = x1;
    tag.z = z1;
  }
};
//悬停在标签上
const mouseover = (index) => {
  tags.value[index].style.fill = "rgb(255,0,0)";
  tags.value[index].style.fontWeight = "bold";
  clearInterval(timer);
};

//取消标签悬停
const mouseout = (index) => {
  tags.value[index].style.fill = randomRgb(20, 220, 1);
  tags.value[index].style.fontWeight = null;
  clearInterval(timer);
  timer = setInterval(() => {
    rotateX();
    rotateY();
  }, 12);
};
//移入SVG
const svgMouseover = () => {
  clearInterval(timer);
  timer = setInterval(() => {
    rotateX();
    rotateY();
  }, 30);
};
//移出SVG
const svgMouseout = () => {
  clearInterval(timer);
  timer = setInterval(() => {
    rotateX();
    rotateY();
  }, 10);
};
let watchWidthTimer = null;
//让标签云自适应
const watchWidth = () => {
  if (watchWidthTimer) {
    return;
  }
  //使用节流器
  watchWidthTimer = setTimeout(() => {
    if (card.value) {
      TagData.value.lengh = card.value.$el.clientWidth;
      TagData.value.RADIUS = card.value.$el.clientWidth / 2 - 10;
      tags.value = [];
      TagData.value.tags.map((item, index) => {
        let k = -1 + (2 * (index + 1) - 1) / TagData.value.tags.length;
        let a = Math.acos(k);
        let b = a * Math.sqrt(TagData.value.tags.length * Math.PI);
        item.x = CX.value + TagData.value.RADIUS * Math.sin(a) * Math.cos(b);
        item.y = CY.value + TagData.value.RADIUS * Math.sin(a) * Math.sin(b);
        item.z = TagData.value.RADIUS * Math.cos(a);
        item.style = { fill: randomRgb(20, 220, 1) };
        tags.value.push(item);
      });
      watchWidthTimer = null;
    }
  }, 20);
};

onMounted(() => {
  api.tags.get().then((res) => {
    TagData.value.tags = res.data;
    watchWidth();
  });

  window.addEventListener("resize", watchWidth);
  timer = setInterval(() => {
    rotateX();
    rotateY();
  }, 10);
});
</script>

<style scoped lang="scss">
.comment :deep(.n-card__content) {
  padding: 0px;
}
</style>
