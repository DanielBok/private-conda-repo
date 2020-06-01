import { FormInstance, FormItemProps } from "antd/lib/form";

export type ValidateStatus = FormItemProps["validateStatus"];

export type RegistrationForm = {
  form: FormInstance;
  validateStatus: ValidateStatus;
  setValidateStatus: (status: ValidateStatus) => void;
};
