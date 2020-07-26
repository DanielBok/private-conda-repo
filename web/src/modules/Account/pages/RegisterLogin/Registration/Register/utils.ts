import { RegistrationForm } from "@/features/channel/types";
import { useRootSelector } from "@/infrastructure/hooks";
import { ValidateStatus } from "antd/es/form/FormItem";

export const useDetails = (
  key: keyof Pick<
    RegistrationForm,
    "channel" | "email" | "password" | "confirm"
  >
): [string, string, ValidateStatus] =>
  useRootSelector(({ channel: { form } }) => {
    const errors = form.errors[key];
    const status = form.pristine[key]
      ? ""
      : errors.length
      ? "error"
      : "success";

    return [form[key], form.errors[key], status];
  });

export const useDisabled = () =>
  useRootSelector(
    ({
      channel: {
        form: { pristine, errors, confirm, password },
      },
    }) =>
      // disabled if
      // any fields have not been touched (edited)
      Object.values(pristine).some((pristine) => pristine) ||
      // or any fields have error
      Object.values(errors).some((e) => e.length > 0) ||
      password !== confirm
  );
