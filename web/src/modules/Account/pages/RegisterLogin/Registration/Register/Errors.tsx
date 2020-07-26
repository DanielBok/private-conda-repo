import React, { FC } from "react";

type Props = {
  errors: string;
};

const Errors: FC<Props> = ({ errors }) => {
  if (errors === "") return null;

  return <span>{errors}</span>;
};

export default Errors;
