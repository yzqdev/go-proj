import http from "./http";
export const notice = (msg: string) => {
  return http.post("/notice", msg);
};
export const loginApi = (msg: string) => {
  // return http.post("/notice", msg);
  return true;
};
export const regApi = (msg: string) => {
  return http.post("/notice", msg);
};
