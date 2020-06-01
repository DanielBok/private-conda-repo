import React, { useContext } from "react";
import { RegistrationForm } from "./types";

export const FormContext = React.createContext<RegistrationForm>({
  form: {} as any,
  validateStatus: "",
  setValidateStatus: () => {},
});

export const useFormContext = () => useContext(FormContext);
