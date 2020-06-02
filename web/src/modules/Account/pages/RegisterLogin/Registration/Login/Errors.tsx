import React, { FC } from "react";
import { useLoginContext } from "./hooks";
import { Credential } from "./reducer";

type Props = {
  field: keyof Credential;
};

const Errors: FC<Props> = ({ field }) => {
  const { errors } = useLoginContext().state;
  if (errors[field].length === 0) return null;

  const messages = errors[field].join(" ");
  return <span>{messages}</span>;
};

export default Errors;
