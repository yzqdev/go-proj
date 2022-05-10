import type { App, Directive } from "vue";
import { networkErrorMsg } from "@/constants";
import { ElMessage } from "element-plus";

export default function setupNetworkDirective(app: App) {
  function listenerHandler(event: MouseEvent) {
    const hasNetwork = window.navigator.onLine;
    if (!hasNetwork) {
      ElMessage({ type: "error", message: networkErrorMsg });
      event.stopPropagation();
    } else {
      ElMessage({ type: "success", message: "网络正常" });
    }
  }

  const networkDirective: Directive<HTMLElement, boolean | undefined> = {
    mounted(el: HTMLElement, binding) {
      if (binding.value === false) return;
      el.addEventListener("click", listenerHandler, { capture: true });
    },
    unmounted(el: HTMLElement, binding) {
      if (binding.value === false) return;
      el.removeEventListener("click", listenerHandler);
    },
  };

  app.directive("network", networkDirective);
}
