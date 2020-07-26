export type Store = {
  channel: string;
  password: string;
  validated: boolean;
  form: RegistrationForm;
  loading: {
    availableCheck: LoadingState;
    validation: LoadingState;
  };
};

export type Channel = {
  channel: string;
  password: string;
};

export type RegistrationForm = {
  channel: string;
  password: string;
  confirm: string;
  email: string;

  errors: Record<
    Exclude<keyof RegistrationForm, "errors" | "pristine">,
    string
  >;
  pristine: Record<
    Exclude<keyof RegistrationForm, "errors" | "pristine">,
    boolean
  >;
};
