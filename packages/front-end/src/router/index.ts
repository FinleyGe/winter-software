import {createRouter, createWebHistory} from "vue-router";
import type {RouteRecordRaw} from "vue-router";

const routes: Array<RouteRecordRaw> = [{
  path: "/",
  name: "Scaffold",
  component: () => import("@views/Scaffold.vue"),
  children: [
    {
      path: "",
      name: "Home",
      components: {
        default: () => import("@views/Home.vue"),
        header: () => import("@views/student/Header.vue"),
      },
    },
  ],
},{
  path:"/login",
  name:"Login",
  component:()=>import("@views/Login.vue")
}
];

export default createRouter({
  history: createWebHistory(),
  routes,
});
