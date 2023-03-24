export interface Config {
  userInfo: {
    title: string;
    name: string;
    motto: string;
    mottoE: string;
  };
  friendChain: Array<{
    name: string;
    href: string;
  }>;
  global: {
    randomImgApi: string;
  };
}
