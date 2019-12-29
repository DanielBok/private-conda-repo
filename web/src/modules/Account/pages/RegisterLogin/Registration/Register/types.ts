import { FormItemProps } from "antd/lib/form";
import { WrappedFormUtils } from "antd/lib/form/Form";

export type ValidateStatus = FormItemProps["validateStatus"];

export type RegistrationForm = {
  form: WrappedFormUtils;
  validateStatus: ValidateStatus;
  setValidateStatus: (status: ValidateStatus) => void;
};
