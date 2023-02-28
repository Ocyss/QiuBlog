export interface Config {
  userInfo: {
    title: string;
    name: string;
    motto: [string, string];
  };
  friendChain: Array<{
    name: string;
    href: string;
  }>;
  global: {
    randomImgApi: string;
  };
}
