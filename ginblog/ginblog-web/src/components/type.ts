const types = {
  404: {
    desc: "抱歉，你访问的页面不存在",
  },
  500: {
    desc: "抱歉，服务器出错了",
  },
};
export interface Menu {
  id: string;
  label: string;
  icon?: string;
  name?: string;
  params?: { [key: string]: string };
  children?: Menu[];
}
export default types;
