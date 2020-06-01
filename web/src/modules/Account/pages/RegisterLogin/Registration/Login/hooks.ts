import { FormInstance } from "antd/es/form";
import React, { useContext } from "react";

export const LoginContext = React.createContext<{
  form: FormInstance;
  submit: () => void;
  valid: boolean;
}>({
  form: {} as any,
  submit: () => {},
  valid: true,
});

export const useLoginContext = () => useContext(LoginContext);
