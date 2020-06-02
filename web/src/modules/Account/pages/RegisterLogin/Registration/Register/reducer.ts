import { some } from "lodash";
import { useImmerReducer } from "use-immer";

export type Fields = {
  username: string;
  password: string;
  confirm: string;
  email: string;
};

export type State = Fields & {
  pristine: Record<keyof Fields, boolean>;
  errors: Record<keyof Fields, string[]>;
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
      type: "SET_CONFIRM";
      payload: {
        confirm: string;
      };
    }
  | {
      type: "SET_EMAIL";
      payload: {
        email: string;
      };
    }
  | {
      type: "USERNAME_AVAILABILITY";
      payload: {
        available: boolean;
      };
    };

const initialState: State = {
  username: "",
  password: "",
  confirm: "",
  email: "",
  pristine: {
    username: true,
    password: true,
    confirm: true,
    email: true,
  },
  errors: {
    username: [],
    password: [],
    confirm: [],
    email: [],
  },
  disabled: true,
};

export const useRegistrationReducer = () => {
  return useImmerReducer((draft, action: Action) => {
    switch (action.type) {
      case "SET_USERNAME": {
        const { username } = action.payload;
        draft.username = action.payload.username;
        draft.pristine.username = true;

        draft.errors.username = validateUsername(username);
        break;
      }
      case "SET_PASSWORD": {
        const { password } = action.payload;
        draft.password = password;
        draft.pristine.password = true;

        draft.errors.password = validatePassword(password);
        break;
      }
      case "SET_EMAIL": {
        const email = action.payload.email;
        draft.email = email;
        draft.pristine.email = true;
        draft.errors.email = validateEmail(email);
        break;
      }
      case "SET_CONFIRM": {
        const confirm = action.payload.confirm;
        draft.confirm = confirm;
        draft.pristine.confirm = true;
        draft.errors.confirm = validateConfirm(confirm, draft.password);
        break;
      }
      case "USERNAME_AVAILABILITY": {
        const message = "Username is not available";
        if (action.payload.available) {
          draft.errors.username = draft.errors.username.filter(
            (e) => e !== message
          );
        } else {
          draft.errors.username.push(message);
        }
      }
    }

    const pristine = some(draft.pristine);
    draft.disabled = pristine || some(draft.errors, (e) => e.length > 0);
  }, initialState);
};

const validateUsername = (username: string) => {
  const errors: string[] = [];
  if (username === undefined) {
    errors.push("Username is required.");
  } else if (username.trim().length < 2) {
    errors.push("Username must be at least 2 characters long.");
  }

  return errors;
};

const validatePassword = (password: string) => {
  const errors: string[] = [];
  if (!password) {
    errors.push("Password is required.");
  } else if (password.length < 4) {
    errors.push("Password must be at least 4 characters long.");
  }

  return errors;
};

const validateEmail = (email: string) => {
  const errors: string[] = [];
  if (!email) {
    errors.push("Email is required.");
  } else if (!/^[^@\s]+@[^@\s]+$/.test(email)) {
    errors.push("Email doesn't seem to be valid");
  }

  return errors;
};

const validateConfirm = (confirm: string, password: string) => {
  const errors: string[] = [];
  if (confirm !== password) {
    errors.push("passwords do not match");
  }
  return errors;
};
