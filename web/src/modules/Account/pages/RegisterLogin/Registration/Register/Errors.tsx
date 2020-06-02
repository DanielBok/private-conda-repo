import React, { FC } from "react";
import { useRegistrationContext } from "./hooks";
import { Fields } from "./reducer";

type Props = {
  field: keyof Fields;
};

const Errors: FC<Props> = ({ field }) => {
  const { errors } = useRegistrationContext().state;
  if (errors[field].length === 0) return null;

  const messages = errors[field].join(" ");
  return <span>{messages}</span>;
};

export default Errors;
