import type {RouteLocationRaw} from "vue-router";

export type NavItem = {
  name: string;
  onClick?: () => void;
  path?: RouteLocationRaw;
};
