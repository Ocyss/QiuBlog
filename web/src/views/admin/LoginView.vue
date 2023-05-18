<template>
  <div class="login_main">
    <Particles id="tsparticles" :options="ParticlesOptions" :particlesInit="particlesInit" />
    <div class="login">
      <div class="change" @click="loginty = !loginty">
        <div v-if="loginty">我要注册</div>
        <div v-else>我要登录</div>
      </div>
      <div class="goHome" @click="router.push({ name: 'home' })">我要回家</div>
      <div class="form">
        <n-form ref="FormRef" :label-width="80" :model="Data" :rules="Rules" size="large" :show-require-mark="false">
          <n-form-item label="账号:" path="username">
            <n-input v-model:value="Data.username" />
          </n-form-item>
          <n-form-item label="密码:" path="password">
            <n-input v-model:value="Data.password" type="password" :maxlength="16" @input="handlePasswordInput"
              @keydown.enter.prevent />
          </n-form-item>
          <n-form-item v-if="!loginty" ref="Password2" label="重复密码:" path="password2">
            <n-input v-model:value="Data.password2" type="password" :disabled="!Data.password" :minlength="6"
              :maxlength="16" @keydown.enter.prevent />
          </n-form-item>
          <GoCaptchaBtn class="go-captcha-btn" v-model:value="status" width="100%" height="80px"
            :image-base64="capt.image" :thumb-base64="capt.thumb" @confirm="handleConfirm"
            @refresh="handleRequestCaptCode" />
          <n-form-item class="button">
            <n-button attr-type="button" @click="Validate">
              {{ loginty ? "登录" : "注册" }}
            </n-button>
          </n-form-item>
        </n-form>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, inject, Ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import api from "@/api";
import { VueCookies } from "vue-cookies";
import { loadFull } from "tsparticles";
import { FormInst, useMessage, FormItemInst } from 'naive-ui'
import lodash from 'lodash'
import GoCaptchaBtn from '@/components/admin/Captcha/GoCaptchaBtn.vue'
const ParticlesOptions = {
  fpsLimit: 60,
  particles: {
    number: {
      value: 0,
      density: {
        enable: false,
        value_area: 800,
      },
    },
    color: {
      value: "#000",
    },
    shape: {
      type: "circle",
      stroke: {
        width: 0,
        color: "#000000",
      },
      polygon: {
        nb_sides: 5,
      },
      image: {
        src: "https://cdn.matteobruni.it/images/particles/github.svg",
        width: 100,
        height: 100,
      },
    },
    opacity: {
      value: 0.5,
      random: false,
      anim: {
        enable: false,
        speed: 1,
        opacity_min: 0.1,
        sync: false,
      },
    },
    size: {
      value: 5,
      random: true,
      anim: {
        enable: false,
        speed: 40,
        size_min: 0.1,
        sync: false,
      },
    },
    line_linked: {
      enable: true,
      distance: 150,
      color: "#000",
      opacity: 0.4,
      width: 1,
    },
    move: {
      enable: true,
      speed: 2,
      direction: "none",
      random: false,
      straight: false,
      out_mode: "out",
      attract: {
        enable: false,
        rotateX: 600,
        rotateY: 1200,
      },
    },
  },
  interactivity: {
    detect_on: "canvas",
    events: {
      onhover: {
        enable: false,
        mode: "repulse",
        parallax: {
          enable: false,
          force: 60,
          smooth: 10,
        },
      },
      onclick: {
        enable: true,
        mode: "emitter",
      },
      resize: true,
    },
    modes: {
      grab: {
        distance: 400,
        line_linked: {
          opacity: 1,
        },
      },
      bubble: {
        distance: 400,
        size: 40,
        duration: 2,
        opacity: 0.8,
        speed: 3,
      },
      repulse: {
        distance: 50,
      },
      push: {
        particles_nb: 4,
      },
      remove: {
        particles_nb: 2,
      },
    },
  },
  retina_detect: true,
  background: {
    color: "#fff",
    image: "",
    position: "50% 50%",
    repeat: "no-repeat",
    size: "cover",
  },
  emitters: {
    direction: "top",
    position: {
      x: 50,
      y: 130,
    },
    rate: {
      delay: 0.1,
    },
    size: {
      width: 100,
      height: 0,
    },
    particles: {
      color: {
        value: ["#5bc0eb", "#fde74c", "#9bc53d", "#e55934", "#fa7921"],
      },
      lineLinked: {
        enable: false,
      },
      opacity: {
        value: 0.5,
      },
      size: {
        value: 400,
        random: {
          enable: true,
          minimumValue: 50,
        },
      },
      move: {
        speed: 10,
        random: false,
        outMode: "destroy",
      },
    },
  },
};
const particlesInit = async (engine) => {
  await loadFull(engine);
};
const $cookies = inject<VueCookies>("$cookies");
const router = useRouter();
const message = useMessage();
const loginty = ref(true);
const FormRef: Ref<FormInst | null> = ref(void 0)
const Password2: Ref<FormItemInst | null> = ref(void 0)
const Data = ref({
  username: "", password: "", password2: "", token: ""
});
const capt = ref({
  image: "",
  thumb: "",
  key: "",
  autoRefreshCount: 0,
})
const status = ref("default")
const Rules = ref({
  username: [{
    required: true,
    trigger: 'blur',
    message: "厉害啊兄弟,账号都不写"
  }],
  password: [{
    required: true,
    trigger: "blur",
    message: "你在干神魔.jpeg"
  }],
  password2: [{
    required: loginty.value ? false : true,
    trigger: ['input', 'blur'],
    message: "请再次输入密码"
  }, {
    validator: (rule, value: string) => {
      return (
        !!Data.value.password &&
        Data.value.password.startsWith(value) &&
        Data.value.password.length >= value.length
      )
    },
    message: '两次密码输入不一致',
    trigger: 'input'
  },
  {
    validator: (rule, value: string) => { return value === Data.value.password },
    message: '两次密码输入不一致',
    trigger: ['blur', 'password-input']
  }]
})

