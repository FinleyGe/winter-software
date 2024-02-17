import { inject } from "vue";
import type { InjectionKey } from "vue";

export type NoticeType = "success" | "error" | "info";

export type Notice = {
  message: string;
  type: NoticeType;
}

export const NoticeKey = Symbol() as
  InjectionKey<
  (notice: Notice) => void
  >;

export function useNotice(): (notice: Notice) =>void {
  return inject(NoticeKey)!;
}
