export interface Config {
  userInfo: {
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
