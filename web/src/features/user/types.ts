export type Store = {
  username: string;
  password: string;

  validated: boolean;
  loading: LoadingState;
};

export type UserInfo = {
  channel: string;
  password: string;
};
