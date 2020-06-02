import { useImmerReducer } from "use-immer";
import { reduce } from "lodash";

export type Credential = {
  username: string;
  password: string;
};

export type State = Credential & {
  valid: boolean;
  touched: Record<keyof Credential, boolean>;
  errors: Record<keyof Credential, string[]>;
  disabled: boolean;
};

export type Action =
  | {
      type: "SET_USERNAME";
      payload: {
        username: string;
      };
    }
  | {
      type: "SET_PASSWORD";
      payload: {
        password: string;
      };
    }
  | {
      type: "SET_VALID";
      payload: {
        valid: boolean;
      };
    };

const initialState: State = {
  username: "",
  password: "",
  valid: true,
  touched: {
    username: false,
    password: false,
  },
  errors: {
    username: [],
    password: [],
  },
  disabled: true,
};

export const useLoginReducer = () => {
  return useImmerReducer((draft, action: Action) => {
    switch (action.type) {
      case "SET_USERNAME": {
        const { username } = action.payload;
        draft.username = action.payload.username;
        draft.touched.username = true;

        draft.errors.username = hasError({ username }).username;
        break;
      }
      case "SET_PASSWORD": {
        const { password } = action.payload;
        draft.password = password;
        draft.touched.password = true;

        draft.errors.password = hasError({ password }).password;
        break;
      }
      case "SET_VALID":
        draft.valid = action.payload.valid;
        break;
    }

    const allTouched = reduce(draft.touched, (a, x) => a && x, true as boolean);
    draft.disabled = reduce(
      draft.errors,
      (a, x) => a || x.length > 0,
      !allTouched
    );
  }, initialState);
};

const hasError = ({ username, password }: Partial<Credential>) => {
  const errors: State["errors"] = {
    username: [],
    password: [],
  };

  if (username === undefined) {
    errors.username.push("Username is required.");
  } else if (username.trim().length < 2) {
    errors.username.push("Username must be at least 2 characters long.");
  }

  if (!password) {
    errors.password.push("Password is required.");
  } else if (password.length < 4) {
    errors.password.push("Password must be at least 4 characters long.");
  }

  return errors;
};
