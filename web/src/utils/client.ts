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
