import { WrappedFormUtils } from "antd/lib/form/Form";
import React, { useContext } from "react";

export const LoginContext = React.createContext<{
  form: WrappedFormUtils;
  submit: () => void;
  valid: boolean;
}>({
  form: {} as any,
  submit: () => {},
  valid: true
});

export const useLoginContext = () => useContext(LoginContext);
