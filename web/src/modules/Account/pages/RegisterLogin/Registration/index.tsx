import { Radio } from "antd";
import React, { useState } from "react";
import styles from "./styles.less";

import LoginForm from "./Login";
import RegistrationForm from "./Register";

export default () => {
  const [action, setAction] = useState<"Register" | "Login">("Register");

  return (
    <div>
      <Radio.Group
        className={styles.radioGroup}
        value={action}
        size="large"
        buttonStyle="solid"
        onChange={e => setAction(e.target.value)}
      >
        <Radio.Button value="Register" className={styles.button}>
          Register
        </Radio.Button>
        <Radio.Button value="Login" className={styles.button}>
          Login
        </Radio.Button>
      </Radio.Group>

      <div className={styles.form}>
        {action === "Login" ? <LoginForm /> : <RegistrationForm />}
      </div>
    </div>
  );
};
