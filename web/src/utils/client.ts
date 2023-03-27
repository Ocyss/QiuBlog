import { createDiscreteApi } from "naive-ui";
const {
  message: _message,
  notification: _notification,
  dialog: _dialog,
  loadingBar: _loadingBar,
} = createDiscreteApi(["message", "dialog", "notification", "loadingBar"]);

export const loadingBar = {
  start() {
    if (!import.meta.env.SSR) {
      _loadingBar.start();
    }
  },
  error() {
    if (!import.meta.env.SSR) {
      _loadingBar.error();
    }
  },
  finish() {
    if (!import.meta.env.SSR) {
      _loadingBar.finish();
    }
  },
};

export const message = {
  create(content, options?) {
    if (!import.meta.env.SSR) {
      return _message.create(content, options);
    }
  },
  error(content, options?) {
    if (!import.meta.env.SSR) {
      return _message.error(content, options);
    }
  },
  info(content, options?) {
    if (!import.meta.env.SSR) {
      return _message.info(content, options);
    }
  },
  loading(content, options?) {
    if (!import.meta.env.SSR) {
      return _message.loading(content, options);
    }
  },
  success(content, options?) {
    if (!import.meta.env.SSR) {
      return _message.success(content, options);
    }
  },
  warning(content, options?) {
    if (!import.meta.env.SSR) {
      return _message.warning(content, options);
    }
  },
};
