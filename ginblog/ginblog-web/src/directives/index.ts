import type { App } from "vue";
import setupNetworkDirective from "./network";

export function setupDirectives(app: App) {
  setupNetworkDirective(app);
}
