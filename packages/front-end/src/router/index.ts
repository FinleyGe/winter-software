import { useUserStore } from "@/store";
import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/student",
    name: "Scaffold",
    component: () => import("@views/Scaffold.vue"),
    children: [
      {
        path: "/",
        name: "StudentInfo",
        components: {
          default: () => import("@views/student/StudentInfo.vue"),
          header: () => import("@views/student/Header.vue"),
        },
        meta: {
          title:"学生信息"
        }
      }
      // , {
      //   path: "teacher",
      //   name: "StudentTeacher",
      //   components: {
      //     default: () => import("@views/student/Teacher.vue"),
      //     header: () => import("@views/student/Header.vue"),
      //   },
      //   meta: {
      //     title:"教师信息"
      //   }
      // }
    ],
  }, {
    path: "/teacher",
    name: "Scaffold",
    component: () => import("@views/Scaffold.vue"),
    children: [
      {
        path: "",
        name: "Application",
        components: {
          default: () => import("@views/teacher/Application.vue"),
          header: () => import("@views/teacher/Header.vue"),
        },
      }, {
        path: "slist",
        name: "List",
        components: {
          default: () => import("@views/teacher/List.vue"),
          header: () => import("@views/teacher/Header.vue"),
        },
      },{
        path: "cipher",
        name: "Teacher",
        components: {
          default: () => import("@views/teacher/Cipher.vue"),
          header: () => import("@views/teacher/Header.vue"),
        },
      }
    ],
  },{
    path:"/login",
    name:"Login",
    component:()=>import("@views/Login.vue"),
    meta: {
      title:"登录"
    }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, _, next) => {
  const userStore = useUserStore();
  document.title = to.meta.title as string;
  if (to.name !== "Login" && !userStore.isLogin) {
    next({ name: "Login" });
  } else {
    next();
  }
});

export default router;
