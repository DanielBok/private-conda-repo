import { some } from "lodash";
import { useImmerReducer } from "use-immer";

export type Credential = {
  username: string;
  password: string;
};

export type State = Credential & {
  valid: boolean;
  pristine: Record<keyof Credential, boolean>;
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
  pristine: {
    username: true,
    password: true,
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
        draft.pristine.username = false;

        draft.errors.username = hasError({ username }).username;
        break;
      }
      case "SET_PASSWORD": {
        const { password } = action.payload;
        draft.password = password;
        draft.pristine.password = false;

        draft.errors.password = hasError({ password }).password;
        break;
      }
      case "SET_VALID":
        draft.valid = action.payload.valid;
        break;
    }

    const pristine = some(draft.pristine);
    draft.disabled = pristine || some(draft.errors, (e) => e.length > 0);
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
