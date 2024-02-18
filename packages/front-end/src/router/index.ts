import { useUserStore } from "@/store";
import {createRouter, createWebHistory} from "vue-router";
import type {RouteRecordRaw} from "vue-router";

const routes: Array<RouteRecordRaw> = [{
  path: "/student",
  name: "Scaffold",
  component: () => import("@views/Scaffold.vue"),
  children: [
    {
      path: "",
      name: "StudentInfo",
      components: {
        default: () => import("@views/student/StudentInfo.vue"),
        header: () => import("@views/student/Header.vue"),
      },
    }, {
      path: "teacher",
      name: "StudentTeacher",
      components: {
        default: () => import("@views/student/Teacher.vue"),
        header: () => import("@views/student/Header.vue"),
      },
    }
  ],
},{
  path:"/login",
  name:"Login",
  component:()=>import("@views/Login.vue")
}
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, _, next) => {
  const userStore = useUserStore();
  if (to.name !== "Login" && !userStore.isLogin) {
    next({name: "Login"});
  } else {
    next();
  }
});

export default router;