function Validate(e: MouseEvent) {
  e.preventDefault()
  FormRef.value?.validate(async (errors) => {
    if (!errors) {
      if (!Data.value.token) {
        message.error("辣么大个验证框看不到？")
        return
      }
      if (loginty.value) {
        const res = await api.user.login(Data.value)
        if (res.status == 200 && res.data.token) {
          $cookies.set("token", res.token);
          message.success("登陆成功！！");
          router.push({ name: "admin" });
        } else {
          message.success("登录失败");
        }
      } else {
        const res = await api.user.register(Data.value)
        if (res.status == 200) {
          message.success("注册成功！！");
        }
      }
    }
  })
}

function handlePasswordInput() {
  if (!loginty.value && Data.value.password2) {
    Password2.value?.validate({ trigger: 'password-input' })
  }
}

function handleRequestCaptCode() {
  api.utils.getCaptcha().then(res => {
    capt.value.image = res.data.image_base64
    capt.value.thumb = res.data.thumb_base64
    capt.value.key = res.data.captcha_key
  })
}

function handleConfirm(dots) {
  if (lodash.size(dots) <= 0) {
    message.warning(`请进行人机验证再操作`)
    return
  }

  let dotArr = []
  lodash.forEach(dots, (dot) => {
    dotArr.push(dot.x, dot.y)
  })

  api.utils.checkCaptcha(capt.value.key, dotArr.join(',')).then((res) => {

    message.success(`人机验证成功`)
    status.value = 'success'
    capt.value.autoRefreshCount = 0
    Data.value.token = res.data

  }).catch(err => {
    message.warning(`人机验证失败`)
    if (capt.value.autoRefreshCount > 5) {
      capt.value.autoRefreshCount = 0
      status.value = 'over'
      return
    }
    capt.value.autoRefreshCount += 1
    status.value = 'error'
    console.log(capt.value);

  })
}

</script>

<style scoped lang="scss">
.login_main {
  width: 100vw;
  min-height: 100vh;

}

.login {
  height: 390px;
  width: 350px;
  background-color: rgba(0, 0, 0, 0.63);
  position: absolute;
  transform: translate(-50%, -50%);
  top: 50%;
  left: 50%;
  border-radius: 10px;
  backdrop-filter: blur(30px);
  -webkit-backdrop-filter: blur(30px);
  border: 2px solid rgba(255, 255, 255, 0.1);
  box-shadow: 0 0 40px rgba(8, 7, 16, 0.6);
  padding: 50px 35px;


}

.form {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 70%;
}

.button {
  display: flex;
  justify-content: center;
  align-items: center;
}

.change,
.goHome {
  position: absolute;
  top: 2%;
  right: 4%;
  cursor: pointer;

  &:hover {
    color: rgb(42, 114, 248);
  }
}

.goHome {
  top: 93%;
  left: 4%;
}
</style>
