import App from "./App.vue";
import router from "@/router";
import { createApp } from "vue";
import "@styles/style.css";
import "@styles/tailwind.css";

const app = createApp(App);
app.use(router);
app.mount("#app");

// createApp(App).mount("#app");
